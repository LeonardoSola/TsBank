package tools

import (
	"regexp"
)

var nonAlphanumericRegex = regexp.MustCompile(`\D`)

func RemoveNonNumbers(str string) (numbers string) {
	numbers = nonAlphanumericRegex.ReplaceAllString(str, "")
	return numbers
}
