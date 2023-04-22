package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("GitHub repository owner:", os.Getenv("GITHUB_REPOSITORY_OWNER"))
	fmt.Println("GitHub repository name:", os.Getenv("GITHUB_REPOSITORY_NAME"))
	fmt.Println("GitHub event name:", os.Getenv("GITHUB_EVENT_NAME"))
	fmt.Println("GitHub ref:", os.Getenv("GITHUB_REF"))
	fmt.Println("GitHub sha:", os.Getenv("GITHUB_SHA"))
	fmt.Println("GitHub event path:", os.Getenv("GITHUB_EVENT_PATH"))

	// Access other environment variables for more context information as needed
}