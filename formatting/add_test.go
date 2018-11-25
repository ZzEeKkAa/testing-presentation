package formatting

import "testing"

func TestAdd(t *testing.T) {
	if res := Add(1, 2); res != 3 {
		t.Fatalf("Add(%d, %d) = %d, want %d", 1, 2, res, 3)
	}
}

//--- FAIL: TestAdd (0.00s)
//    add_test.go:7: Add(1, 2) = 4, want 3
//FAIL
