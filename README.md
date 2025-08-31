# Git Report

A Go CLI tool that analyzes git commit history across all local branches and exports the data to CSV format for reporting and analysis purposes.

## Features

- **Multi-branch Analysis**: Scans all local git branches automatically
- **Date Range Filtering**: Filter commits by date range with flexible date formats
- **Author Filtering**: Filter commits by specific authors (supports multiple authors)
- **CSV Export**: Clean, structured CSV output with comprehensive commit data
- **Detailed Information**: Includes commit hash, author, date, message, and file change statistics

## Installation

### From Source

```bash
git clone https://github.com/djoufson/git-report.git
cd git-report
go build -o git-report ./cmd/git-report
```

### Using Go Install

```bash
go install github.com/djoufson/git-report/cmd/git-report@latest
```

## Usage

### Basic Usage

Generate a report for all commits in the last 30 days:
```bash
./git-report
```

### Advanced Usage

```bash
# Filter by date range
./git-report --since "2023-01-01" --until "2023-12-31"

# Filter by specific authors
./git-report --author "john.doe@company.com" --author "jane.smith"

# Analyze specific branches only
./git-report --branches "main" --branches "develop"

# Custom output file with verbose logging
./git-report --output reports/team-activity.csv --verbose

# Combine filters
./git-report --since "2023-06-01" --author "john" --output june-john.csv
```

## Command Line Options

| Flag | Short | Description | Example |
|------|-------|-------------|---------|
| `--output` | `-o` | Output CSV file path | `--output report.csv` |
| `--since` | `-s` | Start date (YYYY-MM-DD) | `--since "2023-01-01"` |
| `--until` | `-u` | End date (YYYY-MM-DD) | `--until "2023-12-31"` |
| `--author` | `-a` | Filter by author (can specify multiple) | `--author "john@example.com"` |
| `--branches` | `-b` | Specific branches to analyze | `--branches "main"` |
| `--verbose` | `-v` | Enable verbose output | `--verbose` |
| `--help` | `-h` | Show help information | `--help` |

## CSV Output Format

The generated CSV file contains the following columns:

| Column | Description |
|--------|-------------|
| Branch | Git branch name |
| Commit Hash | Full commit hash |
| Short Hash | Abbreviated commit hash |
| Author | Commit author name |
| Email | Commit author email |
| Date | Commit date and time |
| Message | Commit message |
| Files Changed | Number of files modified |
| Lines Added | Total lines added |
| Lines Deleted | Total lines deleted |

## Development

### Requirements

- Go 1.19 or later
- Git installed and available in PATH

### Building

```bash
go mod tidy
go build -o git-report ./cmd/git-report
```

### Testing

```bash
go test ./...
```

### Project Structure

```
git-report/
├── cmd/git-report/     # Main application entry point
├── pkg/
│   ├── git/           # Git log parsing functionality
│   ├── csv/           # CSV export functionality
│   └── filter/        # Commit filtering logic
├── internal/
│   ├── models/        # Data models
│   └── config/        # Configuration structures
└── docs/              # Documentation
```

## Examples

### Team Activity Report

Generate a quarterly report for your development team:

```bash
./git-report \
  --since "2023-07-01" \
  --until "2023-09-30" \
  --output quarterly-report.csv \
  --verbose
```

### Individual Developer Analysis

Analyze commits from a specific developer:

```bash
./git-report \
  --author "jane.smith@company.com" \
  --since "2023-06-01" \
  --output jane-commits.csv
```

### Release Preparation

Get commits from specific branches for release notes:

```bash
./git-report \
  --branches "release/v1.0" \
  --branches "main" \
  --since "2023-08-01" \
  --output release-commits.csv
```

## Error Handling

The tool provides clear error messages for common issues:

- **Not a git repository**: Ensure you're running the command in a git repository
- **Invalid date format**: Use YYYY-MM-DD format for dates
- **No commits found**: Check your date range and author filters
- **Permission issues**: Ensure write permissions for the output directory

## License

This project is open source and available under the [MIT License](LICENSE).