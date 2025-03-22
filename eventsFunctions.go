package main

import (
	"fmt"
	"os"
	"time"
)

func addEvent(events []string, text string) []string {
	currentTime := time.Now()
	isoTime := currentTime.Format(time.RFC3339)
	events = append(events, isoTime+" "+text)
	return events
}

func dumpErrorToFile(errMsg string) {
	file, _ := os.OpenFile("go-err-log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	isoDateTime := time.Now().Format(time.RFC3339)
	fmt.Println("err to file:", errMsg)
	_, _ = file.WriteString(isoDateTime + " " + errMsg + "\n")
}
