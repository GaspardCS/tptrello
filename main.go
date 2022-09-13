package main

import (
	"github.com/01-edu/z01"
)

func main() {
    tab := []int{45, 12, 24, 4, 1}
	tab = SortTab(tab)
	tab = CalculateTab(tab)
	tab = SquareTab(tab)
	tab = PrimeTab(tab)
	for _, k := range tab {
		if k == 7 {
			z01.PrintRune('7')
			z01.PrintRune('\n')
		}
	}
}