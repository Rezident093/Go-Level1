package main

import (
	"fmt"
	"math"
)

func main() {

	var s int

	fmt.Println("Введите площадь окружности")
	fmt.Scanln(&s)
	fmt.Println("Длина окружности равно:", math.Sqrt(4*math.Pi))
	fmt.Println("Диаметр окружности", math.Sqrt(2*math.Pi))
}
