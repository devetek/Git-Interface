.PHONY: fetch-test
fetch-test:
	( \
		export GITHUB_DIRECTORY=./git-kaniko; \
		export GITHUB_USERNAME=prakasa1904; \
		export GITHUB_BRANCH=production; \
		go run fetch/main.go; \
	)

.PHONY: fetch-build
fetch-build:
	@go build -ldflags="-w -s" -o git-fetch fetch/main.go

.PHONY: fetch-create-image
fetch-create-image:
	@docker build -f fetch/Dockerfile -t prakasa1904/git-fetch:latest .
	@docker push prakasa1904/git-fetch:latest

.PHONY: fetch-run-image
fetch-run-image:
	@docker run --env GITHUB_USERNAME=prakasa1904 --env GITHUB_BRANCH=production --env GITHUB_DIRECTORY=./git-kaniko-docker -v `pwd`:/terpusat prakasa1904/git-fetch:latest