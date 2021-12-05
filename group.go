package latexdown

import "fmt"

func NewCurlyGroupNode() *Group {
	node := Group{}
	node.kind = KindCurlyGroup
	return &node
}

type Group struct {
	nodeBase
	Values []Noder
}

func (node Group) String() string {
	value := FlattenNodes(node.Values)
	switch node.kind {
	case KindRoundBracket:
		return fmt.Sprintf("(%s)", value)
	case KindSquareBracket:
		return fmt.Sprintf("[%s]", value)
	case KindCurlyBracket, KindCurlyGroup:
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
}
