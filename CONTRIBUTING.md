# Contributing to Git Report

Thank you for considering contributing to Git Report! We welcome contributions from the community.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/yourusername/git-report.git`
3. Create a feature branch: `git checkout -b feature/your-feature-name`

## Development Setup

### Prerequisites

- Go 1.22 or later
- Git installed and available in PATH

### Building and Testing

```bash
# Install dependencies
go mod tidy

# Build the project
go build -o git-report ./cmd/git-report

# Run tests
go test ./...

# Run specific test package
go test ./pkg/git -v

# Format code
go fmt ./...
```

## Making Changes

### Code Style

- Follow standard Go conventions and formatting
- Use `go fmt` to format your code
- Write clear, descriptive commit messages
- Add tests for new functionality

### Testing

- Write unit tests for all new code
- Ensure all existing tests continue to pass
- Aim for good test coverage

### Documentation

- Update documentation for any new features
- Include usage examples for new CLI flags
- Update README.md if needed

## Submitting Changes

1. Make sure all tests pass: `go test ./...`
2. Format your code: `go fmt ./...`
3. Commit your changes with a clear message
4. Push to your fork: `git push origin feature/your-feature-name`
5. Create a pull request

## Pull Request Guidelines

- Provide a clear description of the changes
- Include the motivation for the changes
- Reference any related issues
- Ensure CI passes
- Be responsive to feedback

## Reporting Issues

When reporting issues, please include:

- Go version (`go version`)
- Operating system and version
- Steps to reproduce the issue
- Expected vs actual behavior
- Any relevant error messages

## Feature Requests

We welcome feature requests! Please:

- Check if the feature already exists
- Provide a clear description of the use case
- Explain why the feature would be valuable
- Consider contributing the implementation

## Code of Conduct

This project follows the [Go Community Code of Conduct](https://golang.org/conduct). Please be respectful and inclusive in all interactions.

## Questions?

Feel free to open an issue for any questions about contributing.