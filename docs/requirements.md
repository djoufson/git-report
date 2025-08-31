# Git Report Requirements

## Overview

Git Report is a Go CLI tool that analyzes git commit history across all local branches and exports the data to CSV format for reporting and analysis purposes.

## Core Features

### 1. Branch Analysis
- Scan all local git branches in the repository
- Extract commit history from each branch
- Handle branch-specific commit data

### 2. Filtering Options
- **Date Range Filtering**: Specify start and end dates for commit analysis
  - Support common date formats (YYYY-MM-DD, relative dates)
  - Default to last 30 days if no range specified
- **Author Filtering**: Filter commits by specific authors
  - Support multiple author names
  - Case-insensitive matching
  - Email address support

### 3. Data Export
- Export to CSV format with comprehensive commit information
- Configurable output file path
- Include the following fields:
  - Branch name
  - Commit hash (short and full)
  - Author name and email
  - Commit date/time
  - Commit message
  - Files changed count
  - Lines added/deleted

### 4. Output Format
- Clean, readable CSV structure
- Proper escaping for special characters in commit messages
- Sortable by date, author, or branch
- Optional summary statistics

## Command Line Interface

```bash
git-report [flags]

Flags:
  --since, -s     Start date (YYYY-MM-DD or relative like "1 week ago")
  --until, -u     End date (YYYY-MM-DD or relative like "yesterday") 
  --author, -a    Filter by author name/email (can specify multiple)
  --output, -o    Output CSV file path (default: git-report.csv)
  --branches, -b  Specific branches to analyze (default: all local branches)
  --verbose, -v   Verbose output with progress information
  --help, -h      Show help information
```

## Technical Requirements

### Dependencies
- Go 1.22+ 
- Git installed and accessible in PATH
- Repository must be a valid git repository

### Performance
- Handle repositories with large commit histories efficiently
- Stream processing for memory efficiency
- Progress indicators for long-running operations

### Error Handling
- Graceful handling of invalid git repositories
- Clear error messages for missing dependencies
- Validation of date ranges and author inputs

## Example Usage

```bash
# Generate report for last 30 days
git-report

# Report for specific date range
git-report --since "2023-01-01" --until "2023-12-31"

# Filter by specific authors
git-report --author "john.doe@company.com" --author "jane.smith"

# Custom output file
git-report --output reports/team-activity.csv --verbose
```
