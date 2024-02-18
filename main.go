package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/vclemenzi/markfire/generator"
	"github.com/vclemenzi/markfire/tokenizer"
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

	var html string
	var tokens []tokenizer.Token
	var previousToken tokenizer.Token
	openableTokens := tokenizer.OpenableTokens{
		Configuration: &tokenizer.Configuration{},
		List:          &tokenizer.List{},
		Blockquote:    &tokenizer.Blockquote{},
		Codeblock:     &tokenizer.Codeblock{},
	}

	for i, line := range strings.Split(string(content), "\n") {
		token := tokenizer.Tokenizer(line, &openableTokens, i)

		html += generator.Html(token, previousToken, &openableTokens, i)
		html += "\n"

		tokens = append(tokens, token)
		previousToken = token
	}

	os.WriteFile(*useOutput, []byte(generator.GetHtmlStrcture(tokens, html, generator.GetHtmlStyle())), 0644)
}
