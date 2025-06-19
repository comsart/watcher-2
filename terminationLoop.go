package main

import (
	"crypto/rand"
	"fmt"
	"github.com/pkg/browser"
	"math/big"
	"time"
)

func terminationLoop(userName string) {
	fmt.Println("starting termination loop")
	eventsDumped := time.Now()
	var events []string
	for {
		instructions, err := askServer(userName)
		if err != nil {
			fmt.Println("can't ask server", err.Error())
			time.Sleep(time.Second * 10)
			continue
		}
		if instructions.LessonIsDone { // --------------------------------------------------------------------------- english not done - internet closed
			if !detectPage() {
				err := browser.OpenURL(serverAddress)
				if err != nil {
					events = addEvent(events, "cant open browser: "+err.Error())
				}
			}

			terminationsStarted := time.Now()
			for {
				terminationMsgs := terminateProcesses(instructions.AppsToTerminate)
				events = append(events, terminationMsgs...)
				if time.Since(terminationsStarted) > 4*time.Second {
					break
				}
				time.Sleep(1 * time.Second)
			}
		} else if !instructions.LessonIsDone { // ---------------------------------------------------------------------------- english done - internet open
			time.Sleep(5 * time.Minute)

			// disturbing chrome and opera use
			randInt, err := rand.Int(rand.Reader, big.NewInt(10))
			if err != nil {
				panic(0)
			}
			if randInt.Cmp(big.NewInt(0)) == 0 {
				terminationMsgs := terminateProcesses([]string{"chrome", "opera"})
				events = append(events, terminationMsgs...)
			}

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
}
