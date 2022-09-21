package fibbonachi

import "errors"

var ErrNegativeNumbers = errors.New("can't be negative numbers")

func Fibbonachi(n int, useOptimized bool) (int, error) {
	if n < 0 {
		return 0, ErrNegativeNumbers
	}

	if useOptimized {
		return cashFibb(n, map[int]int{0: 0, 1: 1}), nil
	}

	return simFibb(n), nil
}

func simFibb(n int) int {
	if n < 2 {
		return n
	}

	return simFibb(n-1) + simFibb(n-2)
}

func cashFibb(n int, cash map[int]int) int {

	if val, ok := cash[n]; ok {
		return val
	}

	cash[n] = cashFibb(n-1, cash) + cashFibb(n-2, cash)

	return cash[n]

}
