FROM golang:1.15-alpine as builder
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
ENV GOCACHE=/go/pkg/.cache/go-build
WORKDIR /work
ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app cmd/main.go

FROM alpine
ENV LUBAN_CONFIG=/work/config.yaml
COPY --from=builder /work/app /work/app
COPY --from=builder /work/public /work/public
WORKDIR /work
RUN ./app config new
CMD ["./app", "server"]