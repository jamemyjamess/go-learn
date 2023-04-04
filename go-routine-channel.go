package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	id   int
	name string
}

func main() {

	var Persons []*Person
	Persons = append(Persons, &Person{id: 1, name: "bom"})
	Persons = append(Persons, &Person{id: 2, name: "atisit"})
	Persons = append(Persons, &Person{id: 3, name: "c"})
	Persons = append(Persons, &Person{id: 4, name: "d"})

	chans := make([]chan *Person, len(Persons))
	for index, _ := range chans {
		chans[index] = make(chan *Person)
	}

	for i, v := range Persons {

		go RenameOrder(v, i, chans[i])
	}

	var arrayPerson []Person
	for index, _ := range chans {
		arrayPerson = append(arrayPerson, *<-chans[index])
	}

	fmt.Println(arrayPerson)

}

// goroutine
func RenameOrder(person *Person, index int, ch chan *Person) {
	temp := strconv.Itoa(person.id)
	person.name = person.name + "_" + temp
	ch <- person

}
