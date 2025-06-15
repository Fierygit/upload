# 文件路径：/root/upload/Makefile

# 定义变量
APP_NAME := upload_app
GO_CMD := /usr/local/go/bin/go  # 替换为你的 go 绝对路径
BUILD_DIR := ./      # 构建目标目录

.PHONY: build clean

build:
	@echo "正在构建应用..."
	cd $(BUILD_DIR) && $(GO_CMD) build -o $(APP_NAME)
	@echo "构建完成！二进制文件: $(BUILD_DIR)/$(APP_NAME)"

clean:
	@echo "清理构建文件..."
	rm -f $(BUILD_DIR)/$(APP_NAME)
	@echo "清理完成"