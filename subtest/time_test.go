package subtest

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	for _, tc := range []struct {
		gmt  string
		loc  string
		want string
	}{
		{"12:31", "Europe/Zuri", "13:31"},     // incorrect location name
		{"12:31", "America/New_York", "7:31"}, // should be 07:31
		{"08:08", "Australia/Sydney", "18:08"},
	} {
		loc, err := time.LoadLocation(tc.loc)
		if err != nil {
			t.Fatalf("could not load location %q", tc.loc)
		}
		gmt, _ := time.Parse("15:04", tc.gmt)
		if got := gmt.In(loc).Format("15:04"); got != tc.want {
			t.Errorf("In(%s, %s) = %s; want %s", tc.gmt, tc.loc, got, tc.want)
		}
	}
}

//--- FAIL: TestTime (0.00s)
//    time_test.go:20: could not load location "Europe/Zuri"
//FAIL
