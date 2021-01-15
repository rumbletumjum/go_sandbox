package main

import "fmt"

type Animal interface {
	Name() string
	Speak() string
}

type dog struct {
	name string
}
type cat struct {
	name string
}

func (d dog) Speak() string {
	return "Woof"
}
func (d dog) Name() string {
	return d.name
}

func (c *cat) Name() string {
	return c.name
}

func (c cat) Speak() string {
	return "meow"
}

func introduce(a Animal) {
	fmt.Printf("Hello, my name is %s. %s\n", a.Name(), a.Speak())
}

func main() {
	doge := dog{"Whopper"}
	//var doges = map[string]dog
	fmt.Print(doge.Speak())
}
