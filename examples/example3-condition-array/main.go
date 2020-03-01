package main

import "fmt"

func main() {
	fmt.Print("test\n\n")

	array1 := [3]int{} // importante: se vuoi dichiarare un array vuoto così
	// devi necessariamente implementarlo con le {} vuote
	fmt.Println(array1)

	fmt.Print("\n\n")

	// gli array li puoi implementare pure così
	var array2 [2]string
	array2[0] = "carmelo"
	array2[1] = "giuseppe"
	fmt.Println(array2)

	var giuseppe = array2[1]
	var carmelo = array2[0]

	fmt.Println(giuseppe, carmelo)

	if carmelo == "carmelo" {
		fmt.Print("grazie al cazzo\n")
	}

	/*	l'if con queste condizioni non funziona, non capisco perché
		if carmelo == "carmelo" ; giuseppe = "giuseppe" {
			fmt.Print("grazie al cazzo")
		}
	*/
	a := 3

	// il comportamento dei due switch scritti sotto è analogo

	switch {
	case a == 3:
		fmt.Print("\nsì manlio, a è uguale a 3, l'hai settato tu.\n")
	}

	switch a {
	case 3:
		fmt.Print("\nsecondo esempio di switch, con a == 3.\n")
	}

	b := 6
	switch b / 3 {
	case 2:
		fmt.Println("b/3 = 2")
	}

Loop:
	for i := 0; i < 6; i++ {
		switch i {
		case 2:
			println("sono nel caso:", i)
		case 4:
			break Loop
		case 1, 3, 5: // nel 5 non ci entra, corretto
			scriviciclo(i)
		}
	}

}

func scriviciclo(i int) {
	fmt.Println(i)
}
