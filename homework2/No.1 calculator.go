package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Println("Введите ширину прямоугольника")
	fmt.Scanln(&a)

	fmt.Println("Введите длину прямоугольника:  ")
	fmt.Scanln(&b)

	fmt.Println("Искомая площадь прямоугольника =  ")
	fmt.Println(a * b)
}
