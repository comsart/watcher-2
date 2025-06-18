package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 2 {
		fmt.Println("STAGE 1 : I will download exec from github. os.Args: ", os.Args)
		downloadExec()
	} else if len(os.Args) == 3 {
		fmt.Println("STAGE 2 : I will overwrite watcher. os.Args: ", os.Args)
		overwriteOrgFileAndWork(os.Args[2])
	}
}
