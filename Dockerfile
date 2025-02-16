## Base Image Go
FROM golang:1.23 AS builder

WORKDIR /app

# Goのモジュールとソースコードをコピー
COPY go.mod go.sum ./

# サーバのソースコードをコピー
COPY . .

WORKDIR /app/client

# Goでビルド
#RUN GOOS=linux GOARCH=amd64 go build -o server .
RUN go build -o client .

# 軽量な実行環境
FROM golang:1.23 AS final

## 必要な依存関係をインストール（gRPCの依存）
#RUN apk --no-cache add ca-certificates

# ビルドしたバイナリをコピー
COPY --from=builder /app/client /client

RUN chmod 777 /client/client

EXPOSE 50051
# コンテナが起動した時に実行するコマンド
CMD ["/client/client"]