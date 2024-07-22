// go:generate mockgen -destination=../../mocks/mock_replaceRule.go -package=mocks isbn-4297132346/fizzbuzz/pkg ReplaceRule

package numberConverter

type ReplaceRule interface {
	Apply(string, int) string
	Match(string, int) bool
}
