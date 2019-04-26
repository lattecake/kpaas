FROM golang:1.12.4 AS build-env

ENV GO111MODULE=on
ENV GO15VENDOREXPERIMENT=1
ENV BUILDPATH=github.com/lattecake/kpaas
RUN mkdir -p /go/src/${BUILDPATH}
COPY ./ /go/src/${BUILDPATH}
# ENV GOPROXY="http://artifactory/api/go/go-virtual"
RUN cd /go/src/${BUILDPATH} && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/kpaas-server

FROM alpine:latest

RUN apk add --no-cache \
		ca-certificates \
		tzdata \
		bash \
		&& cp -r -f /usr/share/zoneinfo/Hongkong /etc/localtime \
		&& rm -rf /var/cache/apk/*

COPY --from=build-env /go/bin/kpaas-server /go/bin/kpaas-server

WORKDIR /go/bin/
CMD ["/go/bin/kpaas-server"]