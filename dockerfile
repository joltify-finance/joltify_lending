FROM golang:1.18

WORKDIR /go/src/app
COPY . .

RUN make install



