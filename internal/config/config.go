package config

import "time"

type Config struct {
	Since     *time.Time
	Until     *time.Time
	Authors   []string
	Output    string
	Branches  []string
	Verbose   bool
}