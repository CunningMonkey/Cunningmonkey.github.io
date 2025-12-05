# Giscus 评论系统配置指南

## 🎯 什么是 Giscus？

Giscus 是一个基于 GitHub Discussions 的评论系统，完全免费且无广告。所有评论数据都存储在你的 GitHub 仓库中。

## 📋 配置步骤

### 1. 启用 GitHub Discussions

1. 访问你的 GitHub 仓库：https://github.com/CunningMonkey/CunningMonkey.github.io
2. 点击 **Settings** 标签
3. 向下滚动找到 **Features** 部分
4. 勾选 **Discussions** 选项

### 2. 获取 Giscus 配置参数

1. 访问 Giscus 官网：https://giscus.app/zh-CN
2. 在网站上填写以下信息：
   - **仓库**：`CunningMonkey/CunningMonkey.github.io`
   - **页面 ↔️ discussion 映射关系**：选择 `pathname` (路径)
   - **Discussion 分类**：选择 `Announcements` 或创建新分类
   - **特性**：
     - ✅ 启用反应
     - 根据需要选择其他选项
   - **主题**：选择 `light`（浅色）或 `preferred_color_scheme`（自动）

3. 页面底部会生成配置代码，复制以下两个重要参数：
   - `data-repo-id="你的仓库ID"`
   - `data-category-id="你的分类ID"`

### 3. 更新模板文件

打开 `templates/base.html`，找到评论部分，替换以下两行：

```html
data-repo-id="YOUR_REPO_ID"        ← 替换为你的仓库ID
data-category-id="YOUR_CATEGORY_ID" ← 替换为你的分类ID
```

### 4. 重新构建博客

```bash
./build.sh
```

### 5. 部署到 GitHub Pages

```bash
./deploy.sh
```

或使用同步脚本：

```bash
./sync.sh
```

## 🎨 自定义主题

在 `templates/base.html` 中修改 `data-theme` 属性：

- `light` - 浅色主题
- `dark` - 深色主题
- `preferred_color_scheme` - 根据系统自动切换
- `transparent_dark` - 透明深色
- `dark_dimmed` - GitHub 暗淡主题

## 💡 其他评论系统选项

如果你想尝试其他评论系统，这里是一些选择：

### Utterances（基于 GitHub Issues）

优点：更轻量，配置简单
缺点：使用 Issues 而非 Discussions

```html
<script src="https://utteranc.es/client.js"
        repo="CunningMonkey/CunningMonkey.github.io"
        issue-term="pathname"
        theme="github-light"
        crossorigin="anonymous"
        async>
</script>
```

### Disqus（第三方服务）

优点：功能最全面，支持实时通知、管理面板
缺点：免费版有广告，加载较慢，需要注册账号

访问 https://disqus.com 注册并获取代码。

### Cusdis（轻量级开源方案）

优点：隐私友好，无广告，轻量级
缺点：需要自己部署服务器

访问 https://cusdis.com 了解更多。

## ❓ 常见问题

### Q: 评论区不显示？

A: 检查以下几点：
1. GitHub Discussions 是否已启用
2. `data-repo-id` 和 `data-category-id` 是否正确
3. 仓库是否为公开（public）
4. 浏览器是否屏蔽了第三方脚本

### Q: 如何管理评论？

A: 访问你的 GitHub 仓库 -> Discussions 标签，在那里可以查看、编辑、删除评论。

### Q: 评论者需要 GitHub 账号吗？

A: 是的，Giscus 和 Utterances 都需要评论者有 GitHub 账号并登录。这也是它们的优势——减少垃圾评论。

### Q: 如何限制谁可以评论？

A: 在 GitHub 仓库设置中配置 Discussions 的权限。

## 📚 更多资源

- [Giscus 官方文档](https://github.com/giscus/giscus)
- [GitHub Discussions 文档](https://docs.github.com/en/discussions)
- [Utterances 文档](https://utteranc.es/)

---

如有问题，欢迎在博客评论区留言！😊
