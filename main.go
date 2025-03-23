package main

import (
	"fmt"
	"os"
)

const watchFlag = "--watch"

func main() {

	var secParam string
	if len(os.Args) == 3 {
		secParam = os.Args[2]
	}

	switch secParam {
	case "":
		fmt.Println("STAGE 1 : I will download exec from github. os.Args: ", os.Args)
		downloadExec()
	case watchFlag:
		fmt.Println("STAGE 3 : case: I will watch. os.Args: ", os.Args)
		watch()
	default:
		fmt.Println("STAGE 2 : I will overwrite watcher. os.Args: ", os.Args)
		overwriteOrgFile(secParam)
	}

}
