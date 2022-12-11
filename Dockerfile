FROM alpine:edge

RUN echo 'http://dl-cdn.alpinelinux.org/alpine/edge/testing' >> /etc/apk/repositories && \
    apk update && apk add \
      go \
      mage

WORKDIR /usr/src/go-playground
COPY . /usr/src/go-playground

ENV GOPATH /root/go
ENV PATH ${GOPATH}/bin:/usr/local/go/bin:$PATH
ENV GOBIN $GOROOT/bin
ENV GO111MODULE on

RUN go get github.com/go-git/go-git/v5 && \
    go get -u github.com/spf13/cobra@latest && \
    mkdir -p ${GOPATH}/src ${GOPATH}/bin && \
    mage build
