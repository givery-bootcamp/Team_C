package test

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func DiffEq(v interface{}, opts ...cmp.Option) gomock.Matcher {
	return &diffMatcher{want: v, opts: opts}
}

type diffMatcher struct {
	want interface{}
	diff string
	opts cmp.Options
}

func (d *diffMatcher) Matches(x interface{}) bool {
	d.diff = cmp.Diff(x, d.want, d.opts...)
	return len(d.diff) == 0
}

func (d *diffMatcher) String() string {
	if d.diff == "" {
		return ""
	}
	return fmt.Sprintf("diff(-got +want) is %s", d.diff)
}
