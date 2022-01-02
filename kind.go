package latexdown

import "fmt"

type NodeKind int

const (
	KindEmpty NodeKind = iota

	KindSubscript
	KindSuperscript

	KindRoundInlineMode
	KindDollarInlineMode
	KindDisplayMode
	KindCurlyGroup
	KindRoundBracket
	KindSquareBracket
	KindCurlyBracket
	KindAngleBracket
	KindVerticalBar
	KindDoublePipe
	KindCeil
	KindFloor

	KindLeft
	KindRight

	KindQualifierbig
	KindQualifierBig
	KindQualifierbigg
	KindQualifierBigg
	KindQualifierAlign
	KindQualifierQuad
	KindQualifierNewline
	KindQualifierLdots

	KindEscapeVerb

	KindBegin

	KindText
	KindTexttt

	KindRawDollar

	KindLeafText
	KindLeafNumber
	KindLeafSpace
	KindLeafPunct
)

func (kind NodeKind) String() string {
	switch kind {
	case KindEmpty:
		return "KindEmpty"
	case KindSubscript:
		return "KindSubscript"
	case KindSuperscript:
		return "KindSuperscript"
	case KindRoundInlineMode:
		return "KindRoundInlineMode"
	case KindDollarInlineMode:
		return "KindDollarInlineMode"
	case KindDisplayMode:
		return "KindDisplayMode"
	case KindCurlyGroup:
		return "KindCurlyGroup"
	case KindRoundBracket:
		return "KindRoundBracket"
	case KindSquareBracket:
		return "KindSquareBracket"
	case KindCurlyBracket:
		return "KindCurlyBracket"
	case KindAngleBracket:
		return "KindAngleBracket"
	case KindVerticalBar:
		return "KindVerticalBar"
	case KindDoublePipe:
		return "KindDoublePipe"
	case KindCeil:
		return "KindCeil"
	case KindFloor:
		return "KindFloor"
	case KindLeft:
		return "KindLeft"
	case KindRight:
		return "KindRight"
	case KindQualifierbig:
		return "KindQualifierbig"
	case KindQualifierBig:
		return "KindQualifierBig"
	case KindQualifierbigg:
		return "KindQualifierbigg"
	case KindQualifierBigg:
		return "KindQualifierBigg"
	case KindQualifierAlign:
		return "KindQualifierAlign"
	case KindQualifierQuad:
		return "KindQualifierQuad"
	case KindQualifierNewline:
		return "KindQualifierNewline"
	case KindQualifierLdots:
		return "KindQualifierLdots"
	case KindEscapeVerb:
		return "KindEscapeVerb"
	case KindBegin:
		return "KindBegin"
	case KindText:
		return "KindText"
	case KindTexttt:
		return "KindTexttt"
	case KindRawDollar:
		return "KindRawDollar"
	case KindLeafText:
		return "KindLeafText"
	case KindLeafNumber:
		return "KindLeafNumber"
	case KindLeafSpace:
		return "KindLeafSpace"
	case KindLeafPunct:
		return "KindLeafPunct"
	default:
		panic(fmt.Errorf("unknown kind: %d", kind))
	}
}
