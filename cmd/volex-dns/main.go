package main

import (
	"log"

	"github.com/TheRemyyy/volex-dns/internal/config"
	"github.com/TheRemyyy/volex-dns/internal/dnsserver"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	server := dnsserver.NewServer(cfg)
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
