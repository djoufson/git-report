package git

import (
	"testing"
	"time"
)

func TestNewParser(t *testing.T) {
	parser := NewParser()
	if parser == nil {
		t.Fatal("NewParser() returned nil")
	}
}

func TestParseCommits(t *testing.T) {
	parser := NewParser()
	
	mockOutput := `abc123|abc1|John Doe|john@example.com|1640995200|Initial commit
1	0	README.md

def456|def4|Jane Smith|jane@example.com|1641081600|Add feature
2	1	main.go
3	0	config.go`

	commits, err := parser.parseCommits(mockOutput, "main")
	if err != nil {
		t.Fatalf("parseCommits failed: %v", err)
	}

	if len(commits) != 2 {
		t.Fatalf("Expected 2 commits, got %d", len(commits))
	}

	commit1 := commits[0]
	if commit1.Hash != "abc123" {
		t.Errorf("Expected hash 'abc123', got '%s'", commit1.Hash)
	}
	if commit1.ShortHash != "abc1" {
		t.Errorf("Expected short hash 'abc1', got '%s'", commit1.ShortHash)
	}
	if commit1.Author != "John Doe" {
		t.Errorf("Expected author 'John Doe', got '%s'", commit1.Author)
	}
	if commit1.Email != "john@example.com" {
		t.Errorf("Expected email 'john@example.com', got '%s'", commit1.Email)
	}
	if commit1.Message != "Initial commit" {
		t.Errorf("Expected message 'Initial commit', got '%s'", commit1.Message)
	}
	if commit1.Branch != "main" {
		t.Errorf("Expected branch 'main', got '%s'", commit1.Branch)
	}
	if commit1.FilesCount != 1 {
		t.Errorf("Expected 1 file changed, got %d", commit1.FilesCount)
	}
	if commit1.LinesAdded != 1 {
		t.Errorf("Expected 1 line added, got %d", commit1.LinesAdded)
	}

	expectedTime := time.Unix(1640995200, 0)
	if !commit1.Date.Equal(expectedTime) {
		t.Errorf("Expected date %v, got %v", expectedTime, commit1.Date)
	}
}