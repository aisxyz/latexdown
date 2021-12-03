package latexdown

import "strings"

func FlattenCommands(cmds []Commander) string {
	var buf strings.Builder
	for _, cmd := range cmds {
		buf.WriteString(cmd.String())
	}
	return buf.String()
}