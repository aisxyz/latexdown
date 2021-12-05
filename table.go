package latexdown

var KeywordNodeTable = map[string]Noder{
	`\(`: nil, // inline math mode
	`\[`: nil, // display math mode

	`_`: NewSubscriptNode(),
	`^`: NewSuperscriptNode(),

	`(`: nil,
	//`{`:       nil,
	`[`:       nil,
	`|`:       nil,
	`\{`:      nil, // -> {x+y}
	`\|`:      nil, // -> ||x+y||
	`\langle`: nil, // \rangle
	`\lceil`:  nil, // \rceil
	`\lfloor`: nil, // \rfloor

	`\left`: nil, // \right
	`\big`:  nil,
	`\Big`:  nil,
	`\bigg`: nil,
	`\Bigg`: nil,
	`\quad`: nil, // similar to \t
}
