FROM golang:1.9 as test
WORKDIR /go/src/app
COPY main.go .
COPY main_test.go .
RUN go test

###############################################################################
FROM golang:1.9 as builder
COPY main.go .
RUN CGO_ENABLED=0 go build -a -ldflags '-s' -o /go-redirect .

###############################################################################
FROM scratch
LABEL maintainer="Scott Miller <scott.miller171@gmail.com>"

COPY --from=builder /go-redirect /go-redirect
EXPOSE 80
ENTRYPOINT [ "/go-redirect" ]
