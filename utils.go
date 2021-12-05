package latexdown

import (
	"fmt"
	"strings"
	"unicode"
)

func FlattenNodes(nodes []Noder) string {
	var buf strings.Builder
	for _, node := range nodes {
		buf.WriteString(node.String())
	}
	return buf.String()
}

func ConsumePrefixSpaces(rd *strings.Reader) (consumed bool) {
	for rd.Len() > 0 {
		ch, _, err := rd.ReadRune()
		if err != nil {
			continue
		}
		if ch == ' ' {
			consumed = true
			continue
		}
		rd.UnreadRune()
		return
	}
	return
}

func TrimRightSpaceNodes(nodes []Noder) []Noder {
	end := len(nodes) - 1
	for ; end >= 0; end-- {
		if nodes[end].GetKind() != KindLeafSpace {
			break
		}
	}
	return nodes[:end+1]
}

func Reverse(nodes []Noder) {
	for low, high := 0, len(nodes)-1; low < high; low, high = low+1, high-1 {
		nodes[low], nodes[high] = nodes[high], nodes[low]
	}
}

func IsPunctOrSymbol(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSymbol(r)
}

func Merge(nodes []Noder) []Noder {
	nodes = TrimRightSpaceNodes(nodes)
	result := make([]Noder, 0, len(nodes))

	for end := len(nodes) - 1; end >= 0; end-- {
		currentNode := nodes[end]

		switch kind := currentNode.GetKind(); kind {
		case KindSubscript, KindSuperscript:
			if end < 1 {
				panic(fmt.Errorf("missing one base node"))
			}
			preNode := nodes[end-1]
			for end > 0 && preNode.GetKind() == KindLeafSpace {
				end--
				preNode = nodes[end-1]
			}
			if end < 1 {
				panic(fmt.Errorf("missing one base node"))
			}
			n := Script{
				Base:  preNode,
				Value: currentNode.(*Script).Value,
			}
			n.Kind = kind
			currentNode = &n
			end-- // skip previous node, which is the base of current node.
		case KindLeafText:
			scratch := make([]Noder, 0, 20)
			// Note: this for loop will perform at least once
			for end >= 0 && currentNode.GetKind() == KindLeafText {
				scratch = append(scratch, currentNode)
				end--
				if end < 0 {
					break
				}
				currentNode = nodes[end]
			}
			Reverse(scratch)
			currentNode = NewTextLeafNode(FlattenNodes(scratch))
			end++
		}

		result = append(result, currentNode)
	}

	Reverse(result)
	return result
}
