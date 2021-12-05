package latexdown

import (
	"fmt"
	"strings"
)

func IsLeftRightKind(kind NodeKind) bool {
	switch kind {
	case KindLeft, KindRight:
		return true
	}
	return false
}

func MakeLeftRightFactory(kind NodeKind) NodeFactory {
	if !IsLeftRightKind(kind) {
		panic(fmt.Errorf("MakeLeftRightFactory: unsupported kind: %s", kind))
	}
	return func() Noder {
		node := LeftRightGroup{}
		node.Kind = kind
		return &node
	}
}

type LeftRightGroup struct {
	NodeBase
	Value Noder // will be nil for "\left." or "\right*"
}

func (node LeftRightGroup) String() string {
	if node.Value != nil {
		return node.Value.String()
	}
	return ""
}

func (node *LeftRightGroup) Feed(rd *strings.Reader) {
	ConsumePrefixSpaces(rd)
	if rd.Len() > 0 {
		ch, _, err := rd.ReadRune()
		if err != nil || ch == '.' {
			return
		}
		rd.UnreadRune()
		if node.Kind == KindRight {
			return
		}
		node.Value = ParseOneNode(rd, false, nil)
	}
}
