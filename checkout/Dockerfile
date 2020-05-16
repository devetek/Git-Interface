FROM prakasa1904/tergit-base:latest AS build-env

WORKDIR /terpusat/checkout

COPY . .

RUN go mod tidy; \
    go build -ldflags="-w -s" -o git-checkout checkout/main.go;

FROM prakasa1904/tergit-run:latest

COPY --from=build-env /terpusat/checkout/git-checkout /go/bin/git-checkout

ENTRYPOINT ["/go/bin/git-checkout"]