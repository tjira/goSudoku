package sudoku

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func EmptyFields(field [][]int) [][]int {
	var fields [][]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if field[i][j] == 0 {
				fields = append(fields, []int{i, j})
			}
		}
	}
	return fields
}

func Column(field [][]int, index int) []int {
	column := make([]int, 9)
	for i := 0; i < 9; i++ {
		column[i] = field[i][index]
	}
	return column
}

func CountEmptyFields(field [][]int) int {
	var count int
	for i := 0; i < 9; i++ {
		for j :=0; j < 9; j++ {
			if field[i][j] == 0 {
				count++
			}
		}
	}
	return count
}

func CheckAvailability(field [][]int, i, j, n int) bool {
	row := !contains(Row(field, i), n)
	column := !contains(Column(field, j), n)
	square := !contains(Square(field, i/3*3 + j/3), n)
	return row && column && square
}

func FilledFields(field [][]int) [][]int {
	var fields [][]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if field[i][j] != 0 {
				fields = append(fields, []int{i, j})
			}
		}
	}
	return fields
}

func LoadString(filename string) [][][]int {
	data, _ := ioutil.ReadFile(filename)
	sudokuStrings := strings.Split(strings.ReplaceAll(string(data), ".", "0"), "\n")
	fields := make([][][]int, 0)
	for _, sudokuString := range sudokuStrings {
		field := Empty()
		for i, char := range sudokuString {
			number, _ := strconv.Atoi(string(char))
			field[i/9][i%9] = number
		}
		fields = append(fields, field)
	}
	return fields
}

func Print(field [][]int) {
	fmt.Println("╔═══════════╦═══════════╦═══════════╗")
	for i := 0; i < 9; i++ {
		if i != 0 {
			fmt.Printf("\n")
		}
		for j := 0; j < 9; j++ {
			if j % 3 == 0 {
				fmt.Printf("║ ")
			} else {
				fmt.Printf("│ ")
			}
			if field[i][j] == 0 {
				fmt.Printf("  ")
			} else {
				fmt.Printf("%d ", field[i][j])
			}
			if j == 8 {
				fmt.Println("║")
			}
		}
		if i%3 == 2 && i != 8 {
			fmt.Printf("╠═══════════╬═══════════╬═══════════╣")
		} else if i != 8 {
			fmt.Printf("║───┼───┼───║───┼───┼───║───┼───┼───║")
		}
	}
	fmt.Println("╚═══════════╩═══════════╩═══════════╝")
}

func Row(field [][]int, index int) []int {
	row := field[index]
	return row
}

func AppendTxt(field [][]int, filename string) {
	data, _ := ioutil.ReadFile(filename)
	var appendix string
	if _, err := os.Stat(filename); err == nil {
		appendix = "\n"
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			appendix += strconv.Itoa(field[i][j])
		}
	}
	ioutil.WriteFile(filename, append(data, []byte(appendix)...), 0644)
}

func Square(field [][]int, index int) []int {
	var square []int
	for i := index/3*3; i < index/3*3 + 3; i++ {
		for j := index%3*3; j < index%3*3 + 3; j++ {
			square = append(square, field[i][j])
		}
	}
	return square
}
