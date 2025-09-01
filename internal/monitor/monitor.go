package monitor

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	WarningThreshold = 5 * time.Minute
	ErrorThreshold   = 10 * time.Minute
	timeLayout       = "15:04:05" // HH:MM:SS
)

// LogEntry represents a single log line
type LogEntry struct {
	Timestamp      time.Time
	JobDescription string
	Status         string
	PID            int32
}

// ParseLog reads a .log file and returns structured log entries
func ParseLog(filePath string) ([]LogEntry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var logs []LogEntry
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 4 {
			continue // skip malformed lines
		}

		// Parse timestamp (HH:MM:SS) into time.Time (use today’s date)
		ts, err := time.Parse(timeLayout, strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid timestamp %q: %w", parts[0], err)
		}

		// Parse PID as int32
		pidVal, err := strconv.Atoi(strings.TrimSpace(parts[3]))
		if err != nil {
			return nil, fmt.Errorf("invalid PID %q: %w", parts[3], err)
		}

		logs = append(logs, LogEntry{
			Timestamp:      ts,
			JobDescription: strings.TrimSpace(parts[1]),
			Status:         strings.TrimSpace(parts[2]),
			PID:            int32(pidVal),
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return logs, nil
}

// MonitorJobs processes jobs, checking durations against thresholds
func MonitorJobs(logs []LogEntry, reportPath string) error {
	jobTracker := make(map[int32]LogEntry)

	reportFile, err := os.Create(reportPath)
	if err != nil {
		return err
	}
	defer reportFile.Close()

	for _, logEntry := range logs {
		pid := logEntry.PID
		switch logEntry.Status {
		case "START":
			jobTracker[pid] = logEntry
		case "END":
			if startLog, exists := jobTracker[pid]; exists {
				duration := logEntry.Timestamp.Sub(startLog.Timestamp)

				// Handle midnight rollover (END past midnight, before START)
				if duration < 0 {
					duration += 24 * time.Hour
				}

				var msg string
				if duration > ErrorThreshold {
					msg = fmt.Sprintf("ERROR: Job '%s' (PID %d) took %v\n", startLog.JobDescription, pid, duration)
				} else if duration > WarningThreshold {
					msg = fmt.Sprintf("WARNING: Job '%s' (PID %d) took %v\n", startLog.JobDescription, pid, duration)
				}

				if msg != "" {
					fmt.Print(msg)                     // print to console
					_, _ = reportFile.WriteString(msg) // write to report
				}
				delete(jobTracker, pid)
			}
		}
	}
	return nil
}
