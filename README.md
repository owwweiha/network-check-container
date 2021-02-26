# network-check-container
Small REST API written in go packed into a container to check network connectivity (host/port/protocol).

## Installation

To build the image using podman, simply run
```bash
podman build -t registry.localhost/network-check:0.1 -f ./Containerfile
```

## Usage

To run the container locally, execute

```bash
podman run --network=host -dit registry.localhost/network-check:0.1
```

Example 1: Connection successful

Request:
```bash
curl -X GET -H "Content-type: application/json" -H "Accept: application/json" -d '{"host": "google.com", "port": "443", "protocol": "tcp" }' localhost:8080/api/v1/connect
```
Result: `{"status": "success", "host": "google.com", "port": "443", "protocol": "tcp", "message": "connection successful"}`

Example 2: Connection closed

Request:
```bash
curl -X GET -H "Content-type: application/json" -H "Accept: application/json" -d '{"host": "google.com", "port": "4443", "protocol": "tcp" }' localhost:8080/api/v1/connect
```
Result: `{"status": "error", "host": "google.com", "port": "4443", "protocol": "tcp", "message": "dial tcp 172.217.16.142:4443: i/o timeout"}`

Example 3: Bad request (missing protocol)

Request:
```bash
curl -X GET -H "Content-type: application/json" -H "Accept: application/json" -d '{"host": "google.com", "port": "443" }' localhost:8080/api/v1/connect
```
Result: `{"status": "error", "message": "bad request"}`
