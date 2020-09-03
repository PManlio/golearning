package main

import (
	"fmt"
	"strings"
	"regexp"
	"strconv"
)

var regex = regexp.MustCompile("[0-9]+")

func main() {
	fmt.Println("esisto")

	
	var a = "cleric 12"
	var b = "guerriero 5"
	var c = "u l t imo  te    st"
	
	var a2 = strings.Replace(a, " ", "", -1)
	var b2 = strings.Replace(b, " ", "", -1)
	var c2 = strings.Replace(c, " ", "", -1)
	
	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)
	
	var anum = regex.FindString(a)
	// returna una stringa
	//fmt.Println(anum+" "+1)
	
	var a3 = strings.Replace(a2, anum, "", -1)
	fmt.Println(a3)
	
	i,err := strconv.Atoi(anum)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}
