package main

import (
	"github.com/miekg/dns"
	"net"
	"log"
	"strconv"
)

var zones map[string]dns.A

var udpServer *dns.Server
var tcpServer *dns.Server
var sig chan bool

func removeZone(domain string) (bool) {
	domain = dns.Fqdn(domain)
	dns.HandleRemove(domain)
	return true
}

func addZone(domain string, ipAddress string) (bool) {
	domain = dns.Fqdn(domain)
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		log.Printf("DNS ERROR: IP address error: %s\n", ipAddress)
		return false
	}
	ar := new(dns.A)
	ar.Hdr = dns.RR_Header{
		Name: domain,
		Rrtype: dns.TypeA,
		Class: dns.ClassINET,
		Ttl: 60,
	}
	ar.A = ip
	rr, _ := dns.NewRR(ar.String())
	log.Printf("Updated DNS Entry: %s %s\n", domain, ipAddress)
	dns.HandleFunc(domain, func(w dns.ResponseWriter, r *dns.Msg) {
			m:= new(dns.Msg)
			m.SetReply(r)
			m.Authoritative = true
			m.Answer = []dns.RR{rr}
			w.WriteMsg(m)
		})
	return true
}

func startDnsServer() {
	udpServer = &dns.Server{Addr: ":" + strconv.Itoa(dnsPort), Net: "udp"}
	go func() {
		log.Printf("Setting up UDP DNS server on port %d\n", dnsPort)
		if err := udpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to set udp listener %s\n", err.Error())
		}
	}()	
	tcpServer = &dns.Server{Addr: ":" + strconv.Itoa(dnsPort), Net: "tcp"}
	go func(){
		log.Printf("Setting up TCP DNS server on port %d\n", dnsPort)
		if err := tcpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to set tcp listener %s\n", err.Error())
		}
		
	}()
}

func stopDnsServer() {
	if udpServer != nil {
		udpServer.Shutdown()
	}
	if tcpServer != nil {
		tcpServer.Shutdown()
	}
	
}