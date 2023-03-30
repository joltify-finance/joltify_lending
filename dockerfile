FROM golang:1.18-alpine as builder
# install make tools
RUN apk add --no-cache make git gcc musl-dev linux-headers


WORKDIR /go/src/github.com/go
#copy the source code to the container
COPY . ./joltify_lending
# enter the directory of joltify_lending
WORKDIR /go/src/github.com/go/joltify_lending
# run the makefile install
RUN make install
## copy the binary to the alpine image
FROM golang:1.18
WORKDIR /root/
COPY --from=builder /go/bin/joltify_lending /usr/bin/joltify_lending
EXPOSE 26656 26657 1317 9090
#

