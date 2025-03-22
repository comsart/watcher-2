package main

import (
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
		downloadExec()
	case watchFlag:
		watch()
	default:
		overwriteOrgFile(secParam)
	}

}
