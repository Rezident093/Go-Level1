package main

import "fmt"

var fibomap map[uint]uint = map[uint]uint{0: 0, 1: 1, 2: 1, 3: 2, 4: 3, 5: 5}

func fib(n uint) uint {
	if val, ok := fibomap[n]; ok {
		return val
	}
	result := fib(n-1) + fib(n-2)
	fibomap[n] = result
	return result
}

func main() {
	var x uint
	fmt.Println("Введите число Фибоначчи")
	fmt.Scanln(&x)
	n := uint(x)

	fmt.Println(fib(n))
}
