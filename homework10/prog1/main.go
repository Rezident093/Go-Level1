package main

import (
	"config/sort"
	"fmt"
)

func main() {
	fmt.Println("Введите значения через пробел")
	var ar [5]int
	for i := range ar {

		fmt.Scan(&ar[i])
	}

	{
		fmt.Println(sort.InsertionSort)
	}
}
