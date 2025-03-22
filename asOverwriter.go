package main

import (
	"fmt"
	"os"
)

func overwriteOrgFile(orgExec string) {

	tempPath, _ := os.Executable()

	err := os.Rename(tempPath, orgExec)
	if err != nil {
		fmt.Println("Błąd nadpisywania pliku:", err)
		return
	}

	fmt.Println("Aktualizacja zakończona, restartowanie...")
	startAnotherIns(tempPath, watchFlag)
	os.Exit(0)
}
