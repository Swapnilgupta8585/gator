package main

import (
	"github.com/Swapnilgupta8585/blog_aggregator/internal/config"
	"github.com/Swapnilgupta8585/blog_aggregator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}
