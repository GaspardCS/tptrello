package piscine

func Sqrt(x int) int {
	var z int = 1
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func SquareTab(tab []int) []int {
	var secondTab = []int{}

	for i := 0; i <= len(tab)-1; i++ {
		secondTab = append(secondTab, Sqrt(tab[i]))
	}

	return secondTab
}
