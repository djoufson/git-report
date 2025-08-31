# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Git Report is a Go CLI tool that inspects git log history for all local branches and exports the data to CSV format. The tool supports:
- Period range filtering
- Author filtering  
- Comprehensive log formatting

## Development Commands

Since this is a Go project, use standard Go tooling:

```bash
# Build the project
go build -o git-report

# Run tests
go test ./...

# Run specific test
go test -run TestFunctionName ./path/to/package

# Format code
go fmt ./...

# Lint (requires golangci-lint)
golangci-lint run

# Get dependencies
go mod tidy

# Run the tool
./git-report [flags]
```

## Architecture Notes

This CLI tool should be structured as:
- Main entry point handling CLI argument parsing
- Core git log parsing logic
- CSV export functionality
- Date range and author filtering components
- Comprehensive output formatting for git history analysis