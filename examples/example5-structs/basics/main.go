package main

import "fmt"

type person struct {
	name, surname string
	age           int
}

func main() {
	// ways to create a struct instance:
	// 1)
	var p person
	p.name = "melo"
	p.surname = "pulvirenti"
	p.age = 34

	// 2)
	var pe *person = new(person) // this is a reference
	pe.name = "ciccio"
	pe.surname = "graziani"
	pe.age = 44

	// 3)
	pers := new(person) // will be empty and 0-ed

	// 4)
	var perso = person{"giuseppe", "retribuito(??)", 21}
	// perso has an implicit declaration. If you want to explicit it, just:
	// var perso = person{name: "giuseppe", surname: "retribuito(??)", age: 21}

	// you can assign existing instance of structs to newly created instances
	var per *person = new(person) // this is a reference
	*per = p
	per.age = 58

	fmt.Println(p, pe, per, pers, perso)

	// this will show you instance presence in memory:
	fmt.Println(&p, &pe, &per, &pers, &perso)
}
