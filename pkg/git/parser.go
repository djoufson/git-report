package git

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/djoufson/git-report/internal/models"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) GetLocalBranches() ([]string, error) {
	cmd := exec.Command("git", "branch", "--format=%(refname:short)")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get branches: %w", err)
	}

	var branches []string
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		branch := strings.TrimSpace(scanner.Text())
		if branch != "" {
			branches = append(branches, branch)
		}
	}

	return branches, scanner.Err()
}

func (p *Parser) GetCommits(branch string, since, until *time.Time) ([]models.Commit, error) {
	args := []string{"log", "--pretty=format:%H|%h|%an|%ae|%at|%s", "--numstat"}
	
	if since != nil {
		args = append(args, fmt.Sprintf("--since=%s", since.Format("2006-01-02")))
	}
	if until != nil {
		args = append(args, fmt.Sprintf("--until=%s", until.Format("2006-01-02")))
	}
	
	args = append(args, branch)
	
	cmd := exec.Command("git", args...)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get commits for branch %s: %w", branch, err)
	}

	return p.parseCommits(string(output), branch)
}

func (p *Parser) parseCommits(output, branch string) ([]models.Commit, error) {
	var commits []models.Commit
	lines := strings.Split(output, "\n")
	
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		
		if !strings.Contains(line, "|") {
			continue
		}
		
		parts := strings.Split(line, "|")
		if len(parts) != 6 {
			continue
		}
		
		timestamp, err := strconv.ParseInt(parts[4], 10, 64)
		if err != nil {
			continue
		}
		
		commit := models.Commit{
			Branch:    branch,
			Hash:      parts[0],
			ShortHash: parts[1],
			Author:    parts[2],
			Email:     parts[3],
			Date:      time.Unix(timestamp, 0),
			Message:   parts[5],
		}
		
		i++
		for i < len(lines) && strings.TrimSpace(lines[i]) != "" && !strings.Contains(lines[i], "|") {
			statLine := strings.TrimSpace(lines[i])
			if statLine != "" {
				parts := strings.Split(statLine, "\t")
				if len(parts) >= 3 {
					if parts[0] != "-" {
						if added, err := strconv.Atoi(parts[0]); err == nil {
							commit.LinesAdded += added
						}
					}
					if parts[1] != "-" {
						if deleted, err := strconv.Atoi(parts[1]); err == nil {
							commit.LinesDeleted += deleted
						}
					}
					commit.FilesCount++
				}
			}
			i++
		}
		if i < len(lines) {
			i--
		}
		
		commits = append(commits, commit)
	}
	
	return commits, nil
}