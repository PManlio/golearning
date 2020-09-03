package main
import(
	"fmt"
	"strconv"
	"regexp"
)

var regex = regexp.MustCompile("[0-9]+")

func main() {
	var stringa = "12 14 15 16 17 18"
	
	var numeristringa = regex.FindAllString(stringa, -1)
	
	fmt.Println(numeristringa)
	
	var valori [6]int
	
	// se siamo interessati all'indice, possiamo
	// scrivere for index, element := .....
	for _, element := range numeristringa {
		var numero, _ = strconv.Atoi(element)
		valori = append(valori, numero)
	}
	fmt.Println(valori)
	
}
