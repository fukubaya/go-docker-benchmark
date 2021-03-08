FROM golang:1.13-alpine3.12 as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go build -o /go/bin/app

FROM alpine:3.12
COPY --from=build /go/bin/app /
CMD ["/app"]

