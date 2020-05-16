include clone/Makefile
include fetch/Makefile
include pull/Makefile
include push/Makefile
include checkout/Makefile

DOCKER_USER := prakasa1904
DOCKER_REG_BUILD := tergit-base
DOCKER_TAG_BUILD := latest
DOCKER_REG_RUN := tergit-run
DOCKER_TAG_RUN := latest

.PHONY: test
test:
	@go test ./...

.PHONY: image-build
image-build:
	@docker build -f base.Dockerfile -t $(DOCKER_USER)/$(DOCKER_REG_BUILD):$(DOCKER_TAG_BUILD) .
	@docker push $(DOCKER_USER)/$(DOCKER_REG_BUILD):$(DOCKER_TAG_BUILD)

.PHONY: image-run
image-run:
	@docker build -f run.Dockerfile -t $(DOCKER_USER)/$(DOCKER_REG_RUN):$(DOCKER_TAG_RUN) .
	@docker push $(DOCKER_USER)/$(DOCKER_REG_RUN):$(DOCKER_TAG_RUN)

.PHONY: build
build:
	@go build -ldflags="-w -s" -o git-clone clone/main.go
	@go build -ldflags="-w -s" -o git-fetch fetch/main.go
	@go build -ldflags="-w -s" -o git-pull pull/main.go
	@go build -ldflags="-w -s" -o git-checkout checkout/main.go