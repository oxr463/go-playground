FROM alpine:edge

RUN apk update && \
    apk add go \
            mage

WORKDIR /usr/src/go-playground
COPY . /usr/src/go-playground

RUN mage
