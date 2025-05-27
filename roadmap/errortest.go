package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Code int
	Msg  string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Ошибка %d: %s", e.Code, e.Msg)
}

func doSomething() error {
	return &CustomError{Code: 404, Msg: "Ресурс не найден"}
}
func mayPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Поймана паника:", r)
		}
	}()

	panic("непредвиденная ошибка!")
}

// Функция, выполняющая деление и возвращающая ошибку при некорректном вводе
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль невозможно")
	}
	return a / b, nil
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println("Произошла ошибка:", err)
	}
	mayPanic()
	fmt.Println("Программа продолжает работу")
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Результат:", result)
}
