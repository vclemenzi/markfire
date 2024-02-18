package utils

import (
	"regexp"

	"github.com/enescakir/emoji"
)

func TextFormat(input string) string {
	bold := regexp.MustCompile(`\*\*(.*?)\*\*`)
	italic := regexp.MustCompile(`\*(.*?)\*`)
	code := regexp.MustCompile("`([^`]+)`")
	url := regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	highlight := regexp.MustCompile(`==([^=]+)==`)
	strikethrough := regexp.MustCompile(`~~([^~]+)~~`)

	input = bold.ReplaceAllString(input, "<b>$1</b>")
	input = italic.ReplaceAllString(input, "<i>$1</i>")
	input = code.ReplaceAllString(input, "<code>$1</code>")
	input = url.ReplaceAllString(input, "<a href='$2'>$1</a>")
	input = highlight.ReplaceAllString(input, "<mark>$1</mark>")
	input = strikethrough.ReplaceAllString(input, "<s>$1</s>")

	return emoji.Parse(input)
}
