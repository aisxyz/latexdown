package latexdown

import (
	"fmt"
	"io"
	"strings"
)

func IsBeginEndKind(kind NodeKind) bool {
	switch kind {
	case KindBegin:
		return true
	}
	return false
}

func MakeBeginEndFactory(kind NodeKind) NodeFactory {
	if !IsBeginEndKind(kind) {
		panic(fmt.Errorf("MakeBeginEndFactory: unsupported kind: %s", kind))
	}
	return func() Noder {
		node := BeginEndGroup{}
		node.Kind = kind
		return &node
	}
}

type BeginEndGroup struct {
	NodeBase
	Values  []Noder
	Command string
}

func (node BeginEndGroup) String() string {
	value := FlattenNodes(node.Values)
	switch node.Command {
	case "matrix":
		return fmt.Sprintf("\n%s\n", value)
	case "pmatrix":
		return fmt.Sprintf("\n(%s)\n", value)
	case "bmatrix":
		return fmt.Sprintf("\n[%s]\n", value)
	case "Bmatrix":
		return fmt.Sprintf("\n{%s}\n", value)
	case "vmatrix":
		return fmt.Sprintf("\n|%s|\n", value)
	case "Vmatrix":
		return fmt.Sprintf("\n||%s||\n", value)
	default:
		panic(fmt.Errorf("unexpected command: %s", node.Command))
	}
}

func (node *BeginEndGroup) Feed(rd *strings.Reader) {
	ConsumePrefixSpaces(rd)
	n := ParseOneNode(rd, false, func(n Noder) {
		if kind := n.GetKind(); kind != KindCurlyGroup {
			panic(fmt.Errorf("unexpected command kind: %s", kind))
		}
	}).(*Group)

	node.Command = FlattenNodes(n.Values)

	if len(node.Command) == 0 {
		panic("missing command node")
	}

	const terminate = `\end`
	scratch := make([]byte, len(terminate))

	for rd.Len() > 0 {
		n, err := rd.Read(scratch)
		if err != nil {
			rd.Seek(int64(1-n), io.SeekCurrent) // advance 1 byte
			continue
		}

		if string(scratch[:n]) == terminate {
			ConsumePrefixSpaces(rd)
			n := ParseOneNode(rd, false, func(n Noder) {
				if kind := n.GetKind(); kind != KindCurlyGroup {
					panic(fmt.Errorf("unexpected command kind: %s", kind))
				}
			}).(*Group)

			endCmd := FlattenNodes(n.Values)
			if endCmd != node.Command {
				panic(fmt.Errorf("unexpected end command: %s, want: %s", endCmd, node.Command))
			}
			break
		}

		rd.Seek(-int64(n), io.SeekCurrent)
		nd := ParseOneNode(rd, false, nil)
		node.Values = append(node.Values, nd)
	}

	node.Values = Merge(node.Values)
}
