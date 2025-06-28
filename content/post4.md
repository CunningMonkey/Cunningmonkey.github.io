---
title: "hhhhhh"
date: 2024-01-15T10:00:00Z
summary: "学习如何向你的静态博客添加新文章并自动部署"
---

# 如何添加新文章到博客

## 步骤1：创建Markdown文件

在 `content/` 目录下创建一个新的 `.md` 文件，例如 `post2.md`。

## 步骤2：添加Front Matter

每个文章都需要在开头添加YAML格式的元数据：

```yaml
---
title: "文章标题"
date: 2024-01-15T10:00:00Z
summary: "文章摘要"
---
```


## 步骤3：编写内容

在Front Matter后面添加你的文章内容，支持Markdown格式：

- **粗体文本**
- *斜体文本*
- `代码片段`
- [链接](https://example.com)

## 步骤4：部署

推送代码到GitHub，博客会自动重新构建和部署！

```bash
git add .
git commit -m "添加新文章：如何添加新文章到博客"
git push origin main
```

就是这么简单！ 