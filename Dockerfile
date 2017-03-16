FROM golang:1.8.0
ENV PROXY_SERVICE=http://localhost:63450

COPY . /go/src/github.com/aofry/go-tee
RUN go get github.com/vulcand/oxy/forward
RUN go install github.com/aofry/go-tee

ENTRYPOINT ["/go/bin/go-tee"]
EXPOSE 8080