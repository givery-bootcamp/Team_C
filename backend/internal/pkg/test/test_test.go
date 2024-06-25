package test_test

import (
	"myapp/internal/pkg/test"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestDiffEq(t *testing.T) {
	matcher := test.DiffEq("test")
	assert.NotNil(t, matcher)
}

func TestDiffMatcherMatches(t *testing.T) {
	t.Run("差分がない場合、Matchesはtrueを返す", func(t *testing.T) {
		matcher := test.DiffEq("test")
		assert.True(t, matcher.Matches("test"))
	})

	t.Run("差分がある場合、Matchesはfalseを返す", func(t *testing.T) {
		matcher := test.DiffEq("test")
		assert.False(t, matcher.Matches("different"))
	})

	t.Run("オプション付きで差分がない場合、Matchesはtrueを返す", func(t *testing.T) {
		customComparer := cmp.Comparer(func(x, y string) bool {
			return len(x) == len(y)
		})
		matcherWithOpts := test.DiffEq("test", customComparer)
		assert.True(t, matcherWithOpts.Matches("same"))
	})
}

func TestDiffMatcherString(t *testing.T) {
	t.Run("差分がない場合、String()は空文字を返す", func(t *testing.T) {
		matcher := test.DiffEq("test")
		matcher.Matches("test")

		assert.Empty(t, matcher.String())
	})

	// 差分がある場合
	t.Run("差分がある場合、String()は差分を返す", func(t *testing.T) {
		matcher := test.DiffEq("test")
		matcher.Matches("different")
		result := matcher.String()

		assert.NotEmpty(t, result)
		assert.True(t, strings.HasPrefix(result, "diff(-got +want)"))
	})
}
