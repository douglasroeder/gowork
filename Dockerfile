# build stage
FROM golang:alpine AS build-env
ADD . /go/src/github.com/douglasroeder/gowork
RUN apk update \
    && apk add --virtual build-dependencies \
        build-base \
        gcc
RUN go install github.com/douglasroeder/gowork

# final stage
FROM alpine
RUN apk add --update \
  ca-certificates
RUN mkdir -p /opt/douglasroeder
WORKDIR /opt/douglasroeder
COPY --from=build-env /go/bin/gowork /opt/douglasroeder/gowork
CMD ["/opt/douglasroeder/gowork"]