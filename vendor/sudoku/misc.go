package sudoku

func contains(array []int, number int) bool {
	for _, element := range array {
		if element == number {
			return true
		}
	}
	return false
}
