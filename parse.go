package latexdown

import (
	"fmt"
	"strings"
	"unicode"
)

var Debug bool

func Parse(latex string) (nodes []Noder) {
	rd := strings.NewReader(strings.TrimSpace(latex))
	for rd.Len() > 0 {
		node := ParseOneNode(rd, false, nil)
		nodes = append(nodes, node)
	}
	return
}

// ParseOneNode returns one node, which will perform beforeFeed (if set) before caling Feed.
func ParseOneNode(rd *strings.Reader, escaped bool, beforeFeed func(n Noder)) (node Noder) {
	defer func() {
		if beforeFeed != nil {
			beforeFeed(node)
		}
		node.Feed(rd)

		if Debug {
			fmt.Printf("[*Debug] Node(kind=%q, value=%q)\n", node.GetKind(), node)
		}
	}()

	if rd.Len() == 0 {
		panic("no more input")
	}
	consumed := ConsumePrefixSpaces(rd)
	if consumed {
		return NewSpaceLeafNode()
	}
	ch, _, _ := rd.ReadRune()
	if unicode.IsDigit(ch) {
		rd.UnreadRune()
		return NewNumberLeafNode()
	}
	key := string(ch)
	isPunct := IsPunctOrSymbol(ch)
	if escaped {
		if isPunct {
			return NewPunctLeafNode(ch)
		}
		return NewTextLeafNode(key)
	}
	if ch == '\\' {
		key += parseLatexKeyword(rd)
	} else if isPunct && KeywordNodeTable[key] == nil {
		return NewPunctLeafNode(ch)
	}
	factory := KeywordNodeTable[key]
	if factory != nil {
		return factory()
	}
	return NewTextLeafNode(key)
}

func parseLatexKeyword(rd *strings.Reader) string {
	var foundLetter bool
	scratch := make([]rune, 0, 20)
	for rd.Len() > 0 {
		ch, _, err := rd.ReadRune()
		if err != nil || ch == ' ' {
			break
		}
		if !unicode.IsLetter(ch) {
			if !foundLetter {
				return string(ch)
			}
			rd.UnreadRune()
			break
		}
		foundLetter = true
		scratch = append(scratch, ch)
	}
	return string(scratch)
}
