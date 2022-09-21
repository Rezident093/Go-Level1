package sort

import (
	"fmt"
)

func intertionSort() {
	fmt.Println("Введите значения через пробел")
	var ar [5]int
	for i := range ar {

		fmt.Scan(&ar[i])
	}
	InsertionSort(ar)
}

func InsertionSort(ar [5]int) {
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
	fmt.Println("\n", ar)
}
