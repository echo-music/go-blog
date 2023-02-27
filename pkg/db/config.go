package db

import "time"

type Config struct {
	Type        string
	Link        string
	Charset     string
	MaxIdle     int
	MaxOpen     int
	MaxLifetime time.Duration
}
