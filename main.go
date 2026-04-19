package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mostlygeek/llama-swap/proxy"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var (
		configFile  = flag.String("config", "config.yaml", "path to config file")
		listenAddr  = flag.String("listen", "127.0.0.1:8080", "address to listen on")
		showVersion = flag.Bool("version", false, "show version information")
		logLevel    = flag.String("log-level", "info", "log level (debug, info, warn, error)")
	)
	flag.Parse()

	if *showVersion {
		fmt.Printf("llama-swap version %s (commit: %s, built: %s)\n", version, commit, date)
		os.Exit(0)
	}

	log.Printf("llama-swap %s starting", version)
	log.Printf("Loading config from: %s", *configFile)

	cfg, err := proxy.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := cfg.Validate(); err != nil {
		log.Fatalf("Invalid config: %v", err)
	}

	log.Printf("Loaded %d model(s)", len(cfg.Models))

	p, err := proxy.New(cfg, *logLevel)
	if err != nil {
		log.Fatalf("Failed to create proxy: %v", err)
	}

	log.Printf("Listening on %s", *listenAddr)
	if err := p.ListenAndServe(*listenAddr); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
