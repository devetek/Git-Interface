.PHONY: checkout-test
checkout-test:
	( \
		export GITHUB_DIRECTORY=./git-kaniko; \
		export GITHUB_BRANCH=master; \
		go run checkout/main.go; \
	)

.PHONY: checkout-build
checkout-build:
	@go build -ldflags="-w -s" -o git-checkout checkout/main.go

.PHONY: checkout-create-image
checkout-create-image:
	@docker build -f checkout/Dockerfile -t prakasa1904/git-checkout:latest .
	@docker push prakasa1904/git-checkout:latest

.PHONY: checkout-run-image
checkout-run-image:
	@docker run --env GITHUB_BRANCH=master --env GITHUB_DIRECTORY=./git-kaniko-docker -v `pwd`:/terpusat prakasa1904/git-checkout:latest