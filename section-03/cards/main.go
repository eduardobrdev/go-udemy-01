package main

import "fmt"

// Fora da func é possível criar variáveis assim:
var teste string     // é obrigatório tipar
const teste2 int = 2 // const obrigatoriamente precisa ser atribuída. Não é obrigatório tipar
var teste3 = true    // Não precisa tipar se passar valor
// teste4 := 3.14 -> assim não é possível

func main() {
	// Formas de declaração de variável
	//var card string = "Ace of Spades"
	//var card = "Ace of Spades"
	//card := "Ace of Spades"
	// Após declarar a variável, podemos reatribuí-la somente com "="
	//card = "Five of Diamonds"
	//card = 12 A reatribuição tem que ser obrigatoriamente do mesmo tipo

	//Array e Slice
	cards := deck{"Ace of Diamonds", newCard()} // tipo deck declarado no arquivo deck.go, no mesmo package main
	fmt.Println(cards)
	cards = append(cards, "Six of Spades") // mesmo comportamento de um array_push
	fmt.Println(cards)

	// range -> semelhante a um foreach onde i é o indice e o card(pode ser qualquer nome) é o value
	cards.print() // Executa a mesma coisa que o for abaixo. A lógica está no arquivo deck.go
	// for i, card := range cards {
	// 	fmt.Println("index:", i, "value:", card)
	// }

	card := newCard()
	fmt.Println(card)

}

// Para fazer com que a função retorne a string, é necessário tipar o seu retorno
func newCard() string {
	return "Five of Diamonds"
}

// Tipos básicos no GO
// bool, string, int,float32, float64

// Array -> possui tamanho fixo
// Slice -> array que pode aumentar ou diminuir. Todos os elementos tem que ser do mesmo tipo
