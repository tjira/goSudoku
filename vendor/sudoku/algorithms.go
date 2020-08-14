package sudoku

import "math/rand"

func Empty() [][]int {
	field := make([][]int, 9)
	for i := 0; i < 9; i++ {
		field[i] = make([]int, 9)
	}
	return field
}

func Filled(seed int64) [][]int {
	field := Empty()
	Fill(&field, seed)
	return field
}

func Fill(field *[][]int, seed int64) bool {
	rand.Seed(seed)
	emptyFields := EmptyFields(*field)
	if len(emptyFields) == 0 {
		return true
	}
	i, j := emptyFields[0][0], emptyFields[0][1]
	for _, number := range rand.Perm(10) {
		if CheckAvailability(*field, i, j, number) {
			(*field)[i][j] = number
			if Fill(field, seed + 1) {
				return true
			}
			(*field)[i][j] = 0
		}
	}
	return false
}

func Reduced(filled [][]int, noSolutions int, seed int64) [][]int {
	field := Empty()
	for i := 0; i < 9; i++ {
		copy(field[i], filled[i])
	}
	Reduce(&field, noSolutions, seed)
	return field
}

func Reduce(field *[][]int, noSolutions int, seed int64) {
	rand.Seed(seed)
	filledFields := FilledFields(*field)
	rand.Shuffle(len(filledFields), func(i, j int) {filledFields[i], filledFields[j] = filledFields[j], filledFields[i]})
	for _, loc := range filledFields {
		prev := (*field)[loc[0]][loc[1]]
		(*field)[loc[0]][loc[1]] = 0
		if len(Solutions(*field, noSolutions + 1)) > noSolutions {
			(*field)[loc[0]][loc[1]] = prev
		}
	}
}

func Solutions(puzzle [][]int, n int) [][][]int {
	solutions := make([][][]int, 0)
	FindSolutions(&solutions, &puzzle, n)
	return solutions
}

func AllSolutions(puzzle [][]int) [][][]int {
	solutions := make([][][]int, 0)
	FindSolutions(&solutions, &puzzle, 1000000)
	return solutions
}

func FindSolutions(solutions *[][][]int, puzzle *[][]int, n int) bool {
	emptyFields := EmptyFields(*puzzle)
	if len(emptyFields) == 0 {
		solution := Empty()
		for i := 0; i < 9; i++ {
			copy(solution[i], (*puzzle)[i])
		}
		*solutions = append(*solutions, solution)
		return false
	}
	i, j := emptyFields[0][0], emptyFields[0][1]
	for number := 1; number < 10; number++ {
		if CheckAvailability(*puzzle, i, j, number) && len(*solutions) < n {
			(*puzzle)[i][j] = number
			if FindSolutions(solutions, puzzle, n) {
				return true
			}
			(*puzzle)[i][j] = 0
		}
	}
	return false
}
