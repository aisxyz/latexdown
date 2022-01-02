package latexdown

import (
	"fmt"
	"io"
	"strings"
)

func IsGroupKind(kind NodeKind) bool {
	switch kind {
	case KindRoundInlineMode, KindDollarInlineMode,
		KindDisplayMode, KindCurlyGroup,
		KindRoundBracket, KindSquareBracket,
		KindCurlyBracket, KindAngleBracket,
		KindVerticalBar, KindDoublePipe,
		KindCeil, KindFloor:
		return true
	}
	return false
}

func MakeGroupFactory(kind NodeKind) NodeFactory {
	if !IsGroupKind(kind) {
		panic(fmt.Errorf("MakeGroupFactory: unsupported kind: %s", kind))
	}
	return func() Noder {
		node := Group{}
		node.Kind = kind
		return &node
	}
}

type Group struct {
	NodeBase
	Values []Noder
	Closed bool

	escape bool
}

func (node *Group) SetEscapeFlag() {
	node.escape = true
}

func (node Group) String() string {
	value := FlattenNodes(node.Values)
	switch node.Kind {
	case KindRoundInlineMode, KindDollarInlineMode:
		return fmt.Sprintf("%s", value)
	case KindDisplayMode:
		return fmt.Sprintf("\n\t\t%s\n", value)
	case KindRoundBracket, KindCurlyGroup:
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
		return fmt.Sprintf("⌈%s⌉", value)
	case KindFloor:
		return fmt.Sprintf("⌊%s⌋", value)
	default:
		panic(fmt.Errorf("unexpected kind: %s", node.Kind))
	}
}

func (node *Group) Feed(rd *strings.Reader) {
	terminate := node.getTerminateStr()
	scratch := make([]byte, len(terminate))
	for rd.Len() > 0 {
		n, err := rd.Read(scratch)
		if err != nil {
			rd.Seek(int64(1-n), io.SeekCurrent) // advance 1 byte
			continue
		}
		if string(scratch[:n]) == terminate {
			node.Closed = true
			break
		}
		rd.Seek(-int64(n), io.SeekCurrent)
		nd := ParseOneNode(rd, node.escape, nil)
		node.Values = append(node.Values, nd)
	}
	node.Values = Merge(node.Values)
}

func (node Group) getTerminateStr() string {
	switch node.Kind {
	case KindRoundInlineMode:
		return `\)`
	case KindDollarInlineMode:
		return "$"
	case KindDisplayMode:
		return `\]`
	case KindRoundBracket:
		return ")"
	case KindSquareBracket:
		return "]"
	case KindCurlyBracket:
		return `\}`
	case KindCurlyGroup:
		return "}"
	case KindAngleBracket:
		return `\rangle`
	case KindVerticalBar:
		return "|"
	case KindDoublePipe:
		return `\|`
	case KindCeil:
		return `\rceil`
	case KindFloor:
		return `\rfloor`
	default:
		panic(fmt.Errorf("unexpected kind: %s", node.Kind))
	}
}
