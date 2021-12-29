FROM golang:1.17 AS builder

WORKDIR /go/src

COPY internal/hasher/go.mod hasher/
COPY internal/hasher/go.sum hasher/

COPY internal/common/go.mod common/
COPY internal/common/go.sum common/

RUN cd hasher && go mod download
RUN cd common && go mod download

COPY ./internal/hasher /go/src/hasher
COPY ./internal/common /go/src/common

RUN cd hasher/cmd/hasher && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-w -s' -o /go/bin/hasher

FROM alpine:3.15.0
COPY --from=builder /go/bin/hasher /go/bin/hasher

CMD /go/bin/hasher
