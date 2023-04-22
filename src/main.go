package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	wd, err := os.Getwd()

	// Check if a Terragrunt command is provided as a command-line argument
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./terragrunt <terragrunt_command>, <terragrunt_path>")
	}

	arguments := strings.Join(os.Args[1:], " ")
	args := strings.Split(arguments, ",")

	terragruntCommand := args[0]
	terragruntPath := func(args []string) string {
		if len(args) > 1 {
			return strings.TrimSpace(args[1])
		}
		return ""
	}(args)

	// Set up the Terragrunt command with the provided command-line argument
	cmd := exec.Command("terragrunt", terragruntCommand)

	// Set the working directory if needed
	cmd.Dir = wd + terragruntPath

	// Set environment variables if needed
	//cmd.Env = append(os.Environ(), "AWS_ACCESS_KEY_ID=your_aws_access_key", "AWS_SECRET_ACCESS_KEY=your_aws_secret_access_key")

	// Redirect the command's stdout and stderr to os.Stdout and os.Stderr respectively
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the Terragrunt command
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Terragrunt command completed successfully!")
}
