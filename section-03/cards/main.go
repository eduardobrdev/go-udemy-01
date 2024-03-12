package main

import "fmt"

// Fora da func é possível criar variáveis assim:
var teste string // é obrigatório tipar
const teste2 int = 2 // const obrigatoriamente precisa ser atribuída. Não é obrigatório tipar
var teste3 = true // Não precisa tipar se passar valor
// teste4 := 3.14 -> assim não é possível


func main() {
	// Formas de declaração de variável
	//var card string = "Ace of Spades"
	//var card = "Ace of Spades"
	card := "Ace of Spades"

	// Após declarar a variável, podemos reatribuí-la somente com "="
	card = "Five of Diamonds"

	//card = 12 A reatribuição tem que ser obrigatoriamente do mesmo tipo

	fmt.Println(card)
}

// Tipos básicos no GO
// bool, string, int,float32, float64
