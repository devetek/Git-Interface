.PHONY: pull-test
pull-test:
	( \
		export GITHUB_DIRECTORY=./git-kaniko; \
		export GITHUB_USERNAME=prakasa1904; \
		go run pull/main.go; \
	)

.PHONY: pull-build
pull-build:
	@go build -ldflags="-w -s" -o git-pull pull/main.go

.PHONY: pull-create-image
pull-create-image:
	@docker build -f pull/Dockerfile -t prakasa1904/git-pull:latest .
	@docker push prakasa1904/git-pull:latest

.PHONY: pull-run-image
pull-run-image:
	@docker run --env GITHUB_USERNAME=prakasa1904 --env GITHUB_DIRECTORY=./git-kaniko-docker -v `pwd`:/terpusat prakasa1904/git-pull:latest