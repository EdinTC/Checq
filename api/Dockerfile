FROM golang:1.13-alpine AS builder

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
RUN CGO_ENABLED=0 go build -o checq-api .

# When building for development pilu/fresh will be used for code reloading via a shared volume.
# When building for production the code will be built as binary.
RUN if [ ${APP_ENV} = development ]; \
	then \
	go get github.com/pilu/fresh && \
	fresh; \
	fi

FROM scratch

COPY --from=builder /go/src/github.com/EdinTC/Checq/api/checq-api .

EXPOSE 1337

CMD ["./checq-api"]
