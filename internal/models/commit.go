package models

import "time"

type Commit struct {
	Branch      string
	Hash        string
	ShortHash   string
	Author      string
	Email       string
	Date        time.Time
	Message     string
	FilesCount  int
	LinesAdded  int
	LinesDeleted int
}