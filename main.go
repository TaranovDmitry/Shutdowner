package main

import (
	"fmt"
	"os"
)

func main() {
	var input int
	fmt.Scan(&input)
	fmt.Print("Set time to shutdown in seconds: ")

	_, err := os.CreateTemp("shutdown.bat", fmt.Sprintf("shutdown /s /t %d", input))
	if err != nil {
		fmt.Errorf("failed to create file: %s", err)
	}
}
