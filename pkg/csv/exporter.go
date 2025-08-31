package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/djoufson/git-report/internal/models"
)

type Exporter struct{}

func NewExporter() *Exporter {
	return &Exporter{}
}

func (e *Exporter) ExportToCSV(commits []models.Commit, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{
		"Branch",
		"Commit Hash",
		"Short Hash",
		"Author",
		"Email",
		"Date",
		"Message",
		"Files Changed",
		"Lines Added",
		"Lines Deleted",
	}

	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write CSV header: %w", err)
	}

	for _, commit := range commits {
		record := []string{
			commit.Branch,
			commit.Hash,
			commit.ShortHash,
			commit.Author,
			commit.Email,
			commit.Date.Format("2006-01-02 15:04:05"),
			commit.Message,
			strconv.Itoa(commit.FilesCount),
			strconv.Itoa(commit.LinesAdded),
			strconv.Itoa(commit.LinesDeleted),
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write CSV record: %w", err)
		}
	}

	return nil
}