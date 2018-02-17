# Build: docker build -t local/json2hcl .
# Run:   docker run -i --rm -v $PWD:/work local/json2hcl --reverse < fixtures/infra.tf
# Run:   docker run -i --rm -v $PWD:/work local/json2hcl < fixtures/infra.tf.json

FROM golang:alpine

RUN apk add --update git 

WORKDIR $GOPATH/src/github.com/kvz/json2hcl
COPY ./main.go .
RUN go get -v && go install -v

WORKDIR /work
ENTRYPOINT ["json2hcl"]