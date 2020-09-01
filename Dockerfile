# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-clairvoyance"
LABEL REPO="https://github.com/reulan/clairvoyance"

ENV PROJPATH=/go/src/github.com/reulan/clairvoyance

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/reulan/clairvoyance
WORKDIR /go/src/github.com/reulan/clairvoyance

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/reulan/clairvoyance"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/clairvoyance/bin

WORKDIR /opt/clairvoyance/bin

COPY --from=build-stage /go/src/github.com/reulan/clairvoyance/bin/clairvoyance /opt/clairvoyance/bin/
RUN chmod +x /opt/clairvoyance/bin/clairvoyance

# Create appuser
RUN adduser -D -g '' clairvoyance
USER clairvoyance

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/clairvoyance/bin/clairvoyance"]
