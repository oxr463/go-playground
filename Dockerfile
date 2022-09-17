FROM alpine:edge

RUN echo 'http://dl-cdn.alpinelinux.org/alpine/edge/testing' >> /etc/apk/repositories && \
    apk update && apk add \
      go \
      mage

WORKDIR /usr/src/go-playground
COPY . /usr/src/go-playground

RUN mage
