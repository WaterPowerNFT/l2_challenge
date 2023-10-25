package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func getTime() time.Time {
	timeCur, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return timeCur
}
func main() {
	timeNow := getTime()
	fmt.Println(timeNow)
}
