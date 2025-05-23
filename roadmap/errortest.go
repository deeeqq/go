package main

import (
	"errors"
	"fmt"
)

// Функция, выполняющая деление и возвращающая ошибку при некорректном вводе
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль невозможно")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Результат:", result)
}
