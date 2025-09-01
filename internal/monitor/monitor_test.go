package monitor

import (
	"os"
	"testing"
	"time"
)

// helper to create a fake log slice
func fakeLog(start time.Time, end time.Time, desc string, pid int32) []LogEntry {
	return []LogEntry{
		{Timestamp: start, JobDescription: desc, Status: "START", PID: pid},
		{Timestamp: end, JobDescription: desc, Status: "END", PID: pid},
	}
}

// small helper for reading log reports
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (len(s) > 0 && (string(s[0:len(substr)]) == substr || contains(s[1:], substr)))
}

func TestMonitorJobs_NoThreshold(t *testing.T) {
	start, _ := time.Parse(timeLayout, "11:00:00")
	end, _ := time.Parse(timeLayout, "11:03:00")
	logs := fakeLog(start, end, "ok task", 1234) // 3 min < 5 min no output
	outfile := "test_report.log"
	defer os.Remove(outfile)

	err := MonitorJobs(logs, outfile)
	if err != nil {
		t.Fatalf("MonitorJobs returned error: %v", err)
	}

	data, _ := os.ReadFile(outfile)
	content := string(data)

	if contains(content, "WARNING") || contains(content, "ERROR") {
		t.Errorf("Expected nothing in output, got:\n%s", content)
	}
}

func TestMonitorJobs_WarningThreshold(t *testing.T) {
	start, _ := time.Parse(timeLayout, "11:00:00")
	end, _ := time.Parse(timeLayout, "11:06:00")
	logs := fakeLog(start, end, "long task", 2345) // 6 min > 5 min warning
	outfile := "test_report_warning.log"
	defer os.Remove(outfile)

	err := MonitorJobs(logs, outfile)
	if err != nil {
		t.Fatalf("MonitorJobs returned error: %v", err)
	}

	data, _ := os.ReadFile(outfile)
	content := string(data)

	if !contains(content, "WARNING") {
		t.Errorf("Expected WARNING in output, got:\n%s", content)
	}
}

func TestMonitorJobs_ErrorThreshold(t *testing.T) {
	start, _ := time.Parse(timeLayout, "12:00:00")
	end, _ := time.Parse(timeLayout, "12:11:00")
	logs := fakeLog(start, end, "too long task", 3456) // 11 min > 10 min error
	outfile := "test_report_error.log"
	defer os.Remove(outfile)

	err := MonitorJobs(logs, outfile)
	if err != nil {
		t.Fatalf("MonitorJobs returned error: %v", err)
	}

	data, _ := os.ReadFile(outfile)
	content := string(data)

	if !contains(content, "ERROR") {
		t.Errorf("Expected ERROR in output, got:\n%s", content)
	}
}
