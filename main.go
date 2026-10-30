package main

import (
	"fmt"
	"net"
)

func main() {
	PORTA := ":9000"
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