package main

import (
	"net/http"
	"log"
	"strconv"
	"regexp"
)

var apiKey string

func startHttpServer(key string) {
	apiKey = key
	http.HandleFunc("/updateDomain", updateHandler)
	go func(){
		log.Printf("Setting up HTTP server on port %d\n", httpPort)
		if err := http.ListenAndServe("0.0.0.0:" + strconv.Itoa(httpPort), nil); err != nil {
			log.Fatalf("Failed to set HTTP server %s\n", err.Error())
		}
	}()
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("apiKey") != apiKey {
		w.WriteHeader(403)
		w.Write([]byte(nil))
		return
	}
	domain := r.FormValue("domain")
	if domain == ""{
		w.WriteHeader(400)
		w.Write([]byte("Bad domain"))
		return
	}
	ipAddress := r.FormValue("ipAddress")
	if ipAddress == "" {
		ipRegex := regexp.MustCompile("(^.*)(?:\\:\\d*$)")
		ipAddress = ipRegex.FindStringSubmatch(r.RemoteAddr)[1]
	}
	success := addZone(domain, ipAddress)
	if !success {
		w.WriteHeader(500)
		w.Write([]byte(nil))
		return
	}

	w.Write([]byte(nil))
}