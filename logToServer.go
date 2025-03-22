package main

import "fmt"

func sendFakeShot(userName, msg string) error {
	fmt.Println("sending to " + userName + " msg: " + msg)
	return logToServer("NOTICE", userName, msg, "")
}

func sendErrorShot(userName, msg string) {
	err := logToServer("ERROR", userName, msg, "")
	if err != nil {
		dumpErrorToFile("cannot send error shot: " + err.Error())
	}
}

func logToServer(level, userName, message, screenshotData string) error {

	dto := RequestDTO[LogRecord]{
		Command:  "log",
		UserName: userName,
		Data: &LogRecord{
			Level:          level,
			LoggerName:     "pc-watcher",
			Message:        message,
			ScreenshotData: screenshotData,
		},
	}
	_, err := request(dto)
	if err != nil {
		return err
	}
	return nil
}
