FROM golang:1.18-alpine AS builder

COPY . /github.com/ksusonic/owl-morning-bot/
WORKDIR /github.com/ksusonic/owl-morning-bot/

RUN go mod download
RUN go build -o ./bin/bot cmd/main.go

FROM alpine:latest
RUN apk update && apk add tzdata
ENV TZ=Europe/Moscow

WORKDIR /root/

COPY --from=0 /github.com/ksusonic/owl-morning-bot/bin/bot .
COPY --from=0 /github.com/ksusonic/owl-morning-bot/config config/

CMD ["./bot", "-c", "config/prod.yaml"]
