package time

import (
	"bytes"
	"sync"
	"testing"
	"time"

	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/assert"
)

func TestConnectGood(t *testing.T) {
	buff := bytes.NewBuffer(nil)
	stdOut = buff
	fakeClock := clockwork.NewFakeClock()
	clock = fakeClock
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := ConnectGood(); err != nil {
			t.Fatalf("could not connect: %v", err)
		}
	}()
	fakeClock.BlockUntil(1) // Waiting until clock.Sleep is called
	assert.Equal(t, buff.String(), "connecting..\n")
	fakeClock.Advance(time.Millisecond * 500)
	fakeClock.BlockUntil(1)
	assert.Equal(t, buff.String(), "connecting..\n")
	fakeClock.Advance(time.Millisecond * 500)
	wg.Wait()
	assert.Equal(t, buff.String(), "connecting..\nconnected..\n")
}

//--- FAIL: TestConnectGood (0.00s)
//    connect-good_test.go:33:
//                Error Trace:    connect-good_test.go:33
//                Error:          Not equal:
//                                expected: "connecting..\ncould not connect..\n"
//                                actual  : "connecting..\nconnected..\n"
//
//                                Diff:
//                                --- Expected
//                                +++ Actual
//                                @@ -1,3 +1,3 @@
//                                 connecting..
//                                -could not connect..
//                                +connected..
//
//                Test:           TestConnectGood
//FAIL
