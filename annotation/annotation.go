package annotation

import (
	"fmt"
	"strings"
	"text/scanner"
)

const (
	initial = iota
	annotationName
	attributeName
	attributeValue
	done
)

type Annotation struct {
	Name       string
	Attributes map[string]string
}

// ParseAnnotation 解析注解字符串
func ParseAnnotation(line string) (Annotation, error) {
	withoutComment := strings.TrimSpace(strings.TrimLeft(line, "/"))

	annotation := Annotation{
		Name:       "",
		Attributes: make(map[string]string),
	}

	var s scanner.Scanner
	s.Init(strings.NewReader(withoutComment))

	var tok rune
	currentStatus := initial
	var attrName string

	for tok != scanner.EOF && currentStatus < done {
		tok = s.Scan()
		switch tok {
		case '@':
			currentStatus = annotationName
		case '(':
			currentStatus = attributeName
		case '=':
			currentStatus = attributeValue
		case ',':
			currentStatus = attributeName
		case ')':
			currentStatus = done
		case scanner.Ident:
			switch currentStatus {
			case annotationName:
				annotation.Name = s.TokenText()
			case attributeName:
				attrName = s.TokenText()
			}
		default:
			switch currentStatus {
			case attributeValue:
				annotation.Attributes[attrName] = strings.Trim(s.TokenText(), "\"")
			}
		}
	}

	if currentStatus != done {
		return annotation, fmt.Errorf("invalid completion status %v for annotation:%s", currentStatus, line)
	}
	return annotation, nil
}
