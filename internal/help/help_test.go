package help

import (
	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test that the argument passed to Handle function is parsed and routed properly
func TestHandle(t *testing.T) {
	cases := []struct {
		arg    []string
		expect []string
	}{
		{arg: []string{at}, expect: CMD[at]},
		{arg: []string{in}, expect: CMD[in]},
		{arg: []string{lsg}, expect: CMD[lsg]},
		{arg: []string{lsc}, expect: CMD[lsc]},
		{arg: []string{update}, expect: CMD[update]},
		{arg: []string{del}, expect: CMD[del]},
		{arg: []string{all}, expect: CMD[all]},
		{arg: []string{"subcommandThatDoesNotExist"}, expect: CMD[all]},
	}

	for _, c := range cases {
		res, err := Handle(c.arg, &discordgo.MessageCreate{})
		assert.NoError(t, err)
		assert.Equal(t, c.expect, res)
	}
}
