# go-tee

[![Build Status](https://travis-ci.org/aofry/go-tee.svg?branch=master)](https://travis-ci.org/aofry/go-tee) [![Coverage Status](https://coveralls.io/repos/github/aofry/go-tee/badge.svg?branch=master)](https://coveralls.io/github/aofry/go-tee?branch=master)

<img src="https://storage.googleapis.com/gopherizeme.appspot.com/gophers/59b95b5c49448dff3581548e138cdb4fbf194036.png" width="120">

split traffic like a tee in order to send traffic to a debug system

Going over logs and metrics sometimes is just not enough. go-tee is basically here to help debug a production system. 
When you see a spcecific http error code happens in production you can fork it with harming the original traffic with go-tee and send the original http request to another system for debugging.


Building:
1. clone this repo: 
git clone git@github.com:aofry/go-tee.git
2. build with docker:
docker build -t aofry/go-tee .

Running:
docker run -d -p 8080:8080 -e REAL_BACKEND=http://server:port/path -e DEBUG_BACKEND=debug-server:port aofry/go-tee


A docker-compose.yml is also available as an example.
To run it:
1. clone this repo: 
git clone git@github.com:aofry/go-tee.git
2. run compose
docker-compose up --build
3. call the go-tee
curl localhost:8080/db