package utils

import "regexp"

func TextFormat(input string) string {
	bold := regexp.MustCompile(`\*\*(.*?)\*\*`)
	italic := regexp.MustCompile(`\*(.*?)\*`)
	code := regexp.MustCompile("`([^`]+)`")

	input = bold.ReplaceAllString(input, "<b>$1</b>")
	input = italic.ReplaceAllString(input, "<i>$1</i>")
	input = code.ReplaceAllString(input, "<code>$1</code>")

	return input
}