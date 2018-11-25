package stable

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestIncGood(t *testing.T) {
	c := NewCounterGood()
	before := c.Val()
	c.Inc()
	after := c.Val()
	assert.Equal(t, after-before, 1, "CounterGood.Inc()")
}

//--- FAIL: TestIncBad (0.00s)
//        assert.go:24: got 2 want 1 CounterGood.Inc()
//FAIL
