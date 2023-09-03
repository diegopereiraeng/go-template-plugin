package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func runCommand(templatePath string, valuesPath string, outputPath string) {
	fmt.Println("running command: go-template -t", templatePath, "-f", valuesPath, "-o", outputPath)

	// Build the command
	cmd := exec.Command("go-template", "-t", templatePath, "-f", valuesPath, "-o", outputPath)

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to execute command: %s", err)
	}

	// Generate Output
	fmt.Println("\033[33mGenerating Output...\033[0m") // Yellow text

	// Log the output
	log.Printf("Command Output:\n%s\n", output)
}

// runPlugin function is called from main.go and receives the parameters
func runPlugin(templatePath string, valuesPath string, outputPath string) {
	// Developed By Diego Pereira
	fmt.Println("\033[35mDeveloped By Diego Pereira\033[0m") // Magenta text

	// Initialize
	fmt.Println("\033[34mInitializing Plugin...\033[0m") // Blue text

	fmt.Printf("Template Path: %s\n", templatePath)
	fmt.Printf("Values Path: %s\n", valuesPath)
	fmt.Printf("Output Path: %s\n", outputPath)
	path, err := exec.Command("pwd").Output()
	if err != nil {
		log.Fatalf("Failed to get current path: %s", err)
	}
	if err != nil {
		log.Fatalf("Failed to get current path: %s", err)
	}
	fmt.Printf("Current Path: %s\n", path)

	// Check if result folder exists, create if not
	fmt.Println("\033[33mChecking Result Folder...\033[0m") // Yellow text
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		fmt.Println("\033[33mResult folder does not exist, creating...\033[0m") // Yellow text
		os.MkdirAll(outputPath, os.ModePerm)
	}

	// Check if the template path is a file or folder
	fileInfo, err := os.Stat(templatePath)
	if err != nil {
		fmt.Println("Error reading template path:", err)
		return
	}

	fmt.Printf("\nDebug: fileInfo: %+v\\n", fileInfo)
	fmt.Printf("\nDebug: Is Directory: %v\\n", fileInfo.IsDir())

	if fileInfo.IsDir() {
		fmt.Println("Template path is a directory")
		fmt.Println("Reading directory:", templatePath)
		fmt.Println("Running command for each YAML file in directory...")
		// If it's a directory, loop through each YAML file and run commands
		filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".yaml" {
				fmt.Println("Running command for file:", path)

				runCommand(path, valuesPath, outputPath)
			}
			return nil
		})
	} else {
		fmt.Println("Template path is a file")
		// If it's a file, run your command
		fmt.Println("Running command for file:", templatePath)
		runCommand(templatePath, valuesPath, outputPath)
	}

	// Complete
	fmt.Println("\033[32mPlugin Completed Successfully\033[0m") // Green text
	fmt.Println("")
	// Developed By Diego Pereira
	fmt.Println("\033[35mDeveloped By Diego Pereira\033[0m") // Magenta text
}
