package xex

import (
	"fmt"
	"strings"
)

//Set up built-in string functions
func registerStringBuiltins() {
	RegisterFunction(
		NewFunction(
			"string",
			`Converts an input into a string using fmt.Sprint`,
			func(in interface{}) string {
				return fmt.Sprint(in)
			},
		),
	)

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

	RegisterFunction(
		NewFunction(
			"len",
			`returns the length of a string`,
			func(in string) int {
				return len(in)
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"substring",
			`returns the substring of the input string from index1 to index2 -1. If index2 is zero, everything to the end of the string is returned`,
			func(input string, index1, index2 int) string {
				if index2 < 1 {
					index2 = len(input)
				}
				return input[index1:index2]
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"instring",
			`returns the start position in the input string of the search string or -1 if the search string is not found`,
			func(input, search string) int {
				return strings.Index(input, search)
			},
		),
	)
}
