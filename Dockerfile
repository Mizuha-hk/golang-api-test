# ベースイメージを指定
FROM golang:latest

# 作業ディレクトリを設定
WORKDIR /golang-api

# ビルドに必要なファイルをコピー
COPY go.mod .
COPY go.sum .

# モジュールをダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# ビルド
RUN go build -o main ./cmd/main.go

EXPOSE 8080

# アプリを実行
CMD ["./main"]