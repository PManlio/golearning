package main

import(
	"fmt"
	"math/rand"
)

func n20() int {
	fmt.Println("sono nella funzione n20")
	n := rand.Intn(20)+1
	return n
}

func n10() int {
	fmt.Println("sono nella funzione n10")
	n := rand.Intn(10)+1
	return n
}

// B O H
type myFunctionType = int

func main() {
	var numero = 6
	
	var a myFunctionType
	
	dado := "n20"
	
	if dado == "n20" {
		a = n20()
	}
	
	var result = 0
	
	for i:=0; i < numero; i++ {
		result = result + a
	}
	
	return result
}
