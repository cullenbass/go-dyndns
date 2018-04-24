package main

import (
	"net/http"
	"log"
	"strconv"
)

func startHttpServer() {
	http.HandleFunc("/updateDomain", updateHandler)
	go func(){
		log.Printf("Setting up HTTP server on port %d\n", httpPort)
		if err := http.ListenAndServe(":" + strconv.Itoa(httpPort), nil); err != nil {
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
		w.WriteHeader(400)
		w.Write([]byte("Bad address"))
		return
	}
	success := addZone(domain, ipAddress)
	if !success {
		w.WriteHeader(500)
		w.Write([]byte(nil))
		return
	}

	w.Write([]byte(nil))
}