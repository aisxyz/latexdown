package latexdown

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func NewLeafNode(content string) *Leaf {
	node := Leaf{}
	node.kind = KindLeaf
	node.Values = ParseLeafNodes(content)
	return &node
}

type Leaf struct {
	nodeBase
	Values []Noder
}

func (node Leaf) String() string {
	return FlattenNodes(node.Values)
}

func ParseLeafNodes(content string) (nodes []Noder) {
	rd := strings.NewReader(content)
	for rd.Len() > 0 {
		ch, _, _ := rd.ReadRune()
		rd.UnreadRune()

		var node Noder
		if unicode.IsDigit(ch) {
			node = newNumberLeafNode()
		} else {
			node = newTextLeafNode()
		}
		node.Feed(rd)
		nodes = append(nodes, node)
	}
	return
}

func newTextLeafNode() *TextLeaf { // private
	node := TextLeaf{}
	node.kind = KindLeafText
	return &node
}

func newNumberLeafNode() *NumberLeaf { // private
	node := NumberLeaf{}
	node.kind = KindLeafNumber
	return &node
}

type TextLeaf struct {
	nodeBase
	Value string
}

func (node TextLeaf) String() string {
	return node.Value
}

func (node *TextLeaf) Feed(rd *strings.Reader) {
	draft := make([]rune, 0, rd.Len())
	for rd.Len() > 0 {
		ch, _, _ := rd.ReadRune()
		if unicode.IsDigit(ch) {
			rd.UnreadRune()
			break
		} else {
			draft = append(draft, ch)
		}
	}
	node.Value = string(draft)
}

type NumberLeaf struct {
	nodeBase
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
