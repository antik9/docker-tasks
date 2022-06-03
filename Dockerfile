FROM golang:1.18 as build

WORKDIR /tmp
COPY main.go /tmp/main.go
RUN go build -o check-file-txt main.go

FROM busybox:1.35.0
COPY --from=build /tmp/check-file-txt /bin/

ENTRYPOINT ["/bin/sh"]
