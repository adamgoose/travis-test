FROM golang:1.12-alpine3.11 as build
RUN apk add --update git
WORKDIR /work
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin .

FROM alpine:3.11
RUN apk add --update --no-cache ca-certificates
COPY --from=build /work/bin /usr/local/bin/travis-test
CMD ["/usr/local/bin/travis-test"]
