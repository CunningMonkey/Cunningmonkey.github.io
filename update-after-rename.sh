#!/bin/bash

echo "🔄 更新仓库配置..."

# 获取用户名
USER_NAME=$(git config --get remote.origin.url | sed -n 's/.*github\.com[:/]\([^/]*\)\/.*/\1/p')

echo "📦 用户: $USER_NAME"

# 更新远程仓库URL
echo "🔧 更新远程仓库URL..."
git remote set-url origin git@github.com:$USER_NAME/$USER_NAME.github.io.git

# 验证更新
echo "✅ 验证远程仓库URL..."
git remote -v

# 推送当前代码到新仓库
echo "📤 推送代码到新仓库..."
git push -u origin main

# 重新部署
echo "🚀 重新部署博客..."
./deploy.sh

echo ""
echo "🎉 更新完成！"
echo "🌐 你的博客现在可以通过以下地址访问："
echo "   https://$USER_NAME.github.io/"
echo ""
echo "⏳ 可能需要等待几分钟才能完全生效..." 