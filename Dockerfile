ARG GO_VERSION=1
FROM golang:${GO_VERSION}-buster as builder

WORKDIR /usr/src/app
COPY go.mod ./
COPY . .
RUN go build -v -o /run-app .


FROM debian:buster

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
