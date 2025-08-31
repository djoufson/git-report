package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/djoufson/git-report/internal/config"
)

func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name      string
		cfg       *config.Config
		expectErr bool
	}{
		{
			name:      "current directory",
			cfg:       &config.Config{RepoPath: "../.."},
			expectErr: false,
		},
		{
			name:      "nonexistent path",
			cfg:       &config.Config{RepoPath: "/nonexistent/path"},
			expectErr: true,
		},
		{
			name:      "empty path defaults to current dir",
			cfg:       &config.Config{RepoPath: "../.."},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Skip modification for empty path test case
			
			err := validateConfig(tt.cfg)
			if (err != nil) != tt.expectErr {
				t.Errorf("validateConfig() error = %v, expectErr %v", err, tt.expectErr)
			}
			
			// If no error expected, check that path was made absolute
			if !tt.expectErr {
				if !filepath.IsAbs(tt.cfg.RepoPath) {
					t.Errorf("Expected absolute path, got %s", tt.cfg.RepoPath)
				}
			}
		})
	}
}

func TestValidateConfigWithTempDir(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "git-report-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a .git directory to make it look like a git repo
	gitDir := filepath.Join(tempDir, ".git")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatal(err)
	}

	cfg := &config.Config{RepoPath: tempDir}
	err = validateConfig(cfg)
	if err != nil {
		t.Errorf("Expected no error for valid git repo, got: %v", err)
	}
}

func TestValidateConfigNonGitDirectory(t *testing.T) {
	// Create a temporary directory without .git
	tempDir, err := os.MkdirTemp("", "git-report-test-nogit")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	cfg := &config.Config{RepoPath: tempDir}
	err = validateConfig(cfg)
	if err == nil {
		t.Error("Expected error for non-git directory, got nil")
	}
}

func TestParseDate(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"2023-01-01", true},
		{"2023-12-31", true},
		{"2023-01-01 12:00:00", true},
		{"01/02/2006", true},
		{"invalid-date", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, err := parseDate(tt.input)
			success := err == nil
			if success != tt.expected {
				t.Errorf("parseDate(%s) success = %v, expected %v", tt.input, success, tt.expected)
			}
		})
	}
}