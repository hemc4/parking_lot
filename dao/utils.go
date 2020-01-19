package dao

func findAndDeleteItem(s []int, e int) []int {
	for i, a := range s {
		if a == e {
			copy(s[i:], s[i+1:])
			s[len(s)-1] = 0
			s = s[:len(s)-1]
			return s
		}
	}
	return s
}
