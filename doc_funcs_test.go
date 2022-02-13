// +build documentation

package xex

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
)

func TestWriteBuiltinFunctionsToMDTable(t *testing.T) {
	out, err := os.Create("./builtins.md")
	if err != nil {
		t.Error(err)
	}
	out.WriteString("| function | args | description\n")
	out.WriteString("| -------- | ---- | -----------\n")
	keys := make([]string, 0, len(functions))
	for k := range functions {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		out.WriteString(fmt.Sprintf("| %s |", functions[k].name))
		for pk, pv := range functions[k].documentation.Parameters {
			out.WriteString(fmt.Sprintf("> %s: %s<br/>", pk, strings.ReplaceAll(pv, "\n", " ")))
		}
		out.WriteString("| " + strings.ReplaceAll(functions[k].documentation.Text, "\n", " ") + "|\n")
	}
}
