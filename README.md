# BlogGo - 静态博客生成器

一个用Go语言编写的简单静态博客生成器，支持Markdown格式的文章。

## 功能特性

- 📝 支持Markdown格式的文章
- 🎨 语法高亮显示
- 📱 响应式设计
- ⚡ 快速生成静态页面
- 🔧 可配置的博客设置
- 📄 自动生成文章列表和关于页面

## 快速开始

### 本地开发

1. 确保已安装Go 1.21或更高版本
2. 克隆仓库：
   ```bash
   git clone <your-repo-url>
   cd BlogGo
   ```
3. 运行博客生成器：
   ```bash
   go run main.go
   ```
4. 打开 `public/index.html` 查看生成的博客

### 添加新文章

1. 在 `content/` 目录下创建新的 `.md` 文件
2. 使用以下格式：

```markdown
---
title: "文章标题"
date: 2024-01-01T00:00:00Z
summary: "文章摘要"
---

这里是文章内容...
```

### 配置博客

编辑 `config.yaml` 文件来自定义博客设置：

```yaml
title: "我的博客"
description: "欢迎来到我的博客"
about:
  content: |
    <p>这里是关于页面的内容</p>
    <p>支持HTML格式</p>
```

## 部署到GitHub Pages

### 自动部署（推荐）

1. 将代码推送到GitHub仓库
2. 在仓库设置中启用GitHub Pages：
   - 进入仓库设置 → Pages
   - Source选择 "Deploy from a branch"
   - Branch选择 "gh-pages"
   - 点击Save

3. 每次推送代码到main分支时，GitHub Actions会自动构建并部署博客

### 手动部署

1. 本地构建博客：
   ```bash
   go run main.go
   ```
2. 将 `public/` 目录的内容推送到 `gh-pages` 分支

## 项目结构

```
BlogGo/
├── content/          # Markdown文章
├── pages/           # 页面文件（如about.md）
├── templates/       # HTML模板
├── static/          # 静态资源（CSS、图片等）
├── public/          # 生成的静态文件
├── config.yaml      # 博客配置
└── main.go          # 主程序
```

## 自定义样式

编辑 `static/css/style.css` 来自定义博客样式。

## 许可证

MIT License 