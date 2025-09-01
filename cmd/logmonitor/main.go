package main

import (
	"flag"
	"fmt"
	"log"
	"logmonitor/internal/monitor"
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

	// Parse the log file
	logs, err := monitor.ParseLog(*logFile)
	if err != nil {
		log.Fatalf("Failed to parse log: %v", err)
	}

	// Monitor jobs and generate report
	err = monitor.MonitorJobs(logs, reportFile)
	if err != nil {
		log.Fatalf("Monitoring failed: %v", err)
	}

	fmt.Printf("Report generated at %s\n", reportFile)
}
