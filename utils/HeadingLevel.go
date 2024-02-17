package utils

import (
	"regexp"
)

func HeadingLevel(input string) int {
	regex := regexp.MustCompile(`^#+`)
	level := regex.FindString(input)

	return len(level)
}
