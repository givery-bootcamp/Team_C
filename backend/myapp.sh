#!/bin/sh
# wrapper.sh

# Goプロセスを開始
./myapp &
app_pid=$!

# シグナルハンドリング
trap 'kill -TERM $app_pid; wait $app_pid' TERM INT

# Goプロセスの終了を待つ
wait $app_pid

# 追加の待機時間を設定（必要に応じて調整）
sleep 5

# 終了
exit 0