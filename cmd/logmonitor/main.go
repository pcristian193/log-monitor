package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}

	// Define command-line flags with defaults
	logFile := flag.String("log-file", filepath.Join(cwd, "examples/logs.log"), "Path to the input log file")
	reportPath := flag.String("report-path", filepath.Join(cwd, "output"), "Path to the output report file")
	flag.Parse()

	// Check if output folder exists or create it
	if _, err := os.Stat(*reportPath); os.IsNotExist(err) {
		err := os.MkdirAll(*reportPath, 0755)
		if err != nil {
			log.Fatalf("Failed to create report directory: %v", err)
		}
	}

	// Add report file to path
	reportFile := *reportPath + "/report.log"

	// Print logFile and reportFile
	fmt.Println(*logFile)
	fmt.Println(reportFile)

}
