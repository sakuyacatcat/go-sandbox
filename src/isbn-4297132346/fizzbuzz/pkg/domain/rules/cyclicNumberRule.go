package rules

import "fmt"

type CyclicNumberRule struct {
	baseNum    int
	replaceStr string
}

func NewCyclicNumberRule(baseNum int, replaceStr string) *CyclicNumberRule {
	return &CyclicNumberRule{baseNum: baseNum, replaceStr: replaceStr}
}

func (c *CyclicNumberRule) Match(s string, n int) bool {
	return n%c.baseNum == 0
}

func (c *CyclicNumberRule) Apply(s string, n int) string {
	return fmt.Sprintf("%s%s", s, c.replaceStr)
}
