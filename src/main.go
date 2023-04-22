package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	if eventPath == "" {
		fmt.Println("GITHUB_EVENT_PATH environment variable is not set")
		return
	}

	eventData, err := ioutil.ReadFile(eventPath)
	if err != nil {
		fmt.Println("Failed to read GITHUB_EVENT_PATH:", err)
		return
	}

	var event interface{}
	err = json.Unmarshal(eventData, &event)
	if err != nil {
		fmt.Println("Failed to parse GITHUB_EVENT_PATH:", err)
		return
	}

	fmt.Println("Contents of GITHUB_EVENT_PATH:")
	fmt.Println(string(eventData))
}
