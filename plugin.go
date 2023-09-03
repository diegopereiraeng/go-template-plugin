package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func runCommand(templatePath string, valuesPath string, outputPath string) error {
	fmt.Println("running command: go-template -t", templatePath, "-f", valuesPath, "-o", outputPath)

	// Build the command
	cmd := exec.Command("go-template", "-t", templatePath, "-f", valuesPath, "-o", outputPath)

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to execute command: %s", err)
		return err
	}

	// Generate Output
	fmt.Println("\033[33mGenerating Output...\033[0m") // Yellow text

	// Log the output
	log.Printf("Command Output:\n%s\n", output)
	return nil
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

	// Read the content of values.yaml
	content, err := os.ReadFile(valuesPath)
	if err != nil {
		fmt.Println("Error reading values.yaml:", err)
		return
	}

	// Use regex to replace all occurrences of "<+...>" with "place_holder"
	fmt.Println("\033[33mReplacing placeholders...\033[0m") // Yellow text
	re := regexp.MustCompile(`<\+[^>]+>`)
	updatedContent := re.ReplaceAllString(string(content), "place_holder")

	// Write the updated content back to values.yaml
	fmt.Println("\033[33mWriting to values.yaml...\033[0m") // Yellow text
	err = os.WriteFile(valuesPath, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Println("\033[31mError writing to values.yaml:\033[0m", err) // Red text
		return
	}
	fmt.Println("\033[32mSuccessfully wrote to values.yaml\033[0m") // Green text
	// Check if the template path is a file or folder
	fileInfo, err := os.Stat(templatePath)
	if err != nil {
		fmt.Println("\033[31mError reading template path:\033[0m", err) // Red text
		return
	}
	fmt.Println("\033[33mChecking Template Path...\033[0m") // Yellow text

	// fmt.Printf("\nDebug: fileInfo: %+v\\n", fileInfo)
	// fmt.Printf("\nDebug: Is Directory: %v\\n", fileInfo.IsDir())

	if fileInfo.IsDir() {
		fmt.Println("Template path is a directory")
		fmt.Println("Reading directory:", templatePath)
		fmt.Println("Running command for each YAML file in directory...")
		// If it's a directory, loop through each YAML file and run commands
		var failedFiles []string
		var successFiles []string
		filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".yaml" {
				fmt.Println("Running command for file:", path)

				err := runCommand(path, valuesPath, outputPath)
				if err != nil {
					failedFiles = append(failedFiles, path)
				} else {
					successFiles = append(successFiles, path)
				}
			}
			return nil
		})
		printStatusTable(successFiles, failedFiles, outputPath, templatePath)
	} else {
		fmt.Println("Template path is a file")
		// If it's a file, run your command
		fmt.Println("Running command for file:", templatePath)
		err := runCommand(templatePath, valuesPath, outputPath)
		if err != nil {
			printStatusTable([]string{}, []string{templatePath}, outputPath, templatePath)
		} else {
			printStatusTable([]string{templatePath}, []string{}, outputPath, templatePath)
		}
	}

	// Complete
	fmt.Println("\033[32mPlugin Completed Successfully\033[0m") // Green text
	fmt.Println("")
	// Developed By Diego Pereira
	fmt.Println("\033[35mDeveloped By Diego Pereira\033[0m") // Magenta text
}

func printStatusTable(successFiles []string, failedFiles []string, outputPath string, templatePath string) {

	fmt.Println("\n\033[1mResults:\033[0m")
	fmt.Println("------------------------------------------------------------------|")
	fmt.Printf("| %-50s | %-10s |\n", "FILE", "STATUS")
	fmt.Println("------------------------------------------------------------------|")
	for _, file := range successFiles {
		fmt.Printf("| %-50s | \033[32m%-10s\033[0m |\n", file, "SUCCESS")
	}
	for _, file := range failedFiles {
		fmt.Printf("| %-50s | \033[31m%-10s\033[0m |\n", file, "FAILED")
	}
	fmt.Println("------------------------------------------------------------------|")
	if len(failedFiles) > 0 {
		fmt.Printf("\033[31mSome files failed to templetize. Check the logs for more details.\033[0m\n")
	} else {
		fmt.Printf("\033[32mAll files templetized successfully. Check the output files in %s.\033[0m\n", outputPath)
	}
	fmt.Printf("\033[1mTemplate result file:\033[0m %s\n", filepath.Join(outputPath, filepath.Base(templatePath)))

}
