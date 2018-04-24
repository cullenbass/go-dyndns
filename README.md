# gd-dns

1. Set the NS record for the (sub)domain to the IP address of the server
2. Clone the repo, then change the API key, HTTP port, or DNS port in `main.go`
3. `go get github.com/miekg/dns` 
4. `go build`
5. `go-dyndns`
6. Create DNS records using the `/updateDomain` REST endpoint;
    - Requires `domain`, `apiKey`, and `ipAddress` in key-value pairs using query strings or form data
    - ie: `curl http://localhost:8080/updateDomain?apiKey=Ex4ample&domain=dynamic.example.com&ipAddress=256.256.256.256`