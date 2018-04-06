package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/kelseyhightower/envconfig"
	"github.com/nstogner/consul-poc/internal/backoff"
)

func main() {
	var cfg struct {
		ServiceName string   `envconfig:"service_name" default:"abc"`
		ServiceTags []string `envconfig:"service_tags"`
		Port        int      `default:"7000"`
		ConsulAddr  string   `envconfig:"consul_addr" default:"localhost:8500"`
	}
	envconfig.MustProcess("", &cfg)

	if err := backoff.Retry(3, time.Second, func() error {
		log.Printf("[server] Registering with consul at address %q", cfg.ConsulAddr)
		return registerWithConsul(cfg.ConsulAddr, cfg.ServiceName, cfg.ServiceTags, cfg.Port)
	}); err != nil {
		log.Fatal("[server] Unable to register service with consul: ", err)
	}

	log.Println("[server] Starting")

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%v", cfg.Port),
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[server] Received request from %q", r.RemoteAddr)

			w.WriteHeader(200)
			w.Write([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
		})),
	)
}

func registerWithConsul(consulAddr, serviceName string, serviceTags []string, servicePort int) error {
	c, err := consulapi.NewClient(&consulapi.Config{
		Address: consulAddr,
	})
	if err != nil {
		return err
	}

	return c.Agent().ServiceRegister(&consulapi.AgentServiceRegistration{
		Name: serviceName,
		Tags: serviceTags,
		Port: servicePort,
	})
}
