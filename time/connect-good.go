package time

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jonboulle/clockwork"
)

var (
	stdOut io.Writer = os.Stdout
	clock            = clockwork.NewRealClock()
)

func ConnectGood() error {
	if _, err := fmt.Fprintln(stdOut, "connecting.."); err != nil {
		return fmt.Errorf("could not write to stdout: %v", err)
	}
	clock.Sleep(time.Second)
	if _, err := fmt.Fprintln(stdOut, "could not connect.."); err != nil {
		return fmt.Errorf("could not write to stdout: %v", err)
	}
	return nil
}
