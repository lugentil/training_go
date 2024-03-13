// Para inicializar digit go mod init <module path>
// <module path> Corresponde ao rep/projeto
package main

// Na linguagem GO é necessario que o código faça parte de um pacote
// É necessário um ponto de entrada para que um programa GO se localize
// 1 Progrmaa só pode ter 1 função "main" por que você só pode ter 1 ponto de entrada

import "fmt"

func main() {
	var conferenceName = "GO conference"
	// Constante não altera durante a execução do programa, sempre manterá o mesmo valor
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend")
}

// Para executar 1 arquivo, executar no terminal: go run <file name> -> Irá compilar e rodar o código
