package latexdown

import "fmt"

type cmdBase struct {
	kind CommandKind
}

func (cmd cmdBase) Kind() CommandKind {
	return cmd.kind
}


type CmdScript struct {
	cmdBase
	Value []Commander
}

func (cmd CmdScript) String() string {
	value := FlattenCommands(cmd.Value)
	switch cmd.kind {
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
}

func NewSubscriptCmd() *CmdScript {
	cmd := CmdScript{}
	cmd.kind = KindSubscript
	return &cmd
}

func NewSuperscriptCmd() *CmdScript {
	cmd := CmdScript{}
	cmd.kind = KindSuperscript
	return &cmd
}