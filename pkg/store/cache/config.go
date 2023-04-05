package cache

import "time"

type Config struct {
	Host        string
	Password    string
	DB          int
	IdleTimeout time.Duration
}
