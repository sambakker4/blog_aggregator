package main

import (
	"github.com/sambakker4/blog_aggregator/internal/database"
	"github.com/sambakker4/blog_aggregator/internal/config"
)

type state struct {
	db     *database.Queries
	config *config.Config
}
