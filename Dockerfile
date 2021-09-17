FROM golang:1.17 as build-env

WORKDIR /go/src/app
COPY . /go/src/app

RUN go mod download
RUN go build -v -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/app /
CMD ["/app"]
