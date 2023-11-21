FROM golang:alpine AS builder

# for running Makefile and builds
RUN apk update && apk upgrade && apk add --no-cache bash make git npm gcc musl-dev

WORKDIR /build
ADD . .

RUN make build

FROM alpine:latest

RUN apk update && \
  apk add --no-cache tzdata
COPY --from=builder /build/bin /bin/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "/bin/http" ]