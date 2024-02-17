package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/vclemenzi/markfire/compiler"
)

func main() {
	useOutput := flag.String("o", "", "Output file")
	useInput := flag.String("i", "", "Input file")

	flag.Parse()

	if *useOutput == "" || *useInput == "" {
		fmt.Println("Welcome to markfire!, Markfire is a simple markdown to html converter")
		fmt.Println("Usage: markfire -i <input file> -o <output file>")

		os.Exit(0)
	}

	content, err := os.ReadFile(*useInput)

	if err != nil {
		fmt.Println("err(intenal): impossible to read the file")
		os.Exit(1)
	}

	var tokens []compiler.Token

	for _, line := range strings.Split(string(content), "\n") {
		tokens = append(tokens, compiler.Tokinizer(line))
	}

	fmt.Println(tokens)
}
