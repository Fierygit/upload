#!/bin/bash
# 文件路径：/root/upload/start.sh

# 定义绝对路径
APP_NAME="upload_app"
APP_PATH="./${APP_NAME}"  # 使用绝对路径
LOG_FILE="./app.log"
PORT=7070

# 1. 检查并终止占用端口的进程
echo "检查端口 ${PORT} 占用情况..."
PID=$(/usr/bin/lsof -t -i:${PORT} || echo "")

if [ ! -z "$PID" ]; then
    echo "发现占用端口 ${PORT} 的进程 (PID: ${PID})，正在终止..."
    /bin/kill -9 $PID
    sleep 1
    echo "进程已终止。"
fi

# 2. 启动应用
echo "启动应用..."
nohup ${APP_PATH} >> ${LOG_FILE} 2>&1 &

# 3. 验证启动
sleep 2
NEW_PID=$(/usr/bin/pgrep -f ${APP_NAME})
if [ ! -z "$NEW_PID" ]; then
    echo "应用启动成功！PID: ${NEW_PID}"
    echo "日志输出到: ${LOG_FILE}"
else
    echo "启动失败，请检查日志！"
    exit 1
fi