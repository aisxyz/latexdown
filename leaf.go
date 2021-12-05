package latexdown

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func IsLeafKind(kind NodeKind) bool {
	switch kind {
	case KindLeafSpace, KindLeafText,
		KindLeafNumber, KindLeafPunct:
		return true
	}
	return false
}

func NewSpaceLeafNode() *TextLeaf {
	node := TextLeaf{Value: " "}
	node.Kind = KindLeafSpace
	return &node
}

func NewPunctLeafNode(punct rune) *TextLeaf {
	if !IsPunctOrSymbol(punct) {
		panic(fmt.Errorf("%q is not a punctuation character", punct))
	}
	node := TextLeaf{Value: string(punct)}
	node.Kind = KindLeafPunct
	return &node
}

func NewTextLeafNode(text string) *TextLeaf {
	node := TextLeaf{Value: text}
	node.Kind = KindLeafText
	return &node
}

func NewNumberLeafNode() *NumberLeaf {
	node := NumberLeaf{}
	node.Kind = KindLeafNumber
	return &node
}

type TextLeaf struct {
	NodeBase
	Value string
}

func (node TextLeaf) String() string {
	return node.Value
}

type NumberLeaf struct {
	NodeBase
	Value float64
}

func (node NumberLeaf) String() string {
	return fmt.Sprintf("%g", node.Value)
}

func (node *NumberLeaf) Feed(rd *strings.Reader) {
	draft := make([]rune, 0, rd.Len())
	var foundDot bool
	for rd.Len() > 0 {
		ch, _, _ := rd.ReadRune()
		if unicode.IsDigit(ch) {
			draft = append(draft, ch)
		} else if ch == '.' && !foundDot { // handle simple number, e.g. 3.14
			foundDot = true
			draft = append(draft, ch)
		} else {
			rd.UnreadRune()
			break
		}
	}
	val, _ := strconv.ParseFloat(string(draft), 64)
	node.Value = val
}
