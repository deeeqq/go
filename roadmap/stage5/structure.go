package stage5

import (
	"fmt"
)

type EmergencyContact struct {
	name  string
	phone string
}
type Address struct {
	street  string
	city    string
	state   string
	zipcode int
}
type Person struct {
	firstname        string
	lastname         string
	age              int
	height           float32
	IsEmployed       bool
	Salary           float32
	Address          Address
	Skills           []string
	Children         []Person
	EmergencyContact EmergencyContact
}

func Structuretest() {

	var vasilii Person = Person{firstname: "Vasilii", lastname: "Petrov", age: 25, height: 80, IsEmployed: true, Salary: 1000,
		Address: Address{street: "Pushkina ", city: "Moscow", state: "42", zipcode: 6}, Skills: []string{"junior"},
		Children: []Person{}, EmergencyContact: EmergencyContact{name: "Sergei", phone: "+7890908"}}
	var anna Person = Person{firstname: "Anna", lastname: "Andreeva", age: 10, height: 30, IsEmployed: false, Salary: 0,
		Address: Address{street: "Dunaiskiy pr.", city: "Saint-Petersburg", state: "60", zipcode: 10}, Skills: []string{"learning"},
		Children: []Person{}, EmergencyContact: EmergencyContact{name: "Marina", phone: "+7892437"}}

	var marina Person = Person{firstname: "Marina", lastname: "Andreeva", age: 30, height: 65, IsEmployed: true, Salary: 1000,
		Address: Address{street: "Dunaiskiy pr.", city: "Saint-Petersburg", state: "60", zipcode: 10}, Skills: []string{"middle"},
		Children: []Person{anna}, EmergencyContact: EmergencyContact{name: "Petr", phone: "+7890657"}}

	fmt.Println("firstname:", vasilii.firstname)
	fmt.Println("lastname:", vasilii.lastname)
	fmt.Println("age:", vasilii.age)
	fmt.Println("height:", vasilii.height)
	fmt.Println("IsEmployed:", vasilii.IsEmployed)
	fmt.Println("Salary:", vasilii.Salary)
	fmt.Println("Address:", vasilii.Address.city, vasilii.Address.street, vasilii.Address.state, vasilii.Address.zipcode)
	fmt.Println("Skills:", vasilii.Skills)
	fmt.Println("Children:", vasilii.Children)
	fmt.Println("EmergencyContact:", vasilii.EmergencyContact.name, vasilii.EmergencyContact.phone)
	fmt.Printf("------------------------------------------------\n")
	fmt.Println("firstname:", marina.firstname)
	fmt.Println("lastname:", marina.lastname)
	fmt.Println("age:", marina.age)
	fmt.Println("height:", marina.height)
	fmt.Println("IsEmployed:", marina.IsEmployed)
	fmt.Println("Salary:", marina.Salary)
	fmt.Println("Address:", marina.Address.city, marina.Address.street, marina.Address.state, marina.Address.zipcode)
	fmt.Println("Skills:", marina.Skills)
	fmt.Println("Children:", marina.Children)
	fmt.Println("EmergencyContact:", marina.EmergencyContact.name, marina.EmergencyContact.phone)
	fmt.Printf("------------------------------------------------\n")

}
