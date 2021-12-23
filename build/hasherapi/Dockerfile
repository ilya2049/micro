FROM golang:1.17 AS builder
COPY ./hasherapi /go/src/hasherapi
COPY ./common /go/src/common
WORKDIR /go/src/hasherapi/cmd/hasherapi

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-w -s' -o /go/bin/hasherapi

FROM alpine:3.15.0
COPY --from=builder /go/bin/hasherapi /go/bin/hasherapi
# TODO: pass host/port via configuration
CMD /go/bin/hasherapi --port 8080 --host '0.0.0.0'
