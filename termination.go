package main

import (
	"github.com/shirou/gopsutil/process"
	"strings"
)

func terminateProcesses(appsToKill []string) []string {
	var events []string
	processes, err := process.Processes()
	if err != nil {
		events = addEvent(events, "cant read os processes"+err.Error())
	}
	for _, p := range processes {
		for _, namePart := range appsToKill {
			procName, err := p.Name()
			if err != nil {
				continue
			}

			if strings.Contains(strings.ToLower(procName), strings.ToLower(namePart)) {
				err = p.Terminate()
				if err != nil {
					events = addEvent(events, "cannot kill: "+procName+" "+err.Error())
				} else {
					events = addEvent(events, "I just killed "+procName)
				}
			}
		}
	}
	return events
}
