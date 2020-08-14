build:
	go build -o bin/sudoku main.go

clean:
	rm -rf bin examples

example:
	mkdir -p examples
	rm -rf examples/examples.txt examples/examples-solutions.txt examples/examples-solved.txt
	go run main.go --mass=5 --txt-output=examples/examples.txt generate
	go run main.go --input=examples/examples.txt --txt-output=examples/examples-solved.txt solve