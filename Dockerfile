FROM golang:1.10-alpine as builder

RUN apk add --no-cache git make gcc musl-dev && \
    go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/tonyhhyip/seau

COPY . /go/src/github.com/tonyhhyip/seau

RUN make all

FROM alpine:3.6

ENV PLUGIN_PATH plugins
ENV PLUGINS php

COPY --from=builder /go/src/github.com/tonyhhyip/seau/dist/server /
COPY --from=builder /go/src/github.com/tonyhhyip/seau/plugins /plugins

CMD ["/server"]