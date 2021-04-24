package xex

import (
	"strings"
)

//Set up built-in string functions
func registerStringBuiltins() {
	RegisterFunction(
		NewFunction(
			"concat",
			`concatenates any number of strings returning a single string result`,
			func(strs ...string) string {
				sb := strings.Builder{}
				for _, s := range strs {
					sb.WriteString(s)
				}
				return sb.String()
			},
		),
	)
}
