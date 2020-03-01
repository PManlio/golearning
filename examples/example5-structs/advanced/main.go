// for more: http://www.golangbootcamp.com/book/methods
package main

import (
	"fmt"
	"math"
)

// complier will throw warning because these struct won't be exported
type A struct {
	ax, bx int
	cx     string
}

// B'll have anonymous field "heired" from A
type B struct {
	A
	bx int
	by float64
}

type thirdStruct struct {
	a float64
}

// this method will be owned by thirdStruct
// infact you kinda declare a struct-variable
// wich allows to access struct's elements
func (t thirdStruct) mysqrt() float64 {
	if t.a < 0 {
		return t.a // whatever, here should be thrown an error
	}
	return math.Sqrt(t.a)
}

func main() {
	b := B{A{2, 3, "asd"}, 43, 3.4543}
	fmt.Println(b)

	c := new(thirdStruct)
	c.a = 37.8905
	fmt.Println(c.mysqrt())
}
