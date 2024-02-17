package compiler

import (
	"strings"

	"github.com/vclemenzi/markfire/utils"
)

type Token struct {
	Kind    int // 0 = Header
	SubKind int // 0 = H1, 1 = H2, 2 = H3, 3 = H4, 4 = H5, 5 = H6
	Content string
	Line    int
}

func Tokinizer(input string) Token {
	// Heading
	if strings.HasPrefix(input, "#") {
		level := utils.HeadingLevel(input)

		return Token{Kind: 0, SubKind: level, Content: strings.TrimSpace(input[level:]), Line: 0}
	}

	return Token{Kind: -1, SubKind: -1, Content: input, Line: 0}
}
