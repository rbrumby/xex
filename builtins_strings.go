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
			FunctionDocumentation{
				Text: `Converts an input into a string using fmt.Sprint`,
				Parameters: map[string]string{
					"in": "The value to convert to a string.",
				},
			},
			func(in interface{}) string {
				return fmt.Sprint(in)
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"concat",
			FunctionDocumentation{
				Text: `concatenates any number of strings returning a single string result`,
				Parameters: map[string]string{
					"strs": "variadic - the strings to concatentate.",
				},
			},
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
			FunctionDocumentation{
				Text: `returns the length of a string`,
				Parameters: map[string]string{
					"in": "The string to measure.",
				},
			},
			func(in string) int {
				return len(in)
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"substring",
			FunctionDocumentation{
				Text: `returns the substring of the input string from index1 to index2 -1. If index2 is zero, everything to the end of the string is returned`,
				Parameters: map[string]string{
					"input": "The string take take a substring from.",
					"start": "The start index (counting from 0).",
					"end":   "The end index. If this is less than 1, defaults to the end of the string.",
				},
			},
			func(input string, start, end int) string {
				if end < 1 {
					end = len(input)
				}
				return input[start:end]
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"instring",
			FunctionDocumentation{
				Text: `returns the start position in the input string of the search string or -1 if the search string is not found`,
				Parameters: map[string]string{
					"input":  "The string to search.",
					"search": "The string to find in the input.",
				},
			},
			func(input, search string) int {
				return strings.Index(input, search)
			},
		),
	)
}
