FROM golang:alpine as build
COPY . /go/src/github.com/nstogner/consul-poc
RUN go install github.com/nstogner/consul-poc/cmd/server
RUN go install github.com/nstogner/consul-poc/cmd/client

FROM alpine
COPY --from=build /go/bin/server /server
COPY --from=build /go/bin/client /client

