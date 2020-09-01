package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

//List of ntp servers
const (
	rupool     = "ru.pool.ntp.org"
	beevikpool = "0.beevik-ntp.pool.ntp.org"
	vniiftri   = "ntp1.vniiftri.ru"
)

func main() {

	fmt.Println("current time:", time.Now().Round(0))

	t, err := ntp.Time(vniiftri)
	if err != nil {
		log.Fatalf("Error of ntp server: %v", err)
	} else {
		fmt.Println("exact time:", t.Round(0))
	}

}
