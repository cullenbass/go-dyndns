package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

const dnsPort int = 53
const httpPort int = 8080
const apiKey string = "CHANGEME"

func main(){
	startHttpServer()
	startDnsServer()
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<- sig
	log.Println("Exiting...")
	stopDnsServer()
}