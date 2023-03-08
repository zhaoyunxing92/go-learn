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
	if len(text) == 0 {
		return ""
	}
	return p.resolver(text)
}

func (p *Placeholder) resolver(text string) string {
	var (
		b     bytes.Buffer
		start int
	)
	if len(text) == 0 {
		return b.String()
	}
	start = strings.Index(text, p.open)
	for start > -1 {
		start = start + len(p.open)
		if end := strings.Index(text, p.close); end != -1 {
			placeholder := text[start:end]
			split := strings.Split(placeholder, ":")
			b.WriteString(split[1])
			text = text[end+1:]
			p.resolver(text)
		} else {
			start = -1
		}
	}
	return b.String()
}
