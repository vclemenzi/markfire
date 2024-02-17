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

type OpenableTokens struct {
	List       *List
	Blockquote *Blockquote
}

type List struct {
	IsOpen  bool
	Index   int
	Subkind int
	Closure int
}

type Blockquote struct {
	IsOpen  bool
	Index   int
	Closure int
}

func Tokenizer(input string, openableTokens *OpenableTokens, i int) Token {
	token := Token{Kind: 0, SubKind: -1, Content: input}

	list := openableTokens.List
	blockquote := openableTokens.Blockquote

	if strings.HasPrefix(input, "#") {
		level := utils.HeadingLevel(input)

		token = Token{Kind: 1, SubKind: level - 1, Content: strings.TrimSpace(input[level:])}
	} else if strings.HasPrefix(input, "-") {
		if !list.IsOpen {
			list.IsOpen = true
			list.Subkind = 0
			list.Index = i
		}

		token = Token{Kind: 2, SubKind: 0, Content: strings.TrimSpace(input[1:])}
	} else if regexp.MustCompile(`[0-9]+.`).MatchString(input) {
		if !list.IsOpen {
			list.IsOpen = true
			list.Subkind = 1
			list.Index = i
		}

		token = Token{Kind: 2, SubKind: 1, Content: strings.TrimSpace(input[2:])}
	} else if strings.HasPrefix(input, ">") {
		if !blockquote.IsOpen {
			blockquote.IsOpen = true
			blockquote.Index = i
		}

		token = Token{Kind: 3, SubKind: -1, Content: strings.TrimSpace(input[1:])}
	} else {
		token = Token{Kind: 0, SubKind: -1, Content: input}
	}

	if list.IsOpen && (token.Kind != 2 || token.SubKind != list.Subkind) {
		list.IsOpen = false
	}

	if blockquote.IsOpen && token.Kind != 3 {
		blockquote.IsOpen = false
	}

	return token
}
