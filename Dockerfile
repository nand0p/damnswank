FROM golang:alpine

MAINTAINER "nando" <nando@hex7.com>

ARG DATE
ARG REVISION

RUN apk add --no-cache git

ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -pv ${GOPATH}/src ${GOPATH}/bin

RUN go version
RUN go get github.com/braintree/manners
RUN go get github.com/kelseyhightower/app/health

COPY . ${GOPATH}
WORKDIR ${GOPATH}

RUN sed -i "s|SEDME|$REVISION -- $DATE|g" handlers/damnswank.go
RUN cat handlers/damnswank.go
RUN go build -o ${GOPATH}/server .

ENTRYPOINT ["/go/server"]
