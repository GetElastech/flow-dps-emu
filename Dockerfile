# NOTE: Must be run in the context of the repo's root directory

## (1) Download a suitable version of Go
FROM golang:1.18 AS build-setup

# Add optional items like apt install -y make cmake gcc g++
RUN apt update && apt install -y cmake make gcc g++

## (2) Build the app binary
FROM build-setup AS build-dependencies

# Cache gopath dependencies for faster builds
# Newer projects should opt for go mod vendor for reliability and security
RUN mkdir /app
RUN mkdir /app/src
COPY src/go.mod /app/src
COPY src/go.sum /app/src
RUN mkdir -p /app/upstream/flow-go
COPY upstream/flow-go/go.mod /app/upstream/flow-go
COPY upstream/flow-go/go.sum /app/upstream/flow-go

# FIX: This generates code marked by `go:build relic` and `+build relic`. See `combined_verifier_v3.go`.
# FIX: This is not needed, if vendor/ is used
# NOTE: crypto@v0.24.4 is the latest stable version from flow-go
WORKDIR /app/src
RUN go mod download
RUN go mod download github.com/onflow/flow-go/crypto@v0.24.4
RUN cd $GOPATH/pkg/mod/github.com/onflow/flow-go/crypto@v0.24.4 && go generate && go build

## (3) Build the app binary
FROM build-dependencies AS build-env

COPY src /app/src
COPY upstream /app/upstream
WORKDIR /app/src

# Fix: make sure no further steps update modules later, so that we can debug regressions
RUN go mod vendor
RUN cp -R $GOPATH/pkg/mod/github.com/onflow/flow-go/crypto@v0.24.4/* /app/src/vendor/github.com/onflow/flow-go/crypto
RUN ls /app/src/vendor/github.com/onflow/flow-go/crypto/relic

# FIX: Without -tags=relic we get undefined: "github.com/onflow/flow-go/consensus/hotstuff/verification".NewCombinedVerifier
RUN go build -v -tags=relic -o /app cmd/flow-dps-emu/main.go

CMD /bin/bash

## (5) Add the statically linked binary to a distroless image
FROM build-env as production

WORKDIR /app/src
COPY --from=build-env /app/main /app/flow-dps-emu

CMD ["/app/flow-dps-emu"]

## (6) Add the statically linked binary to a distroless image
FROM golang:1.18 as production-small

RUN rm -rf /go
RUN rm -rf /app
RUN rm -rf /usr/local/go
COPY --from=production /app/flow-dps-emu /bin/flow-dps-emu

CMD ["/bin/main"]
