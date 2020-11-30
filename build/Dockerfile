FROM golang:1.15-alpine as builder
ENV GOPROXY=https://goproxy.io,direct
ENV GO111MODULE=on
ENV GOCACHE=/go/pkg/.cache/go-build
WORKDIR /work
ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app cmd/server/main.go

FROM centos
ENV LUBAN_CONFIG=/work/config/config.yaml
COPY --from=builder /work/app /work/app
WORKDIR /work
RUN mkdir config \
 && ./app config new -c /work/config/config.yaml
CMD ["./app"]