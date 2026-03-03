# 第一阶段：编译阶段
FROM golang:alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制依赖文件并下载
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 编译二进制文件
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 第二阶段：运行阶段
FROM alpine:latest
WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

# 1. 复制二进制文件
COPY --from=builder /app/main .

# 2. 核心修改：将整个前端目录复制进镜像
# 注意：这会把宿主机的 public 复制到容器的 /app/public
COPY --from=builder /app/public ./public

EXPOSE 8002
CMD ["./main"]