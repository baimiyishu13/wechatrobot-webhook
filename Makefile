# 二进制
bin/wechat-webhook:
	GOOS=linux GOARCH=amd64 go build -ldflags "-w" -o bin/wechat-webhook

# docker

