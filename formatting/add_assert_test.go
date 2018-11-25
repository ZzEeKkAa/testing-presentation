package formatting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd2(t *testing.T) {
	assert.Equal(t, 3, Add(1, 2), "Add(%d,%d)", 1, 2)
}

//--- FAIL: TestAdd2 (0.00s)
//    add_assert_test.go:10:
//                Error Trace:    add_assert_test.go:10
//                Error:          Not equal:
//                                expected: 3
//                                actual  : 4
//                Test:           TestAdd2
//                Messages:       Add(1,2)
//FAIL
