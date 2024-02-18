package linter

import (
	"strings"

	"github.com/vclemenzi/markfire/tokenizer"
)

func Lint(token tokenizer.Token) []LintError {
	var errors []LintError

	if token.Kind == 0 {
		if len(token.Content) > 75 {
			errors = append(errors, LintError{Message: "Line has more than 75 characters", Token: token})
		} else if strings.HasSuffix(token.Content, " ") {
			errors = append(errors, LintError{Message: "Line ends with a space", Token: token})
		}
	}

	return errors
}
