// Para inicializar digit go mod init <module path>
// <module path> Corresponde ao rep/projeto
package main

// Na linguagem GO é necessario que o código faça parte de um pacote
// É necessário um ponto de entrada para que um programa GO se localize
// 1 Programaa só pode ter 1 função "main" por que você só pode ter 1 ponto de entrada

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// Package level variables
var conferenceName = "GO conference"

// var conferenceName = "GO conference"
// Constante não altera durante a execução do programa, sempre manterá o mesmo valor
const conferenceTickets int = 50

var remainingTickets uint = 50

// Exemplo de lista: var bookings = [50]string{}
// var bookings [50]string //outra maneira de declarar lista
// var bookings []string // Isto é um Slice, representa a mesma forma da lista porém de maneira dinâmica
// var bookings = make([]map[string]string, 0) // Criando com a sintaxe compacta
var bookings = make([]UserData, 0) // Criando com a sintaxe compacta

// Ao definir o type, estamos criando dados do tipo userData
// Mixed data type. E defini a estrutura do que os dados do usuário seria

var wg = sync.WaitGroup{}

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetUsers()

	// Podemos dar condições de execução ao nosso for
	// Exemplo for remainingTickets > 00 && len(bookings) < 50 {}
	for {
		// Para loops existe apenas o for em GO
		firstName, lastName, email, userTickets := getUserInput()

		fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)                                              // Este valor é 1, pois estamos apenas com 1 função em paralelo, em caso de 2 funções com a palvra go, colocaremos o valor 2 para esta variavel.
			go sendTicket(userTickets, firstName, lastName, email) // Para aplicar paralelismo, basta colocar a palavra go na frente da chamada da função.

			// Assimilando uma variavel para receber o valor de saída da função getFirtsNames
			firstNames := getFirstNames()

			fmt.Printf("\nO primeiro nome dos compradores: %v\n", firstNames)

			noremainingTickets := remainingTickets == 0
			// if remainingTickets == 0 {
			if noremainingTickets { // end program
				fmt.Println("Os tickets acabaram para esta conferencia. Volte novamente no próximo ano!")
				break
			}

		} else { //else if userTickets == remainingTickets {}
			if !isValidName {
				fmt.Println("O nome informado é muito curto")
			}
			if !isValidEmail {
				fmt.Println("O e-mail não contém o simbolo @")
			}
			if !isValidTicketNumber {
				fmt.Println("Número de tickets que você inseriu é inválido")
			}
			// city := "London"

			// switch city {
			// case "New York":
			// 	// Execute Code for this city
			// case "Singapore":
			// 	// Execute Code for this city
			// case "London", "Mexico":
			// 	// Execute Code for this city
			// case "Berlin":
			// 	// Execute Code for this city
			// case "Hong Kong":
			// 	// Execute Code for this city
			// default: // Caso não seja encontrado nenhum valor, executar o default
			// 	fmt.Println("Cidade inválida")
			// }

			fmt.Printf("O valor informado não está correto, por favor insira novamente.")
			continue
		}
	}
	wg.Wait()
}

// Para executar 1 arquivo, executar no terminal: go run <file name> -> Irá compilar e rodar o código

func greetUsers() {
	fmt.Printf("Bem vindo a %v!\n", conferenceName)
	fmt.Printf("Nós temos o total de %v tickets e %v estão disponíveis atualmente\n", conferenceTickets, remainingTickets)
	fmt.Printf("Garanta seus tickets agora mesmo!\n")

}

func getFirstNames() []string { // Colocamos dentro dos colchetes os parametros de entrada e fora colocamos os parametros que serão retornados / de saida
	firstNames := []string{}
	// Definindo for para iterar o slice Bookings, com a função range conseguimos o index e o valor de cada elemento
	// Tanto para listas ou slices
	for _, booking := range bookings {
		// firstNames = append(firstNames, booking["firstName"]) -> adicionando dados em um MAP
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	// Em GO para utilizar lista é necessário definir quandos itens poderám ser armazenados nesta lista
	// E é necessario dizer qual o tipo de dado que terá na lista, ou seja não é possivel utilizar INT e STRING juntos.
	// bookings[0] = firstName + " " + lastName  -> Exemplo de inserção em 1 lista

	// Create a map for a user
	// Slice example: var slice[]string
	// map example: var map[string]string

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// var userData = make(map[string]string) // Nós não podemos mesclar tipos de dados no map usando GO
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["númeroDeTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData) // Inserindo em nosso slice

	// fmt.Printf("Slice completo: %v\n", bookings)
	// fmt.Printf("Primeiro valor: %v\n", bookings[0])
	// fmt.Printf("Tipo do Slice: %T\n", bookings)
	// fmt.Printf("Tamanho do Slice: %v\n", len(bookings))

	// Função scanf para pedir informações
	fmt.Printf("Bookings list: %v", bookings)
	fmt.Printf("\nObrigado %v %v por comprar %v tickets! Você irá receber a confirmação da compra no seguinte e-mail: %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("Número de tickets em estoque da %v após a compra: %v\n", conferenceName, remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// Função para aplicar paralelismo
	// GO utiliza o que é chamado de Green Thread
	// Em GO utilizamos GOroutine, que é uma abstração do conceito de threads
	// Gerenciado pelo GO runtime, ele se comunicam apenas em um nivel alto de GOroutines
	// Com isso ele fica mais leve e mais barato
	// Podendo rodar diversas GO Routines sem afetar a performance das máquinas

	// Sincronizando -> Utilizamos um package chamado sync que integra a funcionalidade de sincronizar o paralelismo
	// Precisamos montar a sincronização utilizando 3 funções deste pacote
	// Add, Wait e Done. Add -> Define o número de GORoutines para esperar e também gera um contador
	// Função Wait -> Funçao que espera o WaitGroup counter chegar a 0
	// Função Done-> Decrementa o contador em 1, então é chamado pelo GoRoutine para saber quando acabou o processamento.
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	//Sprintf permite salvar o print diretamente em uma variavel
	fmt.Println("############")
	fmt.Printf("Sending ticket: \n%v \nto email address: %v\n", ticket, email)
	fmt.Println("############")
	wg.Done()
}
