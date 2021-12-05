package latexdown

type NodeKind int

const (
	KindSubscript NodeKind = iota
	KindSuperscript

	KindRoundBracket
	KindCurlyGroup
	KindSquareBracket
	KindCurlyBracket
	KindAngleBracket
	KindVerticalBar
	KindDoublePipe
	KindCeil
	KindFloor

	KindQualifierLeft
	KindQualifierbig
	KindQualifierBig
	KindQualifierbigg
	KindQualifierBigg

	KindLeaf
	KindLeafText
	KindLeafNumber
)
