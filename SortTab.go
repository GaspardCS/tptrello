package piscine

import (
	"sort"
)

func SortTab(tab []int) []int {
	sort.Sort(sort.IntSlice(tab))
	return tab
}

/*func SortTab(tab []int) []int {
	i := 0
	u := 0
	var t []int
	for _, r := range tab {
		if r > i {
			t = []int{r}
			u++
		} else {
			t = []int{r}
			u++
		}
	}
	return t
}*/

/*func main() {
	tab := []int{45, 12, 24, 4, 1}
	tab = SortTab(tab)
	fmt.Println(tab)
} */
