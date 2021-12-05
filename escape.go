package latexdown

import (
	"fmt"
	"strings"
)

func IsEscapeKind(kind NodeKind) bool {
	switch kind {
	case KindEscapeVerb:
		return true
	}
	return false
}

func MakeEscapeFactory(kind NodeKind) NodeFactory {
	if !IsEscapeKind(kind) {
		panic(fmt.Errorf("MakeEscapeFactory: unsupported kind: %s", kind))
	}
	return func() Noder {
		node := Escape{}
		node.Kind = kind
		return &node
	}
}

type Escape struct {
	NodeBase
	Values []Noder
}

func (node Escape) String() string {
	value := FlattenNodes(node.Values)
	switch node.Kind {
	case KindEscapeVerb:
		return fmt.Sprintf("%s", value)
	default:
		panic(fmt.Errorf("unexpected kind: %s", node.Kind))
	}
}

func (node *Escape) Feed(rd *strings.Reader) {
	ConsumePrefixSpaces(rd)
	n := ParseOneNode(rd, false, func(n Noder) {
		if kind := n.GetKind(); kind != KindVerticalBar {
			panic(fmt.Errorf("unexpected kind of node: %s", kind))
		}
		n.(*Group).SetEscapeFlag()
	}).(*Group)
	node.Values = n.Values
}
