package main

import (
	"fmt"
	"github.com/pkg/browser"
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
		if instructions.ShouldBeForced {
			if !detectPage() {
				err := browser.OpenURL(serverAddress)
				if err != nil {
					events = addEvent(events, "cant open browser: "+err.Error())
				}
			}
			terminationsStarted := time.Now()
			for {
				terminationReports := terminateProcesses(instructions.AppsToTerminate)
				events = append(events, terminationReports...)
				if time.Now().Sub(terminationsStarted) > 4*time.Second {
					break
				}
				time.Sleep(1 * time.Second)
			}
		} else {
			time.Sleep(5 * time.Minute)
		}

		if len(events) > 200 || (time.Since(eventsDumped) > 30*time.Minute && len(events) > 0) {
			eventsString := buildEventsString(events)
			err := sendFakeShot(userName, eventsString)
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
