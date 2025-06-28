#!/bin/bash

echo "🔨 快速构建博客..."

# 检查Go是否安装
if ! command -v go &> /dev/null; then
    echo "❌ 错误: 未找到Go，请先安装Go 1.21或更高版本"
    exit 1
fi

# 清理旧的构建文件
echo "🧹 清理旧的构建文件..."
rm -rf public/

# 下载依赖
echo "📦 下载依赖..."
go mod download

# 构建博客
echo "🔨 构建博客..."
go run main.go

if [ $? -eq 0 ]; then
    echo "✅ 博客构建成功！"
    echo "📁 输出目录: public/"
    echo "💡 运行 ./test-local.sh 启动本地服务器预览"
else
    echo "❌ 博客构建失败！"
    exit 1
fi 