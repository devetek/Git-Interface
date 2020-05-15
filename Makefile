test:
	@go test ./...

clone-test:
	@rm -rf ./git-kaniko
	( \
		export GITHUB_URL=https://github.com/devetek/Simple-Kaniko.git; \
		export GITHUB_DIRECTORY=./git-kaniko; \
		export GITHUB_USERNAME=prakasa1904; \
		go run clone/main.go; \
	)

fetch-test:
	( \
		export GITHUB_DIRECTORY=./git-kaniko; \
		export GITHUB_USERNAME=prakasa1904; \
		export GITHUB_BRANCH=production; \
		go run fetch/main.go; \
	)

pull-test:
	( \
		export GITHUB_DIRECTORY=./git-kaniko; \
		export GITHUB_USERNAME=prakasa1904; \
		go run pull/main.go; \
	)

checkout-test:
	( \
		export GITHUB_DIRECTORY=./git-kaniko; \
		export GITHUB_BRANCH=master; \
		go run checkout/main.go; \
	)

build:
	@go build -o git-clone clone/main.go
	@go build -o git-pull pull/main.go