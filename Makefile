
SHELL=/bin/bash

# 番兵を置いておく（誤動作防止）
shell_check:
	@echo ${SHELL}

# ビルド
build:
	@docker build --no-cache --tag grpc-cli-chat .

# 起動
run:
	@docker run -it -v C:\Users\yuuzi\go\grpc-cli-chat://usr/src/app --rm --name app grpc-cli-chat /bin/bash

# 自動生成ファイルの作成
proto_gen:
	@protoc --go_out=./chatService/ --go-grpc_out=./chatService/ ./chatService/*.proto

git_config_list:
	@git config --global --list

run_client:
	@go run ./client/.

run_server:
	@go run ./server/.

lint:
	@go vet ./...