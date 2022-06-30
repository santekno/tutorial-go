package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person Person
	Model  string
}

func (a *Android) Talk() {
	a.Person.Talk()
}

func main() {
	a := new(Android)
	a.Person.Name = "Ihsan"
	a.Model = "Samsung"
	// a.Person.Talk()
	a.Talk()
}
