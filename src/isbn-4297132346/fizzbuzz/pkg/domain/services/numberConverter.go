package services

import (
	"github.com/sakuyacatcat/fizzbuzz/pkg/domain/rules"
)

type NumberConverter struct {
	rules []rules.ReplaceRule
}

func NewNumberConverter(rules []rules.ReplaceRule) *NumberConverter {
	return &NumberConverter{rules: rules}
}

func (nc *NumberConverter) Convert(n int) string {
	returnStr := ""

	for _, rule := range nc.rules {
		if rule.Match(returnStr, n) {
			returnStr = rule.Apply(returnStr, n)
		}
	}

	return returnStr
}
