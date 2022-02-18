package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

const dnsPort int = 53
const httpPort int = 8080

func main() {
	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		log.Fatal("Set API_KEY environment variable")
	}
	startHttpServer(apiKey)
	startDnsServer()
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<- sig
	log.Println("Exiting...")
	stopDnsServer()
}