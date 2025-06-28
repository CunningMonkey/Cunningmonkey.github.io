#!/bin/bash

echo "🧪 本地测试博客..."

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
    
    # 检查是否有本地服务器
    if command -v python3 &> /dev/null; then
        echo "🌐 启动本地服务器..."
        echo "访问 http://localhost:8000 查看博客"
        echo "按 Ctrl+C 停止服务器"
        cd public && python3 -m http.server 8000
    elif command -v python &> /dev/null; then
        echo "🌐 启动本地服务器..."
        echo "访问 http://localhost:8000 查看博客"
        echo "按 Ctrl+C 停止服务器"
        cd public && python -m SimpleHTTPServer 8000
    else
        echo "💡 提示: 安装Python可以在本地启动服务器预览博客"
        echo "或者直接打开 public/index.html 文件"
    fi
else
    echo "❌ 博客构建失败！"
    exit 1
fi 