#!/bin/bash

echo "🚀 手动部署博客到GitHub Pages..."

# 获取仓库信息
REPO_NAME=$(basename -s .git $(git config --get remote.origin.url))
USER_NAME=$(git config --get remote.origin.url | sed -n 's/.*github\.com[:/]\([^/]*\)\/.*/\1/p')

echo "📦 仓库: $USER_NAME/$REPO_NAME"

# 检查是否有未提交的更改
if ! git diff-index --quiet HEAD --; then
    echo "⚠️  发现未提交的更改，正在提交..."
    git add .
    git commit -m "Update blog content"
    git push origin main
fi

# 重新构建博客
echo "🔨 重新构建博客..."
go run main.go

if [ $? -eq 0 ]; then
    echo "✅ 博客构建成功"
    
    # 删除现有的gh-pages分支
    echo "🗑️  删除现有的gh-pages分支..."
    git push origin --delete gh-pages 2>/dev/null
    
    # 等待删除完成
    sleep 3
    
    # 创建新的gh-pages分支
    echo "🌐 创建新的gh-pages分支..."
    
    # 切换到public目录
    cd public
    
    # 初始化新的git仓库
    git init
    git add .
    git commit -m "Manual deployment - $(date)"
    
    # 添加远程仓库
    git remote add origin git@github.com:$USER_NAME/$REPO_NAME.git
    
    # 强制推送到gh-pages分支
    git push -f origin HEAD:gh-pages
    
    # 返回原目录
    cd ..
    
    # 清理临时git仓库
    rm -rf public/.git
    
    echo "✅ 手动部署完成"
else
    echo "❌ 博客构建失败"
    exit 1
fi

echo ""
echo "🎉 手动部署成功！"
echo ""
echo "📋 接下来："
echo "1. 访问 https://github.com/$USER_NAME/$REPO_NAME/settings/pages"
echo "2. 确保 'Source' 设置为 'Deploy from a branch'"
echo "3. 确保 'Branch' 设置为 'gh-pages'"
echo "4. 确保 'Custom domain' 字段为空"
echo "5. 点击 'Save'"
echo ""
echo "🌐 部署完成后访问："
echo "   https://$USER_NAME.github.io/$REPO_NAME/"
echo ""
echo "⏳ 可能需要等待5-10分钟才能完全生效..." 