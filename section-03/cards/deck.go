package main

// Para funcionar é necessário executar go run main.go deck.go

import "fmt"

// Criando o novo tipo "deck"
// deck -> slice de strings
type deck []string

// Passando o (d deck) antes do nome da função, permite acessá-la após declarar o tipo "deck" a uma variável
// Ex: teste := deck{"teste1", "teste2"}
// Ex: teste.print() -> permite usar assim pois teste é do tipo deck
// Output: Index: 0 Value: teste1 \n Index: 1 Value: teste2
func (d deck) print() {
	for i, card := range d {
		fmt.Println("Index:", i, "Value:", card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	// Sabemos que o range retorna o indice e o valor do array, neste exemplo como não será utilizado o indice, podemos passar _
	// para que o programa não reclame de variável não utilizada
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}
