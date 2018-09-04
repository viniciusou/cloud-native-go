FROM golang:1.10.3-alpine3.8
MAINTAINER  Vinicius Uemura

ENV SOURCES /go/src/github.com/viniciusou/cloud-native-go/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT cloud-native-go
