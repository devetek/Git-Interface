FROM prakasa1904/tergit-base:latest AS build-env

WORKDIR /terpusat/clone

COPY . .

RUN go mod tidy; \
    go build -ldflags="-w -s" -o git-clone clone/main.go;

FROM prakasa1904/tergit-run:latest

COPY --from=build-env /terpusat/clone/git-clone /go/bin/git-clone

ENTRYPOINT ["/go/bin/git-clone"]