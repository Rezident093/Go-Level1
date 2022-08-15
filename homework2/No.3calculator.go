package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Введите любое трезнаяное число")
	fmt.Scan(&a)
	if a < 999 {
		a = a / 100
		fmt.Println((a), "сотни")
		fmt.Println((a), "десятков")
	} else {
		fmt.Println("вы нарушили условие ввода ")
	}
}
