package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {

	fmt.Println("current time:", time.Now().Round(0))

	t, err := ntp.Time("ntp1.vniiftri.ru")
	if err != nil {
		log.Fatalf("Error of ntp server: %v", err)
	} else {
		fmt.Println("exact time:", t.Round(0))
	}
}
