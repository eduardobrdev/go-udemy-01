package main

// package == project == workspace

// https://pkg.go.dev/fmt -> Pacote fmt
// Descrição -> Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
import "fmt"

func main() {
	fmt.Println("Hi there!")
}

// Go CLI

// go build -> apenas compila o programa
// Ex: ao rodar go build main.go foi gerado o arquivo executável main. Para utilizá-lo, basta rodar ./main

// go run -> compila e executa imediatamente o programa

// go fmt -> formata o código

// go install -> compila e instala um pacote

// go get -> baixa o código fonte do pacote de outra pessoa

// go test -> roda qualquer teste associado ao projeto
