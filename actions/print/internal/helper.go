package internal

import (
	"strings"
)

// TimeExpression convert a string expression into a valid go timestamp format.
func TimeExpression(value string) string {
	value = strings.ReplaceAll(value, "yyyy", "2006")
	value = strings.ReplaceAll(value, "MM", "01")
	value = strings.ReplaceAll(value, "dd", "02")
	value = strings.ReplaceAll(value, "HH", "15")
	value = strings.ReplaceAll(value, "hh", "3")
	value = strings.ReplaceAll(value, "mm", "04")
	value = strings.ReplaceAll(value, "ss", "05")

	return value
}
