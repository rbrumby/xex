// +build documentation

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/rbrumby/xex"
)

func main() {
	out, err := os.Create("../builtins.md")
	if err != nil {
		panic(err)
	}
	out.WriteString("| function | args | description\n")
	out.WriteString("| -------- | ---- | -----------\n")
	keys := xex.GetFunctionNames()
	sort.Strings(keys)
	for _, k := range keys {
		fn, _ := xex.GetFunction(k)
		out.WriteString(fmt.Sprintf("| %s |", fn.Name))
		for pi, pv := range fn.Documentation.Parameters {
			out.WriteString(fmt.Sprintf("[%d] %s: %s<br/>", pi, pv.Name, strings.ReplaceAll(pv.Description, "\n", " ")))
		}
		out.WriteString("| " + strings.ReplaceAll(fn.Documentation.Text, "\n", " ") + "|\n")
	}
}
