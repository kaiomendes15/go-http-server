package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// printf 'POST /x HTTP/1.1\r\nContent-Length: -5\r\n\r\n' | nc localhost 9000
	n := flag.Int("porta", 9000, "Porta dinamica para rodar o servidor")// retorna *int (um ponteiro)
	flag.Parse() // sem isso, *n fica sempre no valor padrão
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

	method, path, err := getRequestLine(reader)
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

	if method == "GET" {
		// Quando o método da requisição for GET, o servidor deve procurar o arquivo correspondente ao path dentro do diretório
		contentPath := filepath.Join("public", path)
		// essa lib ja eh multiplataforma, serve tanto para construir o path no linux `/`, quanto no windows `\`
		fmt.Printf("SYSTEM FILE PATH -> %s\n", contentPath)
		contentType, err := getContentType(contentPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(contentType)

		fileBytes, err := os.ReadFile(contentPath)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				fmt.Println("[ERRO 404] Arquivo não encontrado")
				return // aqui vou montar a resposta com a pagina 404
			}

			fmt.Println("[ERRO] Erro ao ler o arquivo.")
			return
		}

		contentLength := len(fileBytes)
		fmt.Printf("[ARQUIVO] Servindo %s (%s, %d bytes)\n", path, contentType, contentLength)
	}
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

func getRequestLine(reader *bufio.Reader) (string, string, error) {
	requestLine, err := reader.ReadString('\n') // le uma linha so
	if err != nil {
		return "", "", fmt.Errorf("[ERRO] Erro ao ler a request line: %v", err)
	}
	requestLine = strings.TrimSpace(requestLine) // so p tirar os \r\n do final
	// fmt.Printf("\n--- Linha Inicial ---\n%s\n", requestLine)

	requestLineItems := strings.Fields(requestLine) // separa a string independente da quant de espacos em branco (melhor que split(" "))
	if len(requestLineItems) == 3 {
		method := requestLineItems[0]
		path := normalizeHttpPath(requestLineItems[1])
		return method, path, nil
	}

	return "", "", fmt.Errorf("[ERRO] linha inicial do HTTP malformada")


}

func normalizeHttpPath(path string) (string) {
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
			return 0, fmt.Errorf("[ERRO] Erro ao ler o header: %v\n", err)
		}
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}
		// fmt.Println(line)
		if strings.HasPrefix(strings.ToLower(line), "content-length:") {
			_, err := fmt.Sscanf(strings.ToLower(line), "content-length: %d", &contentLength) // endereco de memoria pra atualizar a variavel, coloca %d dentro de contentLenght
			if err != nil {
				return 0, fmt.Errorf("[ERRO] Content-Lenght inválido: %v", err)
			}

			if contentLength < 0 {
				return 0, fmt.Errorf("[ERRO] Content-Lenght inválido\n")
			}
		}
	}

	return contentLength, nil
}