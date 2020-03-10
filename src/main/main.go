package main

import (
	"fmt"
	"model"
)

func main() {

	person := model.NewPerson("john")
	person.SetAge(25)
	person.SetSal(300)
	fmt.Println(person)
	fmt.Println(person.Name,person.GetAge(),person.GetSal())

}
