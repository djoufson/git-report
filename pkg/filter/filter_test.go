package filter

import (
	"testing"
	"time"

	"github.com/djoufson/git-report/internal/models"
)

func TestFilterCommits(t *testing.T) {
	filter := NewFilter()
	
	commits := []models.Commit{
		{
			Author: "John Doe",
			Email:  "john@example.com",
			Date:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Author: "Jane Smith",
			Email:  "jane@example.com",
			Date:   time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Author: "Bob Wilson",
			Email:  "bob@example.com",
			Date:   time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	since := time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)
	until := time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC)

	filtered := filter.FilterCommits(commits, []string{"jane"}, &since, &until)

	if len(filtered) != 1 {
		t.Fatalf("Expected 1 commit after filtering, got %d", len(filtered))
	}

	if filtered[0].Author != "Jane Smith" {
		t.Errorf("Expected author 'Jane Smith', got '%s'", filtered[0].Author)
	}
}

func TestMatchesAuthors(t *testing.T) {
	filter := NewFilter()
	
	commit := models.Commit{
		Author: "John Doe",
		Email:  "john.doe@example.com",
	}

	tests := []struct {
		authors  []string
		expected bool
	}{
		{[]string{}, true},
		{[]string{"john"}, true},
		{[]string{"John"}, true},
		{[]string{"doe"}, true},
		{[]string{"john.doe@example.com"}, true},
		{[]string{"jane"}, false},
		{[]string{"smith"}, false},
	}

	for _, tt := range tests {
		result := filter.matchesAuthors(commit, tt.authors)
		if result != tt.expected {
			t.Errorf("matchesAuthors(%v) = %v, want %v", tt.authors, result, tt.expected)
		}
	}
}