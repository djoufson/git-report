# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Git Report is a Go CLI tool that inspects git log history for all local branches and exports the data to CSV format. The tool supports:
- Period range filtering
- Author filtering  
- Comprehensive log formatting

## Development Commands

```bash
# Build the project
go build -o git-report ./cmd/git-report

# Run tests
go test ./...

# Run specific test package
go test ./pkg/git
go test ./pkg/csv
go test ./pkg/filter

# Format code
go fmt ./...

# Lint (requires golangci-lint)
golangci-lint run

# Get dependencies
go mod tidy

# Run the tool
./git-report [flags]

# Example usage
./git-report --verbose --output test-report.csv
./git-report --since "2023-01-01" --author "username"
./git-report --repo-path /path/to/other/repo --output other-repo.csv
./git-report --repo-path ../sibling-project --verbose
```

## Architecture Notes

The project follows a clean architecture pattern:

- **cmd/git-report/**: Main entry point with Cobra CLI framework
- **pkg/git/**: Git command execution and log parsing
- **pkg/csv/**: CSV file generation and export
- **pkg/filter/**: Date range and author filtering logic
- **internal/models/**: Core data structures (Commit)
- **internal/config/**: Configuration management

Key implementation details:
- Uses `git log --pretty=format` with `--numstat` for comprehensive data
- Processes commits with file change statistics
- Supports multiple date formats and flexible filtering
- Handles all local branches automatically unless specific branches are specified