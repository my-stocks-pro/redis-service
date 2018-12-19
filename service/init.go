package service

import "os"

func NewPSQL() *TypePSQL {
	return &TypePSQL{
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGNAME"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASS"),
	}
}
