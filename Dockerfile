FROM golang:alpine3.12 AS builder
RUN apk add git
ENV CGO_ENABLED=0

########
# Prep
########

# add the source
COPY . /go/src/restapi
WORKDIR /go/src/restapi

########
# Build Go Wrapper
########

# Install go dependencies
RUN go get github.com/gorilla/mux && \
  go get github.com/go-resty/resty && \
  go get github.com/stretchr/testify

# build the package tests
RUN if [ ! -d "package.test" ] ; then echo Test package does not exist ; else rm-rf package.test ; fi
RUN go test -c . -o package.test

#build the go app
RUN go build -o ./restapi ./main.go

########
# Package into runtime image (smaller image size)
########
FROM alpine as release

# copy the executable from the builder image
COPY --from=builder /go/src/restapi .

CMD ["/bin/sh"]

EXPOSE 8080
