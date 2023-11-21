FROM golang:alpine AS builder

WORKDIR /build
ADD . .

RUN make build

FROM alpine:latest

RUN apk update && \
  apk add --no-cache tzdata
RUN mkdir /etc/access-control/
COPY --from=builder /build/bin /bin/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
