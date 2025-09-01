# Log Monitor Application

This Go application parses a `.log` file, tracks job start and end times, and generates warnings or errors based on duration thresholds.


## Folder Structure

```
logmonitor/
├── cmd/
│   └── logmonitor/       # main entrypoint
│       └── main.go
├── internal/
│   └── monitor/          # core parsing + monitoring logic
│       ├── monitor.go
│       └── monitor_test.go
├── examples/
│   └── logs.log          # sample log file
├── output/
│   └── report.log        # generated monitoring report
├── go.mod
├── go.sum
├── .gitignore
├── .gitattributes
├── Makefile
└── README.md
```


## Usage

### Application flags
The application supports the following command-line flags:

* `--log-file`
  Path to the input log file to parse.
  **Default:** `examples/logs.log` (current working directory if not set)

* `--report-path`
  Path to the folder where the monitoring report will be saved.  
  **Default:** `output/` (current working directory if not set).

```bash
# Use default paths
make run

# Specify custom log file and report folder
go run cmd/logmonitor/main.go --log-file /path/to/mylogs.log --report-path /path/to/report-folder
```

### Thresholds

* **Warning**: Job duration exceeds 5 minutes
* **Error**: Job duration exceeds 10 minutes


## Makefile

Build and run the application using Makefile:
* `make lint`: Lint all Go files
* `make run`: Run Go application
* `make test`: Run all Go tests
* `make build`: Compile the Go application
* `make clean`: Clean binaries and reports


## Example Logs

```
12:00:00, ImportData, START, 46578
12:07:10, ImportData, END, 46578
13:00:00, ExportReport, START, 48234
13:12:05, ExportReport, END, 48234
14:30:00, SyncDatabase, START, 49212
14:34:00, SyncDatabase, END, 49212
```

The monitor will print warnings or errors and also write them to a report file.
