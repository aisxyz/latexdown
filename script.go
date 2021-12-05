package latexdown

import (
	"fmt"
	"strings"
)

func IsScriptKind(kind NodeKind) bool {
	switch kind {
	case KindSubscript, KindSuperscript:
		return true
	}
	return false
}

func MakeScriptFactory(kind NodeKind) NodeFactory {
	if !IsScriptKind(kind) {
		panic(fmt.Errorf("MakeScriptFactory: unsupported kind: %s", kind))
	}
	return func() Noder {
		node := Script{}
		node.Kind = kind
		return &node
	}
}

type Script struct {
	NodeBase
	Value Noder
	Base  Noder // only set while merging nodes
}

func (node Script) String() string {
	var base string
	if node.Base != nil {
		base = node.Base.String()
	}

	switch node.Kind {
	case KindSubscript:
		return fmt.Sprintf("%s_%s", base, node.Value)
	case KindSuperscript:
		return fmt.Sprintf("%s^%s", base, node.Value)
	default:
		panic(fmt.Errorf("unexpected kind: %s", node.Kind))
	}
}

func (node *Script) Feed(rd *strings.Reader) {
	n := ParseOneNode(rd, false, func(n Noder) {
		kind := n.GetKind()
		if kind != KindCurlyGroup && !IsLeafKind(kind) {
			panic(fmt.Errorf("unexpected kind of node: %s", kind))
		}
	})
	node.Value = n
}
