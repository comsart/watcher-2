package main

import (
	"fmt"
	"os"
)

// 2-nd run as temp

func overwriteOrgFile(orgExec string) {

	tempPath, _ := os.Executable()

	err := os.Rename(tempPath, orgExec)
	if err != nil {
		fmt.Println("Błąd nadpisywania pliku:", err)
		return
	}

	fmt.Println("Aktualizacja zakończona, restartowanie...")
	startAnotherIns(orgExec, watchFlag)
	os.Exit(0)
}
