package linter

import (
	"fmt"

	"github.com/fatih/color"
)

func Print(errs []LintError, line int) {
	for _, err := range errs {
		color.Yellow(fmt.Sprintf("lint(%d): %s", line, err.Message))
	}
}
