FROM alpine:latest

WORKDIR /terpusat

RUN apk update && apk upgrade; \
    apk add --no-cache bash git openssh; \
    rm -rf /var/cache/apk/*;
