package latexdown

type CommandKind int

const (
	KindSubscript = iota
	KindSuperscript

	KindRoundBracket
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
)