FROM golang:1.8.1

COPY . /go/src/github.com/aofry/go-tee/tracer
RUN go get github.com/vulcand/oxy/trace
RUN go install github.com/aofry/go-tee/tracer

ENTRYPOINT ["/go/bin/tracer"]
EXPOSE 8080