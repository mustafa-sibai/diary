# Diary

## Prerequisites

- Go 1.21+ (or whatever version you use)
- [Task](https://taskfile.dev) - Install with: `go install github.com/go-task/task/v3/cmd/task@latest`

## Setup

1. Clone the repository
2. Install Task: `go install github.com/go-task/task/v3/cmd/task@latest`
3. Install dev tools: `task install-tools`
4. Run the app: `task dev`

## Available Commands

```bash
task install-tools  # Install development dependencies (air, etc.)
task dev           # Run with auto-reload
task build         # Build the application
task run           # Run without auto-reload