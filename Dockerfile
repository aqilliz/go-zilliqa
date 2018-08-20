# Build Gzil in a stock Go builder container
FROM golang:1.10-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go-zilliqa
RUN cd /go-zilliqa && make gzil

# Pull Gzil into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-zilliqa/build/bin/gzil /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["gzil"]
