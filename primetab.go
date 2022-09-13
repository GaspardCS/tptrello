package piscine

func IsPrime(nb int) bool {
	if nb <= 0 || nb == 1 {
		return false
	}

	if nb > 0 {
		for i := 2; i <= nb-1; i++ {
			if nb%i == 0 {
				return false
			}
		}
	}
	return true
}

func PrimeTab(tab []int) []int {
	secondtab := []int{}
	for i := 0; i <= len(tab)-1; i++ {
		if IsPrime(int(tab[i])) {
			secondtab = append(secondtab, tab[i])
		}
	}
	return secondtab
}
