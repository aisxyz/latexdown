package latexdown

import (
	"fmt"
	"strings"
)

func NewSubscriptNode() *Script {
	node := Script{}
	node.kind = KindSubscript
	return &node
}

func NewSuperscriptNode() *Script {
	node := Script{}
	node.kind = KindSuperscript
	return &node
}

type Script struct {
	nodeBase
	Value Noder
}

func (node Script) String() string {
	value := node.Value
	switch node.kind {
	case KindSubscript:
		return fmt.Sprintf("_(%s)", value)
	case KindSuperscript:
		return fmt.Sprintf("^(%s)", value)
	default:
		panic("expect subscript or superscript")
	}
}

func (node *Script) Feed(rd *strings.Reader) {
	for rd.Len() > 0 {
		ch, _, err := rd.ReadRune()
		if err != nil || ch == ' ' {
			continue
		}
		switch ch {
		case '{':
			node.Value = NewCurlyGroupNode()
			node.Value.Feed(rd)
		default:
			node.Value = NewLeafNode(string(ch))
		}
		return
	}
}
