.PHONY: check-env-vars tfinstall deps clean build stdout discord docker push shell test

BIN="clairvoyance"
VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_SHORT=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE := "gcr.io/${BIN}"
TFINSTALL_DIR="./tfinstall/terraform_${CLAIRVOYANCE_TERRAFORM_VERSION}"

# Clairvoyance validation
check-env-vars:
	@if [ -z "${CLAIRVOYANCE_PROJECT_DIR}" ]; then echo "Missing CLAIRVOYANCE_PROJECT_DIR"; exit 1; fi
	@if [ -z "${DISCORD_WEBHOOK_CHANNEL}" ]; then echo "Missing DISCORD_WEBHOOK_CHANNEL"; exit 1; fi
	@if [ -z "${DISCORD_WEBHOOK_SECRET}" ]; then echo "Missing DISCORD_WEBHOOK_SECRET"; exit 1; fi
	
# this installs terraform version specified by the env var
# eventually to be replaced by tfinstall via Go code, but alas this is easier
tfinstall:
	@if [ -z "${CLAIRVOYANCE_TERRAFORM_VERSION}" ]; then echo "Missing CLAIRVOYANCE_TERRAFORM_VERSION"; exit 1; fi
	rm -rf ./tfinstall || true
	mkdir -p $(TFINSTALL_DIR) || true
	wget https://releases.hashicorp.com/terraform/${CLAIRVOYANCE_TERRAFORM_VERSION}/terraform_${CLAIRVOYANCE_TERRAFORM_VERSION}_linux_amd64.zip -P $(TFINSTALL_DIR)
	unzip $(TFINSTALL_DIR)/terraform_${CLAIRVOYANCE_TERRAFORM_VERSION}_linux_amd64.zip -d $(TFINSTALL_DIR)

# Install Golang dependancies
deps:
	rm go.mod go.sum || true
	go mod init || true
	go mod tidy || true

clean:
	@test ! -e bin/${BIN} || rm bin/${BIN}

# Building the binary
build:
	@echo "building Go binary: ${BIN}"
	go build -o bin/${BIN} main.go
	#go build -ldflags '-w -linkmode external -extldflags "-static" -X github.com/reulan/clairvoyance/version.GitCommit=${GIT_SHORT} -X github.com/reulan/clairvoyance/version.BuildDate=${BUILD_DATE}' -o bin/${BIN}

### Clairvoyance invocation
stdout: build check-env-vars
	@echo "running ${BIN} ${VERSION}"
	./bin/${BIN} report stdout

discord: build check-env-vars
	@echo "running ${BIN} ${VERSION}"
	./bin/${BIN} report discord

### Docker bits below
# This will invoke clairvoyance via Docker
docker: build
	docker build -t $(IMAGE):$(GIT_SHORT) .
	docker tag $(IMAGE):$(GIT_SHORT) $(IMAGE):latest
	docker tag $(IMAGE):$(GIT_SHORT) $(IMAGE):${VERSION}
	@echo "built Docker image: $(IMAGE):$(GIT_SHORT)/${VERSION}"

# TODO: configure this for public DockerHub instead of private GCR.
push: docker
	@echo "Uploading image ${IMAGE}:$(GIT_SHORT)/${VERSION} to GCR."
	docker push $(IMAGE):latest
	docker push $(IMAGE):${VERSION}
	docker push $(IMAGE):$(GIT_SHORT)

# Start a container with the Clairvoyance binary and enter the Alpine Shell. (debug)
shell: docker
	docker run --rm -it \
		-e CLAIRVOYANCE_TERRAFORM_VERSION=${CLAIRVOYANCE_TERRAFORM_VERSION} \
		-e CLAIRVOYANCE_PROJECT_DIR="/app/tftest/drift" \
		$(IMAGE):$(GIT_SHORT) ash

# Run clairvoyance in a Docker container against the `tftest` resources.
test: docker
	docker run --rm -it \
		-e CLAIRVOYANCE_TERRAFORM_VERSION=${CLAIRVOYANCE_TERRAFORM_VERSION} \
		-e CLAIRVOYANCE_PROJECT_DIR="/app/tftest/drift" \
		$(IMAGE):$(GIT_SHORT)
