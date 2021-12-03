package latexdown

import (
	"fmt"
)

type Commander interface {
	fmt.Stringer
	Kind() CommandKind
}

var CommandNodeTable = map[string]Commander {
	`\(`: nil,        // inline math mode
	`\[`: nil,        // display math mode

	`_`: NewSubscriptCmd(),
	`^`: NewSuperscriptCmd(),

	`(`: nil,
	`[`: nil,
	`|`: nil,
	`\{`: nil,       // -> {x+y}
	`\|`: nil,       // -> ||x+y||
	`\langle`: nil,  // \rangle
	`\lceil`: nil,       // \rceil
	`\lfloor`: nil,       // \rfloor

	`\left`: nil,   // \right
	`\big`: nil,
	`\Big`: nil,
	`\bigg`: nil,
	`\Bigg`: nil,
	`\quad`: nil,  // similar to \t
}