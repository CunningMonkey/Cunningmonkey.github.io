#!/bin/bash

echo "📚 BlogGo 使用指南"
echo "=================="
echo ""

echo "🔧 可用的脚本命令："
echo ""

echo "📝 内容管理："
echo "  ./sync.sh      - 同步新内容到GitHub（交互式）"
echo "  ./deploy.sh    - 直接部署博客到GitHub Pages"
echo ""

echo "🧪 本地测试："
echo "  ./build.sh     - 快速构建博客（不启动服务器）"
echo "  ./test-local.sh - 构建并启动本地服务器预览"
echo ""

echo "📋 使用流程："
echo ""

echo "1. 添加新文章："
echo "   - 在 content/ 目录创建 .md 文件"
echo "   - 添加Front Matter（标题、日期、摘要）"
echo "   - 编写内容"
echo ""

echo "2. 本地测试："
echo "   ./build.sh        # 构建博客"
echo "   ./test-local.sh   # 本地预览"
echo ""

echo "3. 部署到GitHub："
echo "   ./sync.sh         # 交互式同步"
echo "   或者"
echo "   ./deploy.sh       # 直接部署"
echo ""

echo "🌐 访问地址："
echo "   本地预览: http://localhost:8000"
echo "   在线博客: https://cunningmonkey.github.io/"
echo ""

echo "📁 重要目录："
echo "   content/     - Markdown文章"
echo "   pages/       - 页面文件（如about.md）"
echo "   templates/   - HTML模板"
echo "   static/      - 静态资源（CSS、图片）"
echo "   public/      - 生成的静态文件"
echo ""

echo "💡 提示："
echo "   - 每次修改后建议先本地测试"
echo "   - 部署后可能需要等待几分钟生效"
echo "   - 使用 ./sync.sh 可以交互式地输入提交信息" 