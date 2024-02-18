package linter

import "github.com/vclemenzi/markfire/tokenizer"

type LintError struct {
	Message string
	Token   tokenizer.Token
}
