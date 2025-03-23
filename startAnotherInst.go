package main

import (
	"fmt"
	"os"
	"os/exec"
)

func startAnotherIns(path, secParam string) {

	err := os.Chmod(path, 0755)
	if err != nil {
		fmt.Println("Cant set file as executable", err)
		return
	}

	userName := os.Args[1]
	execCmd := exec.Command(path, userName, secParam)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	err = execCmd.Start()
	if err != nil {
		errMsg := "Cant start file:" + err.Error()
		dumpErrorToFile(errMsg)
		panic(errMsg)
	}
}
