FROM golang:1.23.4-bookworm

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
# COPY ./../.ssh/id_rsa ~/.ssh/
# RUN go build -v -o /usr/local/bin ./...

RUN apt update && apt install -y \
    curl \
    wget \
    git \
    unzip \
    build-essential \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# ref. https://zenn.dev/mokomoka/articles/1cfdddec2d0887
RUN PROTOC_VERSION=$(curl -s "https://api.github.com/repos/protocolbuffers/protobuf/releases/latest" | grep -Po '"tag_name": "v\K[0-9.]+') \
    && curl -Lo protoc.zip "https://github.com/protocolbuffers/protobuf/releases/latest/download/protoc-${PROTOC_VERSION}-linux-x86_64.zip" \
    && unzip -q protoc.zip bin/protoc 'include/*' -d /usr/local \
    && chmod a+x /usr/local/bin/protoc \
    && rm -rf protoc.zip

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# ユーザーを変更
RUN groupadd -r dev && useradd -r -g dev dev
USER dev

# アプリを実行したいとき
# CMD ["app"]