package storage

import (
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrUrlNotFound = errors.New("url not found")
	ErrUrlExists   = errors.New("url already exists")
)
