//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_$GOFILE -package=mocks

package rules

type ReplaceRule interface {
	Apply(string, int) string
	Match(string, int) bool
}
