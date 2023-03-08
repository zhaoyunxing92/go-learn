package v1

import (
	"bytes"
	"strings"
)

type Placeholder struct {
	open  string
	close string
}

func New() *Placeholder {
	return &Placeholder{"${", "}"}
}

func (p *Placeholder) Delims(open, close string) *Placeholder {
	p.open = open
	p.close = close
	return p
}

func (p *Placeholder) Resolver(text string) string {
	return p.resolver(text)
}

func (p *Placeholder) resolver(text string) string {
	var (
		b     bytes.Buffer
		start int
		end   int
	)
	start = strings.Index(text, p.open)
	for start > -1 && start >= end {
		start = start + len(p.open)
		if end = strings.Index(text, p.close); end != -1 {
			placeholder := text[start:end]
			split := strings.Split(placeholder, ":")
			b.WriteString(split[1])
			text = text[end+1:]
			//if strings.Index(text, p.open) > -1 {
			b.WriteString(p.resolver(text))
			//}
		} else {
			start = -1
		}
	}
	return b.String()
}
