package time

import (
	"fmt"
	"time"
)

func ConnectBad() {
	fmt.Print("connecting to server..")
	time.Sleep(time.Second)
	fmt.Print("connected..")
}
