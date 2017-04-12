FROM golang:1.8.1
ENV PROXY_SERVICE=http://localhost:63450

COPY . /go/src/github.com/aofry/go-tee
RUN go get github.com/vulcand/oxy/forward github.com/vulcand/oxy/testutils github.com/mattn/goveralls
RUN go install github.com/aofry/go-tee
RUN go test -cover github.com/aofry/go-tee/util github.com/aofry/go-tee/tee

ENTRYPOINT ["/go/bin/go-tee"]
EXPOSE 8080