FROM golang:1.17 AS builder
ARG PORT

COPY ./internal/hasherapi /go/src/hasherapi
COPY ./internal/common /go/src/common
WORKDIR /go/src/hasherapi/cmd/hasherapi

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-w -s' -o /go/bin/hasherapi

FROM alpine:3.15.0
ARG PORT
ENV PORT ${PORT}

COPY --from=builder /go/bin/hasherapi /go/bin/hasherapi

CMD /go/bin/hasherapi --port ${PORT} --host '0.0.0.0'
