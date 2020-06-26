package help

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test that the argument passed to Handle function is parsed and routed properly
func TestHandle(t *testing.T) {
	cases := []struct {
		arg    string
		expect []string
	}{
		{arg: at, expect: CMD[at]},
		{arg: in, expect: CMD[in]},
		{arg: lsg, expect: CMD[lsg]},
		{arg: lsc, expect: CMD[lsc]},
		{arg: update, expect: CMD[update]},
		{arg: del, expect: CMD[del]},
		{arg: all, expect: CMD[all]},
	}

	for _, c := range cases {
		res, err := Handle(c.arg)
		assert.NoError(t, err)
		assert.Equal(t, c.expect, res)
	}
}
