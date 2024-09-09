# ベースイメージを指定
FROM golang:1.22-alpine

# 作業ディレクトリを設定
WORKDIR /app

# Go モジュールと依存関係をコピー
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# アプリケーションを実行
CMD ["./main"]
