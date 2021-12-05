package latexdown

import "strings"

func FlattenNodes(nodes []Noder) string {
	var buf strings.Builder
	for _, node := range nodes {
		buf.WriteString(node.String())
	}
	return buf.String()
}
