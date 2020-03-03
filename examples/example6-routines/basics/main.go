package main

import (
	"fmt"
	"time"
)

type example struct {
	a int
}

func printvalue(e *example) {
	for i := 0; i < 3; i++ {
		fmt.Println(e.a + i)
		time.Sleep(1000000000) // Time unit is nanosecond
		// 1 second = 1000000000 nanoseconds
	}
}

func main() {
	var ex1 = new(example)
	ex1.a = 12
	var ex2 = new(example)
	ex2.a = 2

	go printvalue(ex1)
	go printvalue(ex2)

	/*
		It is necessary that program has instruction to be
		called continously to avoid main thread to die.
		In this case I use a sort of while(TRUE)
	*/
	for {
	}
}
