FROM golang:1.15 AS builder

WORKDIR /go/src/github.com/afghanistanyn/hydra-wework-auth-server
COPY ./ ./

RUN   go env -w GO111MODULE=on && \
      go env -w GOPROXY=https://goproxy.io && \
      go get -u github.com/ory/hydra-client-go@v1.9.0-alpha.2 && \
      go mod vendor && \
      mkdir -p vendor/github.com/ory/hydra/sdk/go/hydra/ && \
      cp -rf /go/pkg/mod/github.com/ory/hydra-client-go\@v1.9.0-alpha.2/*   vendor/github.com/ory/hydra/sdk/go/hydra/ && \
      GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags "-w -s" -o bin/hydra-wework-auth-server main.go

FROM alpine:edge
RUN apk add --update --no-cache ca-certificates && \
    mkdir -p /hydra-wework/conf && \
    mkdir -p /hydra-wework/bin && \
    mkdir -p /hydra-wework/logs

COPY --from=builder  /go/src/github.com/afghanistanyn/hydra-wework-auth-server/bin/hydra-wework-auth-server  /hydra-wework/bin/
COPY --from=builder  /go/src/github.com/afghanistanyn/hydra-wework-auth-server/conf/config.json.example  /hydra-wework/conf/config.json
WORKDIR /hydra-wework/
EXPOSE 8001
ENV GIN_MODE release
ENV APP_DEBUG false
ENTRYPOINT ["/hydra-wework/bin/hydra-wework-auth-server"]
CMD ["web","-a",":8001"]