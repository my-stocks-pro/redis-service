package infrastructure

import "time"

type Config struct {
	StartDate string
	Name      string
	Port      string
	RedisDB   map[string]int
}

func NewConfig() Config {
	return Config{
		StartDate: time.Now().Format("2006-01-02 15:04"),
		Name:      "redis-service",
		Port:      ":9006",
		RedisDB: map[string]int{
			"earnings": 1,
			"approved": 2,
			"rejected": 3,
		},
	}
}
