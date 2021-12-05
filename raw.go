package latexdown

import (
	"fmt"
)

func IsRawKind(kind NodeKind) bool {
	switch kind {
	case KindRawDollar:
		return true
	}
	return false
}

func MakeRawFactory(kind NodeKind) NodeFactory {
	if !IsRawKind(kind) {
		panic(fmt.Errorf("MakeRawFactory: unsupported kind: %s", kind))
	}
	return func() Noder {
		node := Raw{}
		node.Kind = kind
		return &node
	}
}

type Raw struct {
	NodeBase
}

func (node Raw) String() string {
	switch node.Kind {
	case KindRawDollar:
		return "$"
	default:
		panic(fmt.Errorf("unexpected kind: %s", node.Kind))
	}
}
