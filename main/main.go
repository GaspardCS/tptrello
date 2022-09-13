package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	tab := []int{45, 12, 24, 4, 1}
	tab = piscine.SortTab()
	for _, k := range tab {
		if k == 7 {
			z01.PrintRune('7')
			z01.PrintRune('\n')
		}
	}
}
