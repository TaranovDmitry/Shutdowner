package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Set time to shutdown in minutes: ")
	scanner.Scan()
	text := scanner.Text()
	textInt, err := strconv.Atoi(text)
	if err != nil {
		fmt.Printf("failed to convert the string: %s", err)
		os.Exit(1)
	}
	textInt *= 60

	command := fmt.Sprintf("shutdown /s /t %d", textInt)

	file, err := createFile()
	path := filePath(err)
	_, err = file.WriteString(command)
	if err != nil {
		fmt.Printf("failed to wrtite string: %s", err)
		os.Exit(1)
	}
	execFile(path, err)

	defer func() {
		err = os.Remove("shutdowner.bat")
		if err != nil {
			fmt.Printf("failed to remove file: %s", err)
			os.Exit(1)
		}
	}()
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Printf("failed to close file: %s", err)
			os.Exit(1)
		}
	}()

}

func execFile(path string, err error) {
	exec.Command(`cmd.exe`, `/C`, path+"\\shutdowner.bat").Run()
	if err != nil {
		fmt.Printf("failed to remove file %s", err)
		os.Exit(1)
	}
}

func filePath(err error) string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("failedt to get path: %s", err)
		os.Exit(1)
	}
	return path
}

func createFile() (*os.File, error) {
	file, err := os.OpenFile("shutdowner.bat", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("failed to create file %s", err)
		os.Exit(1)
	}
	return file, err
}
