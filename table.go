package latexdown

import "fmt"

var KeywordNodeTable = map[string]NodeFactory{
	`_`: MakeScriptFactory(KindSubscript),
	`^`: MakeScriptFactory(KindSuperscript),

	`$`:       MakeGroupFactory(KindDollarInlineMode), // inline math mode
	`\(`:      MakeGroupFactory(KindRoundInlineMode),  // inline math mode
	`\[`:      MakeGroupFactory(KindDisplayMode),      // display math mode
	`(`:       MakeGroupFactory(KindRoundBracket),
	`[`:       MakeGroupFactory(KindSquareBracket),
	`{`:       MakeGroupFactory(KindCurlyGroup),
	`\{`:      MakeGroupFactory(KindCurlyBracket),
	`|`:       MakeGroupFactory(KindVerticalBar),
	`\|`:      MakeGroupFactory(KindDoublePipe),
	`\langle`: MakeGroupFactory(KindAngleBracket),
	`\lceil`:  MakeGroupFactory(KindCeil),
	`\lfloor`: MakeGroupFactory(KindFloor),

	`\left`:  MakeLeftRightFactory(KindLeft),
	`\right`: MakeLeftRightFactory(KindRight),

	`\big`:   MakeQualifierFactory(KindQualifierbig),
	`\Big`:   MakeQualifierFactory(KindQualifierBig),
	`\bigg`:  MakeQualifierFactory(KindQualifierbigg),
	`\Bigg`:  MakeQualifierFactory(KindQualifierBigg),
	`&`:      MakeQualifierFactory(KindQualifierAlign),
	`\quad`:  MakeQualifierFactory(KindQualifierQuad),
	`\\`:     MakeQualifierFactory(KindQualifierNewline),
	`\ldots`: MakeQualifierFactory(KindQualifierLdots),

	`\verb`: MakeEscapeFactory(KindEscapeVerb),

	`\begin`: MakeBeginEndFactory(KindBegin),

	`\text`:   MakeTextFactory(KindText),
	`\texttt`: MakeTextFactory(KindTexttt),

	`\$`: MakeRawFactory(KindRawDollar),
}

func RegistKeywordNode(keyword string, factory NodeFactory) {
	if factory == nil {
		panic("the node factory is nil")
	}
	if _, exist := KeywordNodeTable[keyword]; exist {
		fmt.Printf("Warning: overwrite keyword %q\n", keyword)
	}
	KeywordNodeTable[keyword] = factory
}
