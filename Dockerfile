FROM centos:7

MAINTAINER "nando" <nando@hex7.com>

ARG DATE
ARG REVISION

RUN yum install git -y
RUN curl https://storage.googleapis.com/golang/go1.5.linux-amd64.tar.gz | tar -C /usr/local -xzf -

ENV PATH ${PATH}:/usr/local/go/bin
ENV GOPATH /usr/local/go/vendor

RUN mkdir -pv /usr/local/go/vendor
RUN mkdir -pv /damnswank

RUN go version

RUN go get github.com/braintree/manners
RUN go get github.com/kelseyhightower/app/health

COPY . /damnswank
WORKDIR /damnswank
RUN sed -i "s|SEDME|$REVISION -- $DATE|g" handlers/damnswank.go
RUN cat handlers/damnswank.go
RUN go build -o /damnswank/server .

ENTRYPOINT ["/damnswank/server"]
