package latexdown

import (
	"testing"
)

func TestParse(t *testing.T) {
	//const latex = `\( (2a^n)^{r+s} = a^{nr+ ns} \) and \[\langle 3x+7 \rangle\] or use \verb|\(...\)| also \texttt{\$...\$}`
	//const latex = `y  = 1 + & \left(  \frac{1}{x} + \frac{1}{x^2} + \frac{1}{x^3} + \ldots \right. \\ &\left. \quad + \frac{1}{x^{n-1}} + \frac{1}{x^n} \right)`
	//const latex = `F = G \left( \frac{m_1 m_2}{r^2} \right)`
	//const latex = `\left\lceil \begin{matrix} 1 & 2 & 3\\a & b & c \end{matrix} \right\rceil`
	const latex = `\begin{bmatrix} 1 & 2 & 3\\a & b & c \end{bmatrix}`
	Debug = true
	//nodes := Parse(latex)
	nodes := Merge(Parse(latex))
	for _, node := range nodes {
		t.Logf("%+v\n", node)
	}
	t.Logf("latex: %s\n", latex)
	t.Logf("transform: %s\n", FlattenNodes(nodes))
}
