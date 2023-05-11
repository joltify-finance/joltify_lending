FROM golang:1.18

WORKDIR /go/src/app
COPY . .
RUN apt-get update;apt-get install -y jq bc
RUN make install



