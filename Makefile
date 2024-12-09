
SHELL=/bin/bash

# 番兵を置いておく（誤動作防止）
shell_check:
	@echo ${SHELL}

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