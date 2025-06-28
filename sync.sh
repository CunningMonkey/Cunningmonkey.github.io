#!/bin/bash

echo "🔄 同步新内容到GitHub Pages..."

# 检查是否有未提交的更改
if git diff-index --quiet HEAD --; then
    echo "ℹ️  没有发现新的更改"
    echo "💡 提示：如果你添加了新文章，请先保存文件"
    exit 0
fi

# 显示更改的文件
echo "📝 发现以下更改："
git status --porcelain

# 询问用户是否继续
echo ""
read -p "是否继续同步到GitHub？(y/N): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "❌ 同步已取消"
    exit 0
fi

# 添加所有更改
echo "📦 添加更改到Git..."
git add .

# 获取提交信息
echo ""
echo "💬 请输入提交信息（例如：添加新文章、更新样式等）："
read -p "提交信息: " commit_message

if [ -z "$commit_message" ]; then
    commit_message="更新博客内容"
fi

# 提交更改
echo "💾 提交更改..."
git commit -m "$commit_message"

# 推送到GitHub
echo "📤 推送到GitHub..."
git push origin main

if [ $? -eq 0 ]; then
    echo ""
    echo "✅ 同步成功！"
    echo ""
    echo "📋 接下来："
    echo "1. GitHub Actions会自动构建和部署你的博客"
    echo "2. 访问 https://github.com/CunningMonkey/blog_source/actions 查看部署进度"
    echo "3. 部署完成后访问 https://cunningmonkey.github.io/blog_source/"
    echo ""
    echo "⏳ 部署通常需要2-5分钟完成"
else
    echo "❌ 推送失败，请检查网络连接或GitHub权限"
    exit 1
fi 