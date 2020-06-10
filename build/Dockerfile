FROM golang:1.14-alpine as builder
ENV GOPROXY=https://goproxy.io
ENV GO111MODULE=on
ENV GOCACHE=/go/pkg/.cache/go-build
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app cmd/main.go

FROM alpine
COPY --from=builder app /work/app
WORKDIR /work
CMD ["./app"]