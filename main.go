package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
)

func main() {
	n := flag.Int("porta", 9000, "Porta dinamica para rodar o servidor")   // retorna *int (um ponteiro)
	flag.Parse()                              // sem isso, *n fica sempre no valor padrão
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
		fmt.Println("Erro ao iniciar o servidor: ", err)
		return
	}
	defer listener.Close()
	fmt.Printf("[SERVIDOR] Escutando na porta %s\n", PORTA)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar uma conexão: ", err)
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
	buf := make([]byte, 1024)

	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Cliente desconectado ou erro de leitura: %v\n", err)
			break
		}
	}
}