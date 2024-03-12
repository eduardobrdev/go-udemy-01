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
