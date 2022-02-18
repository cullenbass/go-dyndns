FROM golang:alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build

FROM alpine:latest
COPY --from=builder /app/go-dyndns go-dyndns
EXPOSE 8080
EXPOSE 53
EXPOSE 53/udp
CMD ./go-dyndns