package main

import (
	"github.com/sambakker4/gator/internal/database"
	"github.com/sambakker4/gator/internal/config"
)

type state struct {
	db     *database.Queries
	config *config.Config
}
