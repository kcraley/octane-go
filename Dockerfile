# Create based container to do initial compilation
FROM golang:1.14 as base
ARG SOURCE=/go/src/github.com/kcraley/octane-go
COPY . ${SOURCE}
RUN set -ex \
    && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make -C ${SOURCE} build \
    && ls -lah ${SOURCE}

# Copy the compiled binary into a lightweight container
FROM alpine:latest
ARG SOURCE=/go/src/github.com/kcraley/octane-go
RUN set -ex \
    && apk --no-cache add ca-certificates \
    && apk --no-cache upgrade
COPY --from=base ${SOURCE}/octane /usr/local/bin
ENTRYPOINT [ "/usr/local/bin/octane" ]
