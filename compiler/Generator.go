package compiler

import "fmt"

func Generate(token Token, previousToken Token, list *List, i int) string {
	str := ""

	if token.Kind == 1 {
		str = fmt.Sprintf("<h%d>%s</h%d>", token.SubKind+1, token.Content, token.SubKind+1)
	} else if token.Kind == 2 {
		if token.SubKind == 0 && list.Index == i {
			str = fmt.Sprintf("<ul><li>%s</li>", token.Content)
		} else if token.SubKind == 1 && list.Index == i {
			str = fmt.Sprintf("<ol><li>%s</li>", token.Content)
		} else {
			str = fmt.Sprintf("<li>%s</li>", token.Content)
		}
	} else {
		str = token.Content
	}

	if !list.IsOpen && previousToken.Kind == 2 {
		if list.Subkind == 0 {
			return "</ul>" + str
		} else {
			return "</ol>" + str
		}
	}

	return str
}
