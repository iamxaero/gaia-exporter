package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"


	"example.com/gaia-exporter/config"
	"example.com/gaia-exporter/controller"


	"context"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/cloudflare/cfssl/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


func main() {
	// Config
	cfg := config.New()
	ctrl := controller.New(cfg)

	// Router
	h2s := &http2.Server{}
	handler := http.NewServeMux()
	// Prometheus register metrics
	ctrl.PromRegister()
	// Handlers
	// handler.HandleFunc("/webhook", ctrl.Webhook)
	handler.HandleFunc("/", ctrl.Health)
	handler.HandleFunc("/health", ctrl.Health)
	handler.Handle("/metrics", promhttp.Handler())
	// Set option for server
	listen := cfg.GaiaPort
	server := &http.Server{
		Addr:         listen,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 180 * time.Second,
		IdleTimeout:  240 * time.Second,
		Handler:      h2c.NewHandler(handler, h2s),
	}
	// Start Collector
	go func() {
		for {
			// Get gaia status
			cmd := exec.Command("/Users/vk/.pyenv/versions/3.12.5/envs/3/bin/python", cfg.GaiaBin)
			output, err := cmd.Output()
			if err != nil {
				fmt.Printf("Gaia status error: %v", err)
				continue
			}
	
			// Map output
			var status controller.GaiaStatus
			err = json.Unmarshal(output, &status)
			if err != nil {
				fmt.Printf("Parse JSON error: %v", err)
				continue
			}
	
			// Debug struct print
			ctrl.ProcGaiaStatus(status)
			fmt.Printf("Parsed Status:\n%+v\n", status)
	
			// Ожидание перед следующим запуском
			time.Sleep(10 * time.Second)
		}
		}()

	// Start http server
	go func() {
		log.Infof("Running server at %v", listen)
		log.Fatal(server.ListenAndServe())
	}()

	// Wait for an interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = server.Shutdown(ctx)
}
