package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/nstogner/consul-poc/internal/backoff"
)

func main() {
	var cfg struct {
		ServerAddr string `envconfig:"server_addr" default:"localhost:7000"`
	}
	envconfig.MustProcess("", &cfg)
	httpClient := &http.Client{Timeout: 3 * time.Second}

	log.Println("[client] Starting")

	for {
		time.Sleep(time.Second)

		if err := backoff.Retry(3, time.Second, func() error {
			log.Printf("[client] Requesting %q", cfg.ServerAddr)
			return req(httpClient, cfg.ServerAddr)
		}); err != nil {
			log.Fatalf("[client] Error: %s", err)
		}

		log.Println("[client] Success")
	}
}

func req(c *http.Client, addr string) error {
	res, err := c.Get(addr)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if exp, got := 200, res.StatusCode; exp != got {
		return fmt.Errorf("expected status code %v, got %v", exp, got)
	}

	return nil
}
