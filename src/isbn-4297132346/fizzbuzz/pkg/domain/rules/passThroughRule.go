package rules

import "strconv"

type PassThroughRule struct{}

func NewPassThroughRule() *PassThroughRule {
	return &PassThroughRule{}
}

func (ptr *PassThroughRule) Match(s string, n int) bool {
	return s == ""
}

func (ptr *PassThroughRule) Apply(s string, n int) string {
	return strconv.FormatInt(int64(n), 10)
}
