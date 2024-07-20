package numberConverter

import "strings"

type NumberConverter struct {
	rules []ReplaceRule;
}

func NewNumberConverter(rules []ReplaceRule) *NumberConverter {
	return &NumberConverter{rules: rules};
}

func (nc *NumberConverter) Convert(n int) string {
	var returnStrs []string;

	for _, rule := range nc.rules {
		result := rule.Replace(n);
		if result != "" {
			returnStrs = append(returnStrs, result);
		}
	}

	return strings.Join(returnStrs, "");
}
