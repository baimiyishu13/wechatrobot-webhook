FROM golang:1.21.0 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o wechatrobot-webhook

# 第二阶段，使用轻量级的 alpine 镜像作为基础镜像
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/wechatrobot-webhook .
EXPOSE 8999
CMD ["./wechatrobot-webhook", "-addr=:3000", "-RobotKey=default_value"]