package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/djoufson/git-report/internal/config"
	"github.com/djoufson/git-report/internal/models"
	"github.com/djoufson/git-report/pkg/csv"
	"github.com/djoufson/git-report/pkg/filter"
	"github.com/djoufson/git-report/pkg/git"
	"github.com/spf13/cobra"
)

func main() {
	var cfg config.Config

	rootCmd := &cobra.Command{
		Use:   "git-report",
		Short: "Generate CSV reports from git commit history",
		Long:  "A CLI tool that analyzes git commit history across all local branches and exports the data to CSV format.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return generateReport(&cfg)
		},
	}

	rootCmd.Flags().StringVarP(&cfg.Output, "output", "o", "git-report.csv", "Output CSV file path")
	rootCmd.Flags().StringSliceVarP(&cfg.Authors, "author", "a", nil, "Filter by author name/email")
	rootCmd.Flags().StringSliceVarP(&cfg.Branches, "branches", "b", nil, "Specific branches to analyze")
	rootCmd.Flags().BoolVarP(&cfg.Verbose, "verbose", "v", false, "Verbose output")

	rootCmd.Flags().StringP("since", "s", "", "Start date (YYYY-MM-DD or relative)")
	rootCmd.Flags().StringP("until", "u", "", "End date (YYYY-MM-DD or relative)")

	rootCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		return parseDateFlags(cmd, &cfg)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func parseDateFlags(cmd *cobra.Command, cfg *config.Config) error {
	sinceStr, _ := cmd.Flags().GetString("since")
	untilStr, _ := cmd.Flags().GetString("until")

	if sinceStr != "" {
		since, err := parseDate(sinceStr)
		if err != nil {
			return fmt.Errorf("invalid since date: %w", err)
		}
		cfg.Since = &since
	}

	if untilStr != "" {
		until, err := parseDate(untilStr)
		if err != nil {
			return fmt.Errorf("invalid until date: %w", err)
		}
		cfg.Until = &until
	}

	return nil
}

func parseDate(dateStr string) (time.Time, error) {
	layouts := []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
		"01/02/2006",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, dateStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unsupported date format: %s", dateStr)
}

func generateReport(cfg *config.Config) error {
	parser := git.NewParser()
	filter := filter.NewFilter()
	exporter := csv.NewExporter()

	if cfg.Verbose {
		fmt.Println("Starting git report generation...")
	}

	var branches []string
	var err error

	if len(cfg.Branches) > 0 {
		branches = cfg.Branches
	} else {
		branches, err = parser.GetLocalBranches()
		if err != nil {
			return fmt.Errorf("failed to get branches: %w", err)
		}
	}

	if cfg.Verbose {
		fmt.Printf("Analyzing %d branches: %s\n", len(branches), strings.Join(branches, ", "))
	}

	var allCommits []models.Commit

	for _, branch := range branches {
		if cfg.Verbose {
			fmt.Printf("Processing branch: %s\n", branch)
		}

		commits, err := parser.GetCommits(branch, cfg.Since, cfg.Until)
		if err != nil {
			if cfg.Verbose {
				fmt.Printf("Warning: failed to get commits for branch %s: %v\n", branch, err)
			}
			continue
		}

		allCommits = append(allCommits, commits...)
	}

	filteredCommits := filter.FilterCommits(allCommits, cfg.Authors, cfg.Since, cfg.Until)

	if cfg.Verbose {
		fmt.Printf("Found %d commits matching criteria\n", len(filteredCommits))
	}

	if err := exporter.ExportToCSV(filteredCommits, cfg.Output); err != nil {
		return fmt.Errorf("failed to export CSV: %w", err)
	}

	fmt.Printf("Report generated successfully: %s\n", cfg.Output)
	fmt.Printf("Total commits: %d\n", len(filteredCommits))

	return nil
}