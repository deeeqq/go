package main

import (
	"first-step-project/roadmap/stage4"
	"first-step-project/roadmap/stage5"
	"first-step-project/roadmap/stage6"
	"fmt"
)

func main() {
	for {
		var i int
		fmt.Printf("_______________________________________________________________________________________  \n")
		fmt.Printf("1)Основы \n")
		fmt.Printf("2)Ошибки \n")
		fmt.Printf("3)Структуры \n")
		fmt.Printf("4)Вызов функции \n")
		fmt.Printf("5)Рекурсия \n")
		fmt.Printf("6)Интерфейс \n")
		fmt.Printf("Введите номер пункта: ")
		fmt.Scanf("%d", &i)
		switch i {
		case 1:
			stage4.Pracmenu()
		case 2:
			stage5.Errortest()
		case 3:
			stage5.Structuretest()
		case 4:
			stage6.Imp3()
		case 5:
			stage6.Ot0do5rec()
		case 6:
			stage6.Interface1()
		default:
			fmt.Printf("Такого пункта не существует")

		}
	}
}
