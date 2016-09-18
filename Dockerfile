FROM centos:7

MAINTAINER "Fernando J Pando" <nando@hex7.com>

RUN yum -y install git

RUN curl https://storage.googleapis.com/golang/go1.5.linux-amd64.tar.gz | tar -C /usr/local -xzf -

ENV PATH ${PATH}:/usr/local/go/bin

RUN go version

ENV GOPATH /usr/local/go/src

RUN go get github.com/braintree/manners

RUN go get github.com/kelseyhightower/app/health

RUN go get github.com/nand0p/app/handlers

RUN git clone https://github.com/nand0p/damnswank.git /damnswank

WORKDIR /damnswank

RUN go build -o /damnswank/server .

ENTRYPOINT ["/damnswank/server"]
