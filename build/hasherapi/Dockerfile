FROM golang:1.17 AS builder

WORKDIR /go/src

COPY internal/hasherapi/go.mod hasherapi/
COPY internal/hasherapi/go.sum hasherapi/

COPY internal/common/go.mod common/
COPY internal/common/go.sum common/

RUN cd hasherapi && go mod download
RUN cd common && go mod download

ARG PORT

COPY internal/hasherapi hasherapi
COPY internal/common common

RUN cd hasherapi/cmd/hasherapi && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build --ldflags='-w -s' -o /go/bin/hasherapi

FROM alpine:3.15.0
ARG PORT
ENV PORT ${PORT}

COPY --from=builder /go/bin/hasherapi /go/bin/hasherapi

CMD /go/bin/hasherapi --port ${PORT} --host '0.0.0.0'
