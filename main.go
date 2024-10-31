package main

import (
	"fmt"
	"log"

	"github.com/Swapnilgupta8585/blog_aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	err = cfg.SetUser("Swapnil")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg.DbURL)
	fmt.Println(cfg.CurrentUserName)
}
