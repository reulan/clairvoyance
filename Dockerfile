FROM golang:alpine3.12 as clairvoyance_build

RUN apk add --no-cache git

# Download and install pre-reqs for clairvoyance
WORKDIR /clairvoyance-build
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy code + build application
# (tftest is optional)
COPY app/ app/
COPY cmd/ cmd/
COPY config/ config/
COPY log/ log/
COPY tftest/ tftest/
COPY version/ version/
COPY main.go .
RUN go build -o ./bin/clairvoyance .

# Use only alpine this time around
FROM alpine:3.12
LABEL REPO="https://github.com/reulan/clairvoyance"
LABEL maintainer="mpmsimo@gmail.com"

RUN apk add ca-certificates
RUN apk add terraform

# Create and use the clairvoyance user
RUN adduser -D -g '' clairvoyance
USER clairvoyance
WORKDIR /app
COPY --chown=clairvoyance:clairvoyance --from=clairvoyance_build /clairvoyance-build/bin/clairvoyance .
COPY --chown=clairvoyance:clairvoyance tftest/ /app/tftest/

# Run the API
#use this env var${CLAIRVOYANCE_PROJECT_DIR}
WORKDIR /app/tftest/drift
CMD ["/app/clairvoyance", "report", "discord"]
