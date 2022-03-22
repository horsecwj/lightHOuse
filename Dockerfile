FROM golang:1.17.2-alpine AS builder

COPY ./ /app
WORKDIR /app
ENV GOOS linux
ENV CGO_ENABLED 0
ENV GOARCH amd64
ENV GO111MODULE  on
ENV GOPROXY https://goproxy.cn

RUN go build -a -installsuffix cgo -ldflags '-s -w' -o light-house cmd/server/main.go

FROM --platform=linux/amd64 alpine:3.15
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /opt/zoneinfo.zip
COPY --from=builder /app/light-house /
ENV ZONEINFO /opt/zoneinfo.zip
ADD static /static
RUN ln -s /light-house /usr/bin

CMD ["light-house"]