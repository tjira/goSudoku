# goSudoku

goSudoku is a program to generate valid sudoku boards. All the functions are stored in the vendor/sudoku folder, so you can use it also as a package.

## Installation

### Executable

If you have Go installed, you can build the program by navigating to the folder and running


```shell
make build
```

binary file will be generated in bin folder.

### Library

If you want to use the library copy the vendor folder to your project and run

```shell
go mod vendor
```

in your project root directory.

## Usage

Before using the pogram you can run an example by executing

```shell
make example
```

sudokus will be generated in folder examples and printed in terminal.

### Executable

The executable has two modes: generate and solve. Simple sudoku can be generated by

```shell
./sudoku generate
```

and solved by 

```shell
./sudoku --input=filename.txt solve
```

Several parameter can be passed to both modes as described below

```
--solutions          Approximate number of solutions to the resulting sudoku. (default: 1) (mode: generate)
--fill-seed          Random generator seed for filling the board. (default: time.Now().UnixNano()) (mode: generate)
--reduce-seed        Random generator seed for eliminating the board. (default: time.Now().UnixNano()) (mode: generate)
--mass               Number of generated sudokus. (default: 1) (mode: generate)
--save-solutions     Save solutions of the generated sudoku, only works when outputting to txt. (default: true) (mode: generate)
--txt-output         Filename of txt file to output sudoku. (default: None) (mode: generate,solve)
--input              Filename of txt file with unsolved sudokus. (default: sudoku.txt) (mode: solve)
```