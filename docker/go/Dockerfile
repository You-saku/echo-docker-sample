FROM golang:1.17.6-alpine

# テスト実行するためにはgccが必要です
RUN apk add --no-cache gcc libc-dev

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY . /go/src/app

# ワーキングディレクトリの設定
WORKDIR /go/src/app/api

RUN go build
