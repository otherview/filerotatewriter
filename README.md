# RotatingLogWriter Package

## Overview
The RotatingLogWriter package provides a customizable log writing utility in Go, designed to handle log rotation based on file size limits. This package allows for efficient management of log files by automatically rotating them when they exceed a specified size. This is especially useful for applications that generate a large amount of log data and require file management to avoid unbounded log growth.

## Features
- **Automatic Log Rotation**: Automatically rotates log files when they reach a predetermined size, helping manage disk space usage efficiently.
- **Customizable File Naming**: Allows setting the base name for log files, which can be prefixed to uniquely identify each log file.
- **Configurable Directory Setting**: Logs can be written to a specific directory, ensuring that log management can be localized to a designated area.
- **Limit on Number of Files**: Supports setting a maximum number of log files to retain, after which older files can be deleted or archived.

## Installation
To include the RotatingLogWriter in your Go project, you can typically import it using:
```go
import "github.com/otherview/rotatinglogwriter"
```
Ensure that you have the package files in your project or in a reachable location in your GOPATH.

## Usage
Here's how to get started with the `RotatingLogWriter`:

### Creating a New Log Writer
You must configure the log writer using the provided builder options functions. Here is an example of initializing a log writer:
```go
logWriter, err := New(
    WithDir("/var/log/myapp"),
    WithFileBaseName("app_log"),
    WithFileMaxSize(10 * 1024 * 1024), // 10 MB
    WithMaxNumberFiles(5),
)
if err != nil {
    log.Fatalf("Failed to create log writer: %v", err)
}
```

### Writing Logs
Use the `Write` method to write logs. It handles log rotation based on the configurations provided during initialization:
```go
_, err = logWriter.Write([]byte("This is a log message"))
if err != nil {
    log.Fatalf("Failed to write to log: %v", err)
}
```

### Starting the Log Writer
Ensure you start the log writer before beginning to log:
```go
err = logWriter.Start()
if err != nil {
    log.Fatalf("Failed to start log writer: %v", err)
}
```

## Configuration Options
- `WithDir(string)`: Sets the directory where log files will be stored.
- `WithFileBaseName(string)`: Sets the base name for log files.
- `WithFileMaxSize(int64)`: Sets the maximum size of each log file before it is rotated.
- `WithMaxNumberFiles(int)`: Sets the maximum number of log files to retain.

## Contributing
Contributions to the RotatingLogWriter package are welcome. Please submit a pull request or raise an issue on the project's repository.

## License
Apache 2.0 
