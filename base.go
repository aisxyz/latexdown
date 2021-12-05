package latexdown

import (
	"fmt"
	"strings"
)

type Noder interface {
	fmt.Stringer
	Kind() NodeKind
	Feed(*strings.Reader)
}

type nodeBase struct {
	kind NodeKind
}

func (node nodeBase) Kind() NodeKind {
	return node.kind
}

func (node nodeBase) Feed(rd *strings.Reader) {} // do nothing

/*
func (node Script) String() string {
	value := FlattenNodes(node.Value)
	switch node.kind {
	case KindSubscript:
		return fmt.Sprintf("_(%s)", value)
	case KindSuperscript:
		return fmt.Sprintf("^(%s)", value)
	case KindRoundBracket:
		return fmt.Sprintf("(%s)", value)
	case KindSquareBracket:
		return fmt.Sprintf("[%s]", value)
	case KindCurlyBracket:
		return fmt.Sprintf("{%s}", value)
	case KindAngleBracket:
		return fmt.Sprintf("<%s>", value)
	case KindVerticalBar:
		return fmt.Sprintf("|%s|", value)
	case KindDoublePipe:
		return fmt.Sprintf("||%s||", value)
	case KindCeil:
		return fmt.Sprintf("CEIL INTEGER OF (%s)", value)
	case KindFloor:
		return fmt.Sprintf("FLOOR INTEGER OF (%s)", value)
	default:
		panic("expect subscript or superscript")
	}
}*/
