package placeholder

import (
	"bytes"
	"strings"
)

type Placeholder struct {
	open  string
	close string
	data  map[string]string
}

func New() *Placeholder {
	return &Placeholder{open: "${", close: "}"}
}

func (p *Placeholder) Delims(open, close string) *Placeholder {
	p.open = open
	p.close = close
	return p
}

func (p *Placeholder) Defaults(data map[string]string) *Placeholder {
	p.data = data
	return p
}

func (p *Placeholder) Resolver(text string) string {
	var (
		bf    bytes.Buffer
		start int
		end   int
	)
	start = strings.Index(text, p.open)
	for start > -1 {
		bf.WriteString(text[:start])
		end = strings.Index(text, p.close)
		if end > -1 {
			exp := text[start : end+len(p.close)]
			bf.WriteString(p.handle(exp))
			text = text[end+len(p.close):]
		} else {
			break
		}
		start = strings.Index(text, p.open)
	}
	if len(text) > 0 {
		bf.WriteString(text)
	}
	return bf.String()
}

func (p *Placeholder) handle(text string) string {
	exp := text[len(p.open) : len(text)-len(p.close)]
	df := strings.Split(exp, ":")
	if len(df) == 2 {
		return df[1]
	}
	if val, ok := p.data[df[0]]; ok {
		return val
	}
	return ""
}
