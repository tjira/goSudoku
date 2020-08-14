package main

import (
	"sudoku"
	"time"
	"fmt"
	"flag"
)



func main() {
	var solutions = flag.Int("solutions", 1, "Approximate number of puzzle solutions.")
	var fillSeed = flag.Int64("fill-seed", time.Now().UnixNano(), "Random generator seed for generating filled board.")
	var reduceSeed = flag.Int64("reduce-seed", time.Now().UnixNano(), "Random generator seed for eliminating numbers from filled board.")
	var mass = flag.Int("mass", 1, "Number of generated sudokus.")
	var input = flag.String("input", "sudoku.txt", "Filename where sudoku are stored as string.")
	var txtOutput = flag.String("txt-output", "", "Filename where to save output.")
	var saveSolutions = flag.Bool("save-solutions", true, "Wheter or not to save sudoku solutions.")
	flag.Parse()
	mode := flag.Args()[0]
	if mode == "generate" {
		for i := 0; i < *mass; i++ {
			fmt.Println("─────────────────────────────────────")
			fmt.Printf("Sudoku generation no. %d/%d started.\n", i + 1, *mass)
			fmt.Printf("Sudoku will have up to %d solutions.\n", *solutions)
			fmt.Printf("PRNG seed for filling is %d\n", *fillSeed + int64(i))
			fmt.Printf("PRNG seed for reducing is %d\n", *reduceSeed + int64(i))
			start := time.Now()
			filled := sudoku.Filled(*fillSeed + int64(i))
			puzzle := sudoku.Reduced(filled, *solutions, *reduceSeed + int64(i))
			elapsed := time.Since(start)
			emptyFields := sudoku.CountEmptyFields(puzzle)
			allSolutions := sudoku.AllSolutions(puzzle)
			if *txtOutput != "" {
				sudoku.AppendTxt(puzzle, *txtOutput)
				if *saveSolutions {
					filename := (*txtOutput)[:len(*txtOutput) - 4] + "-solutions.txt"
					sudoku.AppendTxt(filled, filename)
				}
			}
			sudoku.Print(filled)
			sudoku.Print(puzzle)
			fmt.Printf("Sudoku generation took %s.\n", elapsed.Round(time.Millisecond))
			fmt.Printf("Sudoku has %d empty fields.\n", emptyFields)
			fmt.Printf("Sudoku has %d solutions.\n", len(allSolutions))
			fmt.Println("─────────────────────────────────────")
		}
	} else if mode == "solve" {
		puzzles := sudoku.LoadString(*input)
		for i := 0; i < len(puzzles); i++ {
			emptyFields := sudoku.CountEmptyFields(puzzles[i])
			fmt.Println("─────────────────────────────────────")
			fmt.Printf("Sudoku solving no. %d/%d started.\n", i + 1, len(puzzles))
			fmt.Printf("Sudoku has %d empty fields.\n", emptyFields)
			start := time.Now()
			allSolutions := sudoku.AllSolutions(puzzles[i])
			elapsed := time.Since(start)
			if *txtOutput != "" {
				sudoku.AppendTxt(allSolutions[0], *txtOutput)
			}
			sudoku.Print(puzzles[i])
			sudoku.Print(allSolutions[0])
			fmt.Printf("Sudoku solving took %s.\n", elapsed.Round(time.Millisecond))
			fmt.Printf("Sudoku has %d solutions.\n", len(allSolutions))
			fmt.Println("─────────────────────────────────────")
		}
	}
}