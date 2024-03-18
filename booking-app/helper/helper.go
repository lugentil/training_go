package helper

import "strings"

//Para exportar uma função em GO basta colocar a primeira Letra em caps lock. Obs: Serve para variaveis, constantes e etc -> Variavel global
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") // Verifica se a string email possui o simbolo "@"
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	// isValidCity := city == "Singapore" || city == "London"
	// isInvalidValidCity := city =! "Singapore" && city != "London"
	// Uma maneira de negar uma afirmação, exemplo: !isValidCity

	// Em GO podemos retornar multiplos parâmetros de uma função, diferente de outras linguagens de programação
	return isValidName, isValidEmail, isValidTicketNumber
}
