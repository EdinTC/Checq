FROM golang:1.12-alpine as builder

# To fix go get and build with cgo
RUN apk add --no-cache --virtual .build-deps \
    bash \
    gcc \
    git \
    musl-dev

# We create an /app directory within our
# image that will hold our application source
# files
RUN mkdir /app
# We copy everything in the root directory
# into our /app directory
ADD . /app
# We specify that we now wish to execute 
# any further commands inside our /app
# directory
WORKDIR /app

RUN go get
# we run go build to compile the binary
# executable of our Go program
RUN go build -o main .
# Our start command which kicks off
# our newly created binary executable
CMD ["/app/main"]
