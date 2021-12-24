FROM golang:1.17 AS builder
COPY ./internal/hasher /go/src/hasher
COPY ./internal/common /go/src/common
WORKDIR /go/src/hasher/cmd/hasher

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-w -s' -o /go/bin/hasher

FROM alpine:3.15.0
COPY --from=builder /go/bin/hasher /go/bin/hasher

CMD /go/bin/hasher
