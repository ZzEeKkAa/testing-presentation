package simple

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Fatal("fatal")
	}
}

//--- FAIL: TestAdd (0.00s)
//    add_test.go:9: fatal
//FAIL
