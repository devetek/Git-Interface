.PHONY: clone-test
clone-test:
	@rm -rf ./git-kaniko
	( \
		export GITHUB_URL=https://github.com/devetek/Simple-Kaniko.git; \
		export GITHUB_DIRECTORY=./git-kaniko; \
		export GITHUB_USERNAME=prakasa1904; \
		go run clone/main.go; \
	)

.PHONY: clone-build
clone-build:
	@go build -ldflags="-w -s" -o git-clone clone/main.go

.PHONY: clone-create-image
clone-create-image:
	@docker build -f clone/Dockerfile -t prakasa1904/git-clone:latest .
	@docker push prakasa1904/git-clone:latest

.PHONY: clone-run-image
clone-run-image:
	@docker run --env GITHUB_USERNAME=prakasa1904 --env GITHUB_DIRECTORY=./git-kaniko-docker --env GITHUB_URL=https://github.com/devetek/Simple-Kaniko.git -v `pwd`:/terpusat prakasa1904/git-clone:latest