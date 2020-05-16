FROM prakasa1904/tergit-base:latest AS build-env

WORKDIR /terpusat/fetch

COPY . .

RUN go mod tidy; \
    go build -ldflags="-w -s" -o git-fetch fetch/main.go;

FROM prakasa1904/tergit-run:latest

COPY --from=build-env /terpusat/fetch/git-fetch /go/bin/git-fetch

ENTRYPOINT ["/go/bin/git-fetch"]