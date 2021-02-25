# network-check-container
Small go app packed into a container to check network connectivity (host/port).

## Installation

To build the image using podman, simply run
```bash
podman build -t registry.localhost/network-check:0.1 -f ./Containerfile
```

## Usage

```bash
podman run -it registry.localhost/network-check:0.1 google.com 443
```
will result in `Connection succesful: google.com:443`
```bash
podman run -it registry.localhost/network-check:0.1 google.com 4443
```
will result in `Connecting failed: dial tcp 172.217.16.142:4443: i/o timeout`
