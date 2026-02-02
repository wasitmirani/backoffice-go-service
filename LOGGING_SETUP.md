# Daily File Logging Setup

## Overview

The application now supports daily file-based logging with automatic rotation. Log files are created daily with the format: `app-YYYY-MM-DD.log` (e.g., `app-2024-01-15.log`).

## Features

- ✅ **Daily Rotation**: Log files are automatically rotated daily
- ✅ **Size-based Rotation**: Log files are rotated when they reach a maximum size
- ✅ **Compression**: Old log files can be compressed to save space
- ✅ **Retention**: Configurable retention period for old log files
- ✅ **Multiple Outputs**: Support for stdout, file, or both (stack)

## Configuration

### Environment Variables

Add these to your `.env` file:

```bash
# Logging Configuration
LOG_CHANNEL=file          # Options: stdout, file, stack
LOG_LEVEL=debug           # Options: debug, info, warn, error
LOG_FILE_PATH=./storage/logs
LOG_FILE_NAME=app
LOG_MAX_SIZE=10           # Maximum size in MB before rotation
LOG_MAX_BACKUPS=5         # Maximum number of old log files to retain
LOG_MAX_AGE=28           # Maximum number of days to retain old log files
LOG_COMPRESS=true         # Whether to compress rotated log files
LOG_DAILY_ROTATE=true     # Enable daily rotation
```

### Logger Types

1. **stdout** - Logs to standard output (console)
2. **file** - Logs to daily rotated files
3. **stack** - Logs to both stdout and file

## Log File Structure

Log files are stored in the directory specified by `LOG_FILE_PATH`:

```
storage/
└── logs/
    ├── app-2024-01-15.log
    ├── app-2024-01-16.log
    ├── app-2024-01-17.log
    └── app-2024-01-17.log.gz  (compressed old files)
```

## Log Format

Each log entry includes:
- Timestamp
- Log level (DEBUG, INFO, WARN, ERROR, FATAL)
- File location (for debugging)
- Message
- Structured fields (key-value pairs)

Example:
```
2024/01/15 10:30:45 [INFO] internal/app/app.go:123: Application starting | version=1.0.0
2024/01/15 10:30:46 [INFO] internal/app/app.go:145: Primary database connected | driver=postgresql
```

## Usage

### Default (stdout)

If no logging configuration is provided, logs go to stdout:

```bash
LOG_CHANNEL=stdout
```

### File Logging

To enable file logging with daily rotation:

```bash
LOG_CHANNEL=file
LOG_FILE_PATH=./storage/logs
LOG_FILE_NAME=app
LOG_DAILY_ROTATE=true
LOG_MAX_SIZE=10
LOG_MAX_BACKUPS=5
LOG_MAX_AGE=28
LOG_COMPRESS=true
```

### Stack Logging (Both stdout and file)

To log to both console and file:

```bash
LOG_CHANNEL=stack
LOG_FILE_PATH=./storage/logs
LOG_FILE_NAME=app
LOG_DAILY_ROTATE=true
```

## Automatic Rotation

### Daily Rotation

When `LOG_DAILY_ROTATE=true`:
- A new log file is created each day
- File names include the date: `app-2024-01-15.log`
- Old files are automatically managed based on retention settings

### Size-based Rotation

When a log file reaches `LOG_MAX_SIZE` (in MB):
- The current file is rotated
- A new file is created
- Old files are compressed if `LOG_COMPRESS=true`

### Retention

- **Max Backups**: Keeps the last N log files (`LOG_MAX_BACKUPS`)
- **Max Age**: Deletes log files older than N days (`LOG_MAX_AGE`)

## Example Configuration

### Development

```bash
LOG_CHANNEL=stdout
LOG_LEVEL=debug
```

### Production

```bash
LOG_CHANNEL=file
LOG_LEVEL=info
LOG_FILE_PATH=/var/log/backoffice
LOG_FILE_NAME=app
LOG_MAX_SIZE=100
LOG_MAX_BACKUPS=30
LOG_MAX_AGE=90
LOG_COMPRESS=true
LOG_DAILY_ROTATE=true
```

### Development with File Logging

```bash
LOG_CHANNEL=stack
LOG_LEVEL=debug
LOG_FILE_PATH=./storage/logs
LOG_FILE_NAME=app
LOG_DAILY_ROTATE=true
```

## Code Example

The logger is automatically initialized in `cmd/main.go`:

```go
// Logger is initialized based on LOG_CHANNEL environment variable
// No code changes needed - just configure via environment variables
```

## Log Levels

- **debug**: Detailed information for debugging
- **info**: General informational messages
- **warn**: Warning messages
- **error**: Error messages
- **fatal**: Fatal errors (exits application)

## Troubleshooting

### Log files not created

1. Check that `LOG_FILE_PATH` directory exists or is writable
2. Ensure `LOG_CHANNEL=file` or `LOG_CHANNEL=stack`
3. Check file permissions

### Logs not rotating daily

1. Ensure `LOG_DAILY_ROTATE=true`
2. Check system time/date is correct
3. Verify the application is running continuously

### Too many log files

1. Reduce `LOG_MAX_BACKUPS`
2. Reduce `LOG_MAX_AGE`
3. Enable compression: `LOG_COMPRESS=true`

## Best Practices

1. **Production**: Use `LOG_CHANNEL=file` with appropriate retention
2. **Development**: Use `LOG_CHANNEL=stdout` or `LOG_CHANNEL=stack`
3. **Log Level**: Use `info` in production, `debug` in development
4. **Retention**: Set `LOG_MAX_AGE` based on compliance requirements
5. **Compression**: Enable compression to save disk space
6. **Path**: Use absolute paths in production (`/var/log/app`)

