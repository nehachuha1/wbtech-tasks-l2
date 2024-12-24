package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func getTime() time.Time {
	currentTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Printf("%v", err)
	}
	return currentTime
}

func main() {
	fmt.Println(getTime())
}
