package main

import (
	"log"

	"github.com/forever1young/account-onboard-alltoken/config"
	"github.com/forever1young/account-onboard-alltoken/router"
)

func main() {
	cfg := config.Load()
	r := router.New(cfg)

	log.Printf("account-onboard listening on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
