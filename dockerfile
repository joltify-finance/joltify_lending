FROM golang:1.20

WORKDIR /go/src/app
COPY . .
RUN apt-get update;apt-get install -y jq bc vim
RUN make install;make integration-test
