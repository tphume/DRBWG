package in

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	cases := []struct {
		dur    string
		name   string
		now    time.Time
		expect time.Time
		err    error
	}{
		{dur: "1h20m", name: "valid name", now: time.Time{}, expect: time.Time{}.Add(time.Minute * 80), err: nil},
		{dur: "1h", name: "valid name", now: time.Time{}, expect: time.Time{}.Add(time.Minute * 60), err: nil},
		{dur: "20m", name: "valid name", now: time.Time{}, expect: time.Time{}.Add(time.Minute * 20), err: nil},
		{dur: "fvdsvfewf", name: "valid name", now: time.Time{}, expect: time.Time{}, err: INVALID_DURATION},
		{dur: "1h20m", name: "d", now: time.Time{}, expect: time.Time{}, err: INVALID_NAME},
	}

	for _, c := range cases {
		res, err := parse(c.dur, c.name, c.now)
		assert.Equal(t, c.err, err)
		assert.Equal(t, c.expect, res)
	}
}
