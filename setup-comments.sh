#!/bin/bash

echo "🔧 Giscus 评论系统配置助手"
echo "================================"
echo ""
echo "📖 在开始之前，请确保："
echo "1. 你的 GitHub 仓库已启用 Discussions 功能"
echo "2. 仓库是公开的（public）"
echo ""
echo "🌐 请访问以下网址获取配置参数："
echo "   https://giscus.app/zh-CN"
echo ""
echo "================================"
echo ""

# 提示用户输入配置参数
read -p "请输入你的 data-repo-id: " repo_id
read -p "请输入你的 data-category-id: " category_id

if [ -z "$repo_id" ] || [ -z "$category_id" ]; then
    echo ""
    echo "❌ 错误: repo-id 和 category-id 不能为空"
    echo ""
    echo "💡 获取方法："
    echo "1. 访问 https://giscus.app/zh-CN"
    echo "2. 输入仓库名：CunningMonkey/CunningMonkey.github.io"
    echo "3. 选择映射方式和分类"
    echo "4. 在页面底部的代码中找到这两个ID"
    exit 1
fi

# 备份原文件
cp templates/base.html templates/base.html.backup
echo "✅ 已备份原文件到 templates/base.html.backup"

# 替换配置
sed -i.tmp "s/data-repo-id=\"YOUR_REPO_ID\"/data-repo-id=\"$repo_id\"/g" templates/base.html
sed -i.tmp "s/data-category-id=\"YOUR_CATEGORY_ID\"/data-category-id=\"$category_id\"/g" templates/base.html
rm templates/base.html.tmp 2>/dev/null

echo ""
echo "✅ 配置完成！"
echo ""
echo "📋 配置信息："
echo "   Repo ID: $repo_id"
echo "   Category ID: $category_id"
echo ""
echo "🔨 接下来的步骤："
echo "1. 运行 ./build.sh 重新构建博客"
echo "2. 运行 ./test-local.sh 本地预览（可选）"
echo "3. 运行 ./deploy.sh 或 ./sync.sh 部署到 GitHub Pages"
echo ""
echo "💡 如需恢复原配置，运行："
echo "   cp templates/base.html.backup templates/base.html"
