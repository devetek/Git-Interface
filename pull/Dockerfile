FROM prakasa1904/tergit-base:latest AS build-env

WORKDIR /terpusat/pull

COPY . .

RUN go mod tidy; \
    go build -ldflags="-w -s" -o git-pull pull/main.go;

FROM prakasa1904/tergit-run:latest

COPY --from=build-env /terpusat/pull/git-pull /go/bin/git-pull

ENTRYPOINT ["/go/bin/git-pull"]