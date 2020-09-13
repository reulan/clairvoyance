.PHONY: build build-alpine clean test help default

BIN_NAME=clairvoyance

VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := "reulan/clairvoyance"

TFINSTALL_DIR="./tfinstall/terraform_${CLAIRVOYANCE_TERRAFORM_VERSION}"

default: test


# Clairvoyance build
deps:
	rm go.mod go.sum || true
	go mod init || true
	go mod tidy || true

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}
	
build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X github.com/reulan/clairvoyance/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/reulan/clairvoyance/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}


build-alpine:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags '-w -linkmode external -extldflags "-static" -X github.com/reulan/clairvoyance/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/reulan/clairvoyance/version.BuildDate=${BUILD_DATE}' -o bin/${BIN_NAME}


package:
	@echo "building image ${BIN_NAME} ${VERSION} $(GIT_COMMIT)"
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=$(GIT_COMMIT) -t $(IMAGE_NAME):local .

	
# Clairvoyance validation
check-env-vars:
	@if [ -z "${CLAIRVOYANCE_WORKING_DIR}" ]; then echo "Missing CLAIRVOYANCE_WORKING_DIR"; exit 1; fi
	@if [ -z "${DISCORD_WEBHOOK_CHANNEL}" ]; then echo "Missing DISCORD_WEBHOOK_CHANNEL"; exit 1; fi
	@if [ -z "${DISCORD_WEBHOOK_SECRET}" ]; then echo "Missing DISCORD_WEBHOOK_SECRET"; exit 1; fi


# Clairvoyance runtime
# this installs terraform version specified by the env var
tfinstall:
	@if [ -z "${CLAIRVOYANCE_TERRAFORM_VERSION}" ]; then echo "Missing CLAIRVOYANCE_TERRAFORM_VERSION"; exit 1; fi
	rm -rf ./tfinstall || true
	mkdir -p $(TFINSTALL_DIR) || true
	wget https://releases.hashicorp.com/terraform/${CLAIRVOYANCE_TERRAFORM_VERSION}/terraform_${CLAIRVOYANCE_TERRAFORM_VERSION}_linux_amd64.zip -P $(TFINSTALL_DIR)
	unzip $(TFINSTALL_DIR)/terraform_${CLAIRVOYANCE_TERRAFORM_VERSION}_linux_amd64.zip -d $(TFINSTALL_DIR)
	
report-stdout: build check-env-vars
	@echo "running ${BIN_NAME} ${VERSION}"
	./bin/clairvoyance report stdout

report-discord: build check-env-vars
	@echo "running ${BIN_NAME} ${VERSION}"
	./bin/clairvoyance report discord


# Docker
tag: 
	@echo "Tagging: latest ${VERSION} $(GIT_COMMIT)"
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):$(GIT_COMMIT)
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):${VERSION}
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):latest

push: tag
	@echo "Pushing docker image to registry: latest ${VERSION} $(GIT_COMMIT)"
	docker push $(IMAGE_NAME):$(GIT_COMMIT)
	docker push $(IMAGE_NAME):${VERSION}
	docker push $(IMAGE_NAME):latest
