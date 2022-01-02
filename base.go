package latexdown

import (
	"fmt"
	"strings"
)

type Noder interface {
	fmt.Stringer
	GetKind() NodeKind
	Feed(*strings.Reader)
}

type NodeFactory func() Noder

type NodeBase struct {
	Kind NodeKind
}

func (node NodeBase) GetKind() NodeKind {
	return node.Kind
}

func (node NodeBase) Feed(rd *strings.Reader) {} // do nothing

func (node NodeBase) String() string {
	return ""
}

var EmptyNode = NodeBase{Kind: KindEmpty}
