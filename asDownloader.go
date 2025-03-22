package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func downloadExec() {

	orgExec, _ := os.Executable()
	thisDir := filepath.Dir(orgExec)
	exe := filepath.Ext(filepath.Base(orgExec))
	updateURL := fmt.Sprintf("https://raw.githubusercontent.com/comsart/watcher/master/%s", "watcher"+exe)

	// Pobranie nowej wersji

	pathToWriteTemp := filepath.Join(thisDir, "temp"+exe)

	tmpFile, err := os.Create(pathToWriteTemp)
	if err != nil {
		fmt.Println("Błąd tworzenia pliku tymczasowego:", err)
		return
	}
	defer tmpFile.Close()

	resp, err := http.Get(updateURL)
	if err != nil {
		fmt.Println("Błąd pobierania nowej wersji:", err)
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		fmt.Println("Błąd zapisu nowej wersji:", err)
		return
	}

	tmpFilePath := tmpFile.Name()
	fmt.Println("Nowa wersja pobrana:", tmpFilePath)

	// Ustawienia uprawnień do wykonania
	err = os.Chmod(tmpFilePath, 0755)
	if err != nil {
		fmt.Println("Błąd ustawiania uprawnień:", err)
		return
	}

	// Uruchomienie nowej wersji i zamknięcie starej
	startAnotherIns(tmpFilePath, orgExec)
	fmt.Println("Nowa wersja uruchomiona, zamykanie starej...")
	os.Exit(0)
}

func startAnotherIns(path, secParam string) {
	userName := os.Args[1]
	execCmd := exec.Command(path, userName, secParam)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	err := execCmd.Start()
	if err != nil {
		fmt.Println("Błąd uruchamiania nowej wersji:", err)
		return
	}
}
