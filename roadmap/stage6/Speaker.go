package stage6

import "fmt"

type Speaker interface {
	Speak() string
}
type Dog struct{ name string }
type Person struct{ name string }

func (c Dog) Speak() string {
	fmt.Println(c.name, "Dog speak")
	return "Dog speak"
}
func (a Person) Speak() string {
	fmt.Println(a.name, "Person speak")
	return "Person speak"

}
func Interface1() {

	var Rex Speaker = Dog{}
	var Alex Speaker = Person{}
	Rex.Speak()
	Alex.Speak()
}
