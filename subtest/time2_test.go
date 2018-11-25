package subtest

import (
	"fmt"
	"testing"
	"time"
)

func TestTime2(t *testing.T) {
	for _, tc := range []struct {
		gmt  string
		loc  string
		want string
	}{
		{"12:31", "Europe/Zuri", "13:31"},
		{"12:31", "America/New_York", "7:31"},
		{"08:08", "Australia/Sydney", "18:08"},
	} {
		t.Run(fmt.Sprintf("%s in %s", tc.gmt, tc.loc), func(t *testing.T) {
			loc, err := time.LoadLocation(tc.loc)
			if err != nil {
				t.Fatal("could not load location")
			}
			gmt, _ := time.Parse("15:04", tc.gmt)
			if got := gmt.In(loc).Format("15:04"); got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

//--- FAIL: TestTime2 (0.00s)
//    --- FAIL: TestTime2/12:31_in_Europe/Zuri (0.00s)
//        time2_test.go:22: could not load location
//    --- FAIL: TestTime2/12:31_in_America/New_York (0.00s)
//        time2_test.go:26: got 07:31; want 7:31\
//    --- PASS: TestTime2/08:08_in_Australia/Sydney (0.00s)
//FAIL
