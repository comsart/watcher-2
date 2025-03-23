package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// 1-st run as watcher

func downloadExec() {

	orgExec, _ := os.Executable()
	thisDir := filepath.Dir(orgExec)
	exe := filepath.Ext(filepath.Base(orgExec))
	updateURL := fmt.Sprintf("https://raw.githubusercontent.com/comsart/watcher-2/master/%s", "watcher"+exe)

	tmpFilePath := filepath.Join(thisDir, "temp"+exe)

	tmpFile, err := os.Create(tmpFilePath)
	if err != nil {
		fmt.Println("Błąd tworzenia pliku tymczasowego:", err)
		return
	}

	var githubResp *http.Response
	for {
		githubResp, err = http.Get(updateURL)
		if err != nil {
			fmt.Println("Błąd pobierania nowej wersji:", err)
			time.Sleep(time.Minute * 1)
			continue
		}
		break
	}

	defer githubResp.Body.Close()

	_, err = io.Copy(tmpFile, githubResp.Body)
	if err != nil {
		fmt.Println("Błąd zapisu nowej wersji:", err)
		return
	}

	fmt.Println("Nowa wersja pobrana:", tmpFilePath)

	tmpFile.Close()
	startAnotherIns(tmpFilePath, orgExec)
	fmt.Println("Nowa wersja uruchomiona, zamykanie starej...")
	os.Exit(0)
}
