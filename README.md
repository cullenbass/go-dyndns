# go-dyndns

1. Set the NS record for the (sub)domain to the IP address of the server
2. `go build`
3. Set `API_KEY` environment variable
4. `go-dyndns`
5. Create DNS records using the `/updateDomain` REST endpoint;
    - Requires `domain`, `apiKey` in key-value pairs using query strings or form data
    - `ipAddress` is also accepted in query strings or form data. If absent, the domain is updated with the origination address of the request
    - ie: `curl http://localhost:8080/updateDomain?apiKey=Ex4ample&domain=dynamic.example.com&ipAddress=256.256.256.256`