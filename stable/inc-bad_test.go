package stable

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestIncBad(t *testing.T) {
	c := NewCounterBad()
	c.Inc()
	assert.Equal(t, c.Val(), 1, "CounterBad.Inc()")
}

//--- FAIL: TestIncBad (0.00s)
//        assert.go:24: got 2 want 1 CounterBad.Inc()
//FAIL
