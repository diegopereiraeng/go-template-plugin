package main

import (
	"log"
	"os/exec"
)

// runPlugin function is called from main.go and receives the parameters
func runPlugin(templatePath string, valuesPath string, outputPath string) {
	// Build the command
	cmd := exec.Command("./go-template", "-t", templatePath, "-f", valuesPath, "-o", outputPath)

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to execute command: %s", err)
	}

	// Log the output
	log.Printf("Command Output:\n%s\n", output)
}
