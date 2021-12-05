package latexdown

import "fmt"

func IsQualifierKind(kind NodeKind) bool {
	switch kind {
	case KindQualifierbig, KindQualifierBig,
		KindQualifierbigg, KindQualifierBigg,
		KindQualifierAlign, KindQualifierQuad,
		KindQualifierNewline, KindQualifierLdots:
		return true
	}
	return false
}

func MakeQualifierFactory(kind NodeKind) NodeFactory {
	if !IsQualifierKind(kind) {
		panic(fmt.Errorf("MakeQualifierFactory: unsupported kind: %s", kind))
	}
	return func() Noder {
		node := Qualifier{}
		node.Kind = kind
		return &node
	}
}

type Qualifier struct {
	NodeBase
}

func (node Qualifier) String() string {
	switch node.Kind {
	case KindQualifierQuad:
		return "\t"
	case KindQualifierNewline:
		return "\n"
	case KindQualifierLdots:
		return "..."
	default:
		return ""
	}
}
