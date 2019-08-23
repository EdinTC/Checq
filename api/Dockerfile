FROM golang:1.12-alpine

# To fix go get and build with cgo.
RUN apk add --no-cache --virtual .build-deps \
    bash \
    gcc \
    git \
    musl-dev

# This runs as a development build by default.
# To build for production run: docker build ./api --build-args app_env=production
ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/github.com/EdinTC/Checq/api
WORKDIR /go/src/github.com/EdinTC/Checq/api

# If you add any vendor services add them here (RUN go get ./vendor/...)


RUN go get ./
RUN go build

# When building for development pilu/fresh will be used for code reloading via a shared volume.
# When building for production the code will be built as binary.
CMD if [ ${APP_ENV} = production ]; \
	then \
	api; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi

EXPOSE 1337