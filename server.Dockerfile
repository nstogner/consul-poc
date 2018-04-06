FROM golang:alpine as build
COPY . /go/src/github.com/nstogner/consul-poc
RUN go install github.com/nstogner/consul-poc/cmd/server

FROM alpine
COPY --from=build /go/bin/server /server

ENTRYPOINT ["/server"]
