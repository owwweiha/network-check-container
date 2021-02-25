FROM golang:alpine AS golang-base
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY src/go.mod .
COPY src/check.go .

RUN go build -o network-check .
WORKDIR /dist
RUN cp /build/network-check .

FROM scratch
COPY --from=golang-base /dist/network-check /
ENTRYPOINT ["/network-check"]
