#!/bin/bash

echo "🌐 设置直接访问 cunningmonkey.github.io..."

# 获取当前仓库信息
REPO_NAME=$(basename -s .git $(git config --get remote.origin.url))
USER_NAME=$(git config --get remote.origin.url | sed -n 's/.*github\.com[:/]\([^/]*\)\/.*/\1/p')

echo "📦 当前仓库: $USER_NAME/$REPO_NAME"

if [ "$REPO_NAME" = "$USER_NAME.github.io" ]; then
    echo "✅ 仓库名称已经是正确的格式！"
    echo "🌐 你的博客应该可以通过 https://$USER_NAME.github.io/ 访问"
    exit 0
fi

echo ""
echo "🚨 要实现直接访问 $USER_NAME.github.io，需要重命名仓库。"
echo ""
echo "📋 重命名步骤："
echo ""
echo "🔧 步骤1：在GitHub上重命名仓库"
echo "1. 访问: https://github.com/$USER_NAME/$REPO_NAME/settings"
echo "2. 在 'Repository name' 部分"
echo "3. 将仓库名从 '$REPO_NAME' 改为 '$USER_NAME.github.io'"
echo "4. 点击 'Rename' 按钮"
echo ""
echo "🔧 步骤2：更新本地仓库"
echo "重命名完成后，运行以下命令："
echo "git remote set-url origin git@github.com:$USER_NAME/$USER_NAME.github.io.git"
echo ""
echo "🔧 步骤3：重新部署"
echo "运行以下命令重新部署："
echo "./deploy.sh"
echo ""
echo "🌐 重命名完成后，你的博客将可以通过以下地址访问："
echo "   https://$USER_NAME.github.io/"
echo ""
echo "⚠️  注意："
echo "- 重命名后，旧的URL将自动重定向到新URL"
echo "- 重命名可能需要几分钟才能生效"
echo "- 建议在重命名前备份重要数据" 