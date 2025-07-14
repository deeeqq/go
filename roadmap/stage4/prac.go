package stage4

import (
	"fmt"
)

var a, i, b, x, c int
var array, myMap string

func Dell() {
	fmt.Printf("Введите число: \n")
	fmt.Scanf("%d", &a)
	if a%2 == 0 {
		fmt.Printf("Чётное")
	} else {
		fmt.Printf("Нечётное")
	}

}

func Change() {
	fmt.Printf("Введите число от 1 до 5: \n")
	fmt.Scanf("%d", &i)
	switch i {
	case 1:
		fmt.Printf("one")
	case 2:
		fmt.Printf("two")
	case 3:
		fmt.Printf("three")
	case 4:
		fmt.Printf("four")
	case 5:
		fmt.Printf("five")
	default:
		fmt.Printf("Число вне диапазона")
	}
}
func Cikl() {
	for b := 1; b <= 10; b++ {
		fmt.Println("current number is", 0+b)
	}
}
func Kogda() {
	fmt.Printf("while отсутствует в go")

}
func Massive() {
	array := [5]string{"one", "two", "three", "four", "five"}
	for i, c := range array {
		fmt.Println(i, c)
	}
}
func Karta() {
	myMap := map[int]string{1: "Арбуз", 2: "Дыня", 3: "Виноград"}
	fmt.Println(myMap)
}
func Pracmenu() {
	fmt.Printf("1)if/else \n")
	fmt.Printf("2)switch \n")
	fmt.Printf("3)for \n")
	fmt.Printf("4)while \n")
	fmt.Printf("5)for+array \n")
	fmt.Printf("6)map \n")
	fmt.Printf("Введите номер пункта: ")
	fmt.Scanf("%d", &i)
	switch i {
	case 1:
		Dell()
	case 2:
		Change()
	case 3:
		Cikl()
	case 4:
		Kogda()
	case 5:
		Massive()
	case 6:
		Karta()
	default:
		fmt.Printf("Такого пункта не существует")

	}
}
