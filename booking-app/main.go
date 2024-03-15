// Para inicializar digit go mod init <module path>
// <module path> Corresponde ao rep/projeto
package main

// Na linguagem GO é necessario que o código faça parte de um pacote
// É necessário um ponto de entrada para que um programa GO se localize
// 1 Progrmaa só pode ter 1 função "main" por que você só pode ter 1 ponto de entrada

import "fmt"

func main() {
	conferenceName := "GO conference"
	// var conferenceName = "GO conference"
	// Constante não altera durante a execução do programa, sempre manterá o mesmo valor
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// Exemplo de lista: var bookings = [50]string{}
	// var bookings [50]string //outra maneira de declarar lista
	// var bookings []string // Isto é um Slice, representa a mesma forma da lista porém de maneira dinâmica
	bookings := []string{} // Criando com a sintaxe compacta
	fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your e-mail: ")
	fmt.Scan(&email)

	fmt.Printf("Enter your number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	// Em GO para utilizar lista é necessário definir quandos itens poderám ser armazenados nesta lista
	// E é necessario dizer qual o tipo de dado que terá na lista, ou seja não é possivel utilizar INT e STRING juntos.
	// bookings[0] = firstName + " " + lastName  -> Exemplo de inserção em 1 lista
	bookings = append(bookings, firstName+" "+lastName) // Inserindo em nosso slice

	// fmt.Printf("Slice completo: %v\n", bookings)
	// fmt.Printf("Primeiro valor: %v\n", bookings[0])
	// fmt.Printf("Tipo do Slice: %T\n", bookings)
	// fmt.Printf("Tamanho do Slice: %v\n", len(bookings))

	// Função scanf para pedir informações
	fmt.Printf("\nObrigado %v %v por comprar %v tickets! Você irá receber a confirmação da compra no seguinte e-mail: %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("Número de tickets em estoque após a compra: %v", remainingTickets)

	fmt.Printf("Todos Bookings na aplicação: %v\n", bookings)
}

// Para executar 1 arquivo, executar no terminal: go run <file name> -> Irá compilar e rodar o código
