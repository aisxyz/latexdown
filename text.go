package latexdown

import (
	"fmt"
	"strings"
)

func IsTextKind(kind NodeKind) bool {
	switch kind {
	case KindText, KindTexttt:
		return true
	}
	return false
}

func MakeTextFactory(kind NodeKind) NodeFactory {
	if !IsTextKind(kind) {
		panic(fmt.Errorf("MakeTextFactory: unsupported kind: %s", kind))
	}
	return func() Noder {
		node := Text{}
		node.Kind = kind
		return &node
	}
}

type Text struct {
	NodeBase
	Values []Noder
}

func (node Text) String() string {
	value := FlattenNodes(node.Values)
	switch node.Kind {
	case KindText, KindTexttt:
		return fmt.Sprintf("%s", value)
	default:
		panic(fmt.Errorf("unexpected kind: %s", node.Kind))
	}
}

func (node *Text) Feed(rd *strings.Reader) {
	ConsumePrefixSpaces(rd)
	n := ParseOneNode(rd, false, func(n Noder) {
		if kind := n.GetKind(); kind != KindCurlyGroup {
			panic(fmt.Errorf("unexpected kind of node: %s", kind))
		}
	}).(*Group)
	node.Values = n.Values
}
