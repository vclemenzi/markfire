package tokenizer

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
// Configuration: 4
//   title: 0
//   unknown: -1
// Image: 5
// Codeblock: 6
// Tasklist: 7
//   Unchecked: 0
//   Checked: 1

func Tokenizer(input string, openableTokens *OpenableTokens, i int) Token {
	token := Token{Kind: 0, SubKind: -1, Content: input}

	list := openableTokens.List
	blockquote := openableTokens.Blockquote

	if strings.HasPrefix(input, "#") {
		level := utils.HeadingLevel(input)

		token = Token{Kind: 1, SubKind: level - 1, Content: strings.TrimSpace(input[level:])}
	} else if !strings.HasPrefix(input, "---") && regexp.MustCompile(`^-\s*[^\s\[\]]+\s*$`).MatchString(input) || strings.HasPrefix(input, "*") { // This regex is so long because it needs to check if it is not a checkbox
		if !list.IsOpen {
			list.IsOpen = true
			list.Subkind = 0
			list.Index = i
		}

		token = Token{Kind: 2, SubKind: 0, Content: strings.TrimSpace(input[1:])}
	} else if regexp.MustCompile(`^[0-9]+\.`).MatchString(input) {
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
	} else if strings.HasPrefix(input, "---") {
		openableTokens.Configuration.IsOpen = !openableTokens.Configuration.IsOpen

		token = Token{Kind: 4, SubKind: -1, Content: strings.TrimSpace(input[3:])}
	} else if openableTokens.Configuration.IsOpen && strings.HasPrefix(input, "title:") {
		token = Token{Kind: 4, SubKind: 0, Content: strings.TrimSpace(input[6:])}
	} else if regexp.MustCompile(`!\[.*\]\(.*\)`).MatchString(input) {
		// alt:token
		alt := regexp.MustCompile(`!\[(.*)\]`).FindStringSubmatch(input)[1]
		url := regexp.MustCompile(`\((.*)\)`).FindStringSubmatch(input)[1]

		token = Token{Kind: 5, SubKind: -1, Content: alt + ":" + url}
	} else if strings.HasPrefix(input, "```") {
		openableTokens.Codeblock.IsOpen = !openableTokens.Codeblock.IsOpen

		token = Token{Kind: 6, SubKind: -1, Content: ""}
	} else if openableTokens.Codeblock.IsOpen {
		token = Token{Kind: 6, SubKind: 0, Content: input}
	} else if regexp.MustCompile(`^- \[x\]`).MatchString(input) {
		if !openableTokens.Tasklist.IsOpen {
			openableTokens.Tasklist.IsOpen = true
			openableTokens.Tasklist.Index = i
		}

		token = Token{Kind: 7, SubKind: 1, Content: strings.TrimSpace(input[5:])}
	} else if regexp.MustCompile(`^- \[ \]`).MatchString(input) {
		if !openableTokens.Tasklist.IsOpen {
			openableTokens.Tasklist.IsOpen = true
			openableTokens.Tasklist.Index = i
		}

		token = Token{Kind: 7, SubKind: 0, Content: strings.TrimSpace(input[5:])}
	} else {
		token = Token{Kind: 0, SubKind: -1, Content: input}
	}

	if list.IsOpen && (token.Kind != 2 || token.SubKind != list.Subkind) {
		list.IsOpen = false
	}

	if blockquote.IsOpen && token.Kind != 3 {
		blockquote.IsOpen = false
	}

	if openableTokens.Tasklist.IsOpen && token.Kind != 7 {
		openableTokens.Tasklist.IsOpen = false
	}

	return token
}
