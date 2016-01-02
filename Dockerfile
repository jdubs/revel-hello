FROM golang:latest

ADD . /go/src/github.com/jdubs/revel-hello

RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

ENTRYPOINT revel run github.com/jdubs/revel-hello

EXPOSE 9000
