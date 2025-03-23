package main

import (
	"fmt"
	"os"
)

const watchFlag = "--watch"

func main() {
	fmt.Println("hallo from main of", os.Args[0])

	var secParam string
	if len(os.Args) == 3 {
		secParam = os.Args[2]
	}

	switch secParam {
	case "":
		fmt.Println("case: I will download exec from github")
		downloadExec()
	case watchFlag:
		fmt.Println("case: I will watch")
		watch()
	default:
		fmt.Println("case: I will overwrite watcher")
		overwriteOrgFile(secParam)
	}

}
