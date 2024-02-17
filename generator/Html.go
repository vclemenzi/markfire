package generator

import (
	"fmt"

	"github.com/vclemenzi/markfire/tokenizer"
	"github.com/vclemenzi/markfire/utils"
)

func Html(token tokenizer.Token, previousToken tokenizer.Token, openableTokens *tokenizer.OpenableTokens, i int) string {
	str := ""
	list := openableTokens.List
	blockquote := openableTokens.Blockquote

	if token.Kind == 1 {
		str = fmt.Sprintf("<h%d>%s</h%d>", token.SubKind+1, utils.TextFormat(token.Content), token.SubKind+1)
	} else if token.Kind == 2 {
		if token.SubKind == 0 && list.Index == i {
			str = fmt.Sprintf("<ul><li>%s</li>", utils.TextFormat(token.Content))
		} else if token.SubKind == 1 && list.Index == i {
			str = fmt.Sprintf("<ol><li>%s</li>", utils.TextFormat(token.Content))
		} else {
			str = fmt.Sprintf("<li>%s</li>", utils.TextFormat(token.Content))
		}
	} else if token.Kind == 3 {
		if blockquote.Index == i {
			str = fmt.Sprintf("<blockquote>%s", utils.TextFormat(token.Content))
		} else {
			str = utils.TextFormat(token.Content)
		}

		str += "<br>"
	} else {
		str = utils.TextFormat(token.Content)
	}

	if !list.IsOpen && previousToken.Kind == 2 {
		if list.Subkind == 0 {
			return "</ul>" + str
		} else {
			return "</ol>" + str
		}
	}

	if !blockquote.IsOpen && previousToken.Kind == 3 {
		return "</blockquote>" + str
	}

	return str
}
