package piscine

func CalculateTab(tab []int) []int{
	result := []int{}
	for i := 0; i <= len(tab)-1; i++ {
		result = append(result, 2*(3*tab[i]+4*tab[i]))
	}
	return result
}