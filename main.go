package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Fields struct {
	Key   string
	Value string
}

func main() {
	// printf 'POST /x HTTP/1.1\r\nContent-Length: -5\r\n\r\n' | nc localhost 9000
	n := flag.Int("porta", 9000, "Porta dinamica para rodar o servidor") // retorna *int (um ponteiro)
	flag.Parse()                                                         // sem isso, *n fica sempre no valor padrão
	fmt.Println(*n)
	PORTA := ""

	if *n >= 1 && *n <= 65535 {
		PORTA = ":" + strconv.Itoa(*n)
	} else {
		PORTA = ":9000"
		fmt.Printf("[SERVIDOR] Porta %d esta no intervalo invalido\n[SERVIDOR] Executando na porta de fallback\n", *n)
	}

	listener, err := net.Listen("tcp", PORTA)

	if err != nil {
		fmt.Println("[ERRO] Erro ao iniciar o servidor: ", err)
		return
	}
	defer listener.Close()
	fmt.Printf("[SERVIDOR] Escutando na porta %s\n", PORTA)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[ERRO] Erro ao aceitar uma conexão: ", err)
			continue
		}

		//           coroutine
		// Criar uma goroutine aqui
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer fmt.Printf("[CONEXÃO] Encerrada %s\n", conn.RemoteAddr())
	defer conn.Close()

	fmt.Printf("[CONEXÃO] Nova conexão de %s\n", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	// lib que permite ler linha, independente de quantos pacotes tcp chegaram

	method, path, version, err := getRequestLine(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	contentLenght, err := getContentLenght(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[REQUEST] Método: %s | Path: %s | Content-Length: %d\n", method, path, contentLenght)

	switch method {
	case "GET":
		response, err := getContent(path, version)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = conn.Write(response)
		if err != nil {
			fmt.Printf("[ERRO] Problema ao retornar response para o client: %v\n", err)
		}
	case "POST":
		buf, err := postContent(path, contentLenght, reader)
		if err != nil {
			fmt.Println(err)
			return
		}

		fields, err := sliceFields(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(fields)
		return
	}
}

func postContent(path string, contentLength int, reader *bufio.Reader) ([]byte, error) {
	buf := make([]byte, contentLength)
	_, err := io.ReadFull(reader, buf)
	if err != nil {
		return nil, fmt.Errorf("[ERRO] Nenhum ou menos bytes que o Content-Length foram lidos.")
	}
	// fmt.Println("POST para ", path)
	return buf, nil
}

func sliceFields(body []byte) ([]Fields, error) {
	strBody := string(body)
	fmt.Printf("[POST] Body bruto (%d bytes): %s\n", len(body), strBody)
	keyValues := strings.Split(strBody, "&")
	// bodyMap := make(map[string]string)
	// essa merda de map nn preservava a ordem dos valores
	// fiz uma struct pra padronizar um objeto chave/valor
	slicedFields := []Fields{}
	for _, kv := range keyValues {
		// strings.Cut retorna: antes do '=', depois do '=', e se achou o '='
		key, value, found := strings.Cut(kv, "=")
		decodedKey, err := decodeValue(key)
		if err != nil {
			return nil, err
		}

		decodedValue, err := decodeValue(value)
		if err != nil {
			return nil, err
		}

		if found {
			fmt.Printf("[POST] Campo: %s = %s\n", decodedKey, decodedValue)
			slicedFields = append(slicedFields, Fields{Key: decodedKey, Value: decodedValue})
		}
	}

	return slicedFields, nil
}

func decodeValue(value string) (string, error) {
	decodedValue, err := url.QueryUnescape(value)
	if err != nil {
		return "", fmt.Errorf("[ERRO] Formato inválido para decode no body: %w", err)
	}

	return decodedValue, nil
}

func getContent(
	path string,
	version string) ([]byte, error) {
	// Quando o método da requisição for GET, o servidor deve procurar o arquivo correspondente ao path dentro do diretório
	contentPath := filepath.Join("public", path)
	// essa lib ja eh multiplataforma, serve tanto para construir o path no linux `/`, quanto no windows `\`
	// fmt.Printf("SYSTEM FILE PATH -> %s\n", contentPath)
	contentType, err := getContentType(contentPath)
	if err != nil {
		return nil, err
	}
	fmt.Println(contentType)

	fileBytes, err := os.ReadFile(contentPath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("[404] Recurso não encontrado:", path)
			notFoundResponse, err2 := notFound(version)
			if err2 != nil {
				return nil, err2
			}

			return notFoundResponse, nil
		}

		return nil, fmt.Errorf("[ERRO] Erro ao ler o arquivo: %v", err)
	}

	contentLength := len(fileBytes)
	fmt.Printf("[ARQUIVO] Servindo %s (%s, %d bytes)\n", path, contentType, contentLength)

	responseHeader := buildHeader(version, 200, "OK", contentType, contentLength)
	fmt.Println(responseHeader)
	response := buildResponse(responseHeader, fileBytes)

	return response, nil
}

func notFound(
	version string,
) ([]byte, error) {
	notFoundPagePath := filepath.Join("public", "not_found.html")
	notFoundPage, err := os.ReadFile(notFoundPagePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, fmt.Errorf("[ERRO 404] Arquivo não encontrado: %v", err)
		}

		return nil, fmt.Errorf("[ERRO] Erro ao ler o arquivo: %v", err)
	}
	contentLength := len(notFoundPage)
	contentType, err := getContentType(notFoundPagePath)
	if err != nil {
		return nil, err
	}

	header := buildHeader(version, 404, "Not Found", contentType, contentLength)
	response := buildResponse(header, notFoundPage)

	return response, nil
}

func buildHeader(
	version string,
	status int,
	message string,
	contentType string,
	contentLength int,
) string {
	header := version + " " + strconv.Itoa(status) + " " + message + "\r\nContent-Type: " + contentType + "\r\n" + "Content-Length: " + strconv.Itoa(contentLength) + "\r\n\r\n"

	return header
}

func buildResponse(header string, body []byte) []byte {
	return append([]byte(header), body...) // ... pra aceitar slice de bytes
}

func getContentType(contentPath string) (string, error) {
	// devo considerar:
	// ./public/style.css
	// ./public/img/imagem.png
	fileExtension := filepath.Ext((contentPath))
	// fmt.Println("File Extension:", fileExtension)
	contentType := ""
	switch strings.TrimSpace(fileExtension) {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".png":
		contentType = "image/png"
	case ".jpg":
		contentType = "image/jpeg"
	default:
		return contentType, fmt.Errorf("[ERRO] formato de arquivo inválido.")
	}

	return contentType, nil
}

func getRequestLine(reader *bufio.Reader) (string, string, string, error) {
	requestLine, err := reader.ReadString('\n') // le uma linha so
	if err != nil {
		return "", "", "", fmt.Errorf("[ERRO] Erro ao ler a request line: %v", err)
	}
	requestLine = strings.TrimSpace(requestLine) // so p tirar os \r\n do final
	// fmt.Printf("\n--- Linha Inicial ---\n%s\n", requestLine)

	requestLineItems := strings.Fields(requestLine) // separa a string independente da quant de espacos em branco (melhor que split(" "))
	if len(requestLineItems) == 3 {
		method := requestLineItems[0]
		path := normalizeHttpPath(requestLineItems[1])
		version := requestLineItems[2]
		return method, path, version, nil
	}

	return "", "", "", fmt.Errorf("[ERRO] linha inicial do HTTP malformada")

}

func normalizeHttpPath(path string) string {
	if path == "/" {
		return "/index.html"
	}

	return path
}

func getContentLenght(reader *bufio.Reader) (int, error) {
	var contentLength int
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return 0, fmt.Errorf("[ERRO] Erro ao ler o header: %v", err)
		}
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}
		// fmt.Println(line)
		if strings.HasPrefix(strings.ToLower(line), "content-length:") {
			_, err := fmt.Sscanf(strings.ToLower(line), "content-length: %d", &contentLength) // endereco de memoria pra atualizar a variavel, coloca %d dentro de contentLenght
			if err != nil {
				return 0, fmt.Errorf("[ERRO] Content-Length inválido: %v", err)
			}

			if contentLength < 0 {
				return 0, fmt.Errorf("[ERRO] Content-Length inválido")
			}
		}
	}

	return contentLength, nil
}
