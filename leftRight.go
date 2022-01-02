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
		node := LeftRightGroup{Value: EmptyNode}
		node.Kind = kind
		return &node
	}
}

type LeftRightGroup struct {
	NodeBase
	Value Noder
}

func (node LeftRightGroup) String() string {
	return node.Value.String()
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
		ConsumePrefixSpaces(rd)
		node.Value = ParseOneNode(rd, false, func(n Noder) {
			if kind := n.GetKind(); !IsGroupKind(kind) {
				panic(fmt.Errorf("unexpected kind: %s", kind))
			}
		})
	}
}
