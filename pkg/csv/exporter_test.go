package csv

import (
	"os"
	"testing"
	"time"

	"github.com/djoufson/git-report/internal/models"
)

func TestExportToCSV(t *testing.T) {
	exporter := NewExporter()
	
	commits := []models.Commit{
		{
			Branch:       "main",
			Hash:         "abc123def456",
			ShortHash:    "abc123d",
			Author:       "John Doe",
			Email:        "john@example.com",
			Date:         time.Date(2023, 6, 1, 12, 0, 0, 0, time.UTC),
			Message:      "Initial commit",
			FilesCount:   2,
			LinesAdded:   10,
			LinesDeleted: 0,
		},
	}

	filename := "test-output.csv"
	defer os.Remove(filename)

	err := exporter.ExportToCSV(commits, filename)
	if err != nil {
		t.Fatalf("ExportToCSV failed: %v", err)
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatal("CSV file was not created")
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read CSV file: %v", err)
	}

	contentStr := string(content)
	if !contains(contentStr, "Branch") || !contains(contentStr, "Commit Hash") {
		t.Errorf("CSV header missing. Got: %s", contentStr)
	}

	if !contains(contentStr, "main") || !contains(contentStr, "John Doe") {
		t.Errorf("CSV data missing. Got: %s", contentStr)
	}
}

func contains(s, substr string) bool {
	return len(substr) <= len(s) && (substr == s[:len(substr)] || (len(s) > 0 && contains(s[1:], substr)))
}