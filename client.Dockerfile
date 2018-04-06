FROM golang:alpine as build
COPY . /go/src/github.com/nstogner/consul-poc
RUN go install github.com/nstogner/consul-poc/cmd/client

FROM alpine
COPY --from=build /go/bin/client /client

ENTRYPOINT ["/client"]
