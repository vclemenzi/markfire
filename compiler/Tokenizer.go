package compiler

import (
	"regexp"
	"strings"

	"github.com/vclemenzi/markfire/utils"
)

// Markdown Token ID
// Text: 0
//  - There is a subtokenizer for bold, italic and link
// Heading: 1
//  - H1: 0
//  - H2: 1
//  - H3: 2
//  - H4: 3
//  - H5: 4
//  - H6: 5
// List: 2
//  - Unordered: 0
//  - Ordered: 1
// Blockquote: 3
//   No one

type Token struct {
	Kind    int
	SubKind int
	Content string
	Line    int
}

func Tokinizer(input string) Token {
	if strings.HasPrefix(input, "#") {
		level := utils.HeadingLevel(input)

		return Token{Kind: 1, SubKind: level, Content: strings.TrimSpace(input[level:])}
	} else if strings.HasPrefix(input, "-") {
		return Token{Kind: 2, SubKind: 0, Content: strings.TrimSpace(input[1:])}
	} else if regexp.MustCompile(`[0-9]+.`).MatchString(input) {
		return Token{Kind: 2, SubKind: 1, Content: strings.TrimSpace(input[2:])}
	} else if strings.HasPrefix(input, ">") {
		return Token{Kind: 3, SubKind: -1, Content: strings.TrimSpace(input[1:])}
	} else {
		return Token{Kind: 0, SubKind: -1, Content: input}
	}
}
