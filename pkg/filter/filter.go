package filter

import (
	"strings"
	"time"

	"github.com/djoufson/git-report/internal/models"
)

type Filter struct{}

func NewFilter() *Filter {
	return &Filter{}
}

func (f *Filter) FilterCommits(commits []models.Commit, authors []string, since, until *time.Time) []models.Commit {
	var filtered []models.Commit
	
	for _, commit := range commits {
		if !f.matchesAuthors(commit, authors) {
			continue
		}
		
		if !f.matchesTimeRange(commit, since, until) {
			continue
		}
		
		filtered = append(filtered, commit)
	}
	
	return filtered
}

func (f *Filter) matchesAuthors(commit models.Commit, authors []string) bool {
	if len(authors) == 0 {
		return true
	}
	
	for _, author := range authors {
		if strings.Contains(strings.ToLower(commit.Author), strings.ToLower(author)) ||
		   strings.Contains(strings.ToLower(commit.Email), strings.ToLower(author)) {
			return true
		}
	}
	
	return false
}

func (f *Filter) matchesTimeRange(commit models.Commit, since, until *time.Time) bool {
	if since != nil && commit.Date.Before(*since) {
		return false
	}
	
	if until != nil && commit.Date.After(*until) {
		return false
	}
	
	return true
}