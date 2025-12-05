# 🌗 暗色模式功能说明

## ✨ 新增功能

你的博客现在支持**亮色/暗色模式切换**！

### 主要特性

1. **右上角主题切换按钮** 🌓
   - 一键切换亮色/暗色模式
   - 平滑动画过渡效果
   - 图标随主题变化（太阳/月亮）

2. **自动保存偏好设置** 💾
   - 使用 localStorage 保存用户选择
   - 下次访问自动应用上次的主题

3. **评论区自动适配** 💬
   - Giscus 评论系统跟随主题切换
   - 亮色模式：白色背景
   - 暗色模式：深色背景

4. **全站响应式设计** 📱
   - 在移动设备上完美显示
   - 主题切换按钮自适应屏幕大小

## 🎨 主题颜色方案

### 亮色模式（Light）
- 背景：浅灰色 (#fafafa)
- 卡片：白色 (#ffffff)
- 文字：深色 (#2c3e50)
- 链接：蓝色 (#3182ce)

### 暗色模式（Dark）
- 背景：深灰色 (#0d1117)
- 卡片：深色 (#161b22)
- 文字：浅色 (#e6edf3)
- 链接：亮蓝色 (#58a6ff)

## 🔧 技术实现

### CSS 变量系统
使用 CSS 自定义属性实现主题切换：

```css
:root {
    --bg-primary: #fafafa;
    --text-primary: #2c3e50;
    /* 更多变量... */
}

[data-theme="dark"] {
    --bg-primary: #0d1117;
    --text-primary: #e6edf3;
    /* 暗色变量... */
}
```

### JavaScript 主题管理
- 监听按钮点击事件
- 切换 `data-theme` 属性
- 保存到 localStorage
- 重新加载 Giscus 评论框以应用新主题

## 📝 如何使用

### 用户端
1. 访问博客任意页面
2. 点击右上角的主题切换按钮（☀️/🌙）
3. 主题立即切换，并自动保存

### 开发者自定义

#### 修改主题颜色
编辑 `static/css/style.css` 中的 CSS 变量：

```css
:root {
    --bg-primary: #你的颜色;
    --text-primary: #你的颜色;
}

[data-theme="dark"] {
    --bg-primary: #你的暗色;
    --text-primary: #你的暗色;
}
```

#### 修改按钮位置
在 `static/css/style.css` 中调整 `.theme-toggle` 样式：

```css
.theme-toggle {
    position: fixed;  /* 固定位置 */
    top: 20px;
    right: 20px;
    /* 其他样式... */
}
```

#### 修改默认主题
编辑 `templates/base.html` 中的 JavaScript：

```javascript
const savedTheme = localStorage.getItem('theme') || 'dark';  // 改为 'dark'
```

## 🎯 兼容性

- ✅ Chrome/Edge 88+
- ✅ Firefox 85+
- ✅ Safari 14+
- ✅ 移动浏览器

## 📋 文件修改清单

### 更新的文件

1. **templates/base.html**
   - 添加主题切换按钮
   - 添加主题切换 JavaScript
   - 更新 Giscus 集成方式

2. **static/css/style.css**
   - 添加 CSS 变量定义
   - 将所有硬编码颜色改为变量
   - 添加暗色模式样式
   - 添加主题切换按钮样式
   - 优化过渡动画

## 🚀 部署说明

### 本地测试
```bash
./build.sh
./test-local.sh
```

访问 http://localhost:8000 查看效果

### 部署到 GitHub Pages
```bash
./deploy.sh
```

或使用同步方式：
```bash
./sync.sh
```

## 💡 进阶定制

### 1. 添加更多主题
可以扩展为支持多个主题（如：蓝色、绿色、紫色）：

```css
[data-theme="blue"] {
    --bg-primary: #0c4a6e;
    --text-primary: #e0f2fe;
    /* ... */
}
```

### 2. 跟随系统主题
修改 JavaScript 以检测系统主题：

```javascript
const systemTheme = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
const savedTheme = localStorage.getItem('theme') || systemTheme;
```

### 3. 添加主题选择器
创建下拉菜单以选择多个主题选项。

## 📸 效果预览

### 亮色模式
- 清爽明亮的白色背景
- 适合白天阅读
- 减少眼部疲劳

### 暗色模式
- 优雅的深色背景
- 适合夜间阅读
- 保护视力，省电（OLED 屏幕）

## 🔍 故障排除

### 问题1：主题切换后评论区没有变化
**解决**：刷新页面，Giscus 会重新加载

### 问题2：移动端按钮显示异常
**解决**：检查 CSS 媒体查询是否正确加载

### 问题3：主题不保存
**解决**：检查浏览器是否禁用了 localStorage

## 🎉 总结

现在你的博客拥有了：
- ✅ 现代化的暗色模式
- ✅ 平滑的主题切换动画
- ✅ 评论系统完美适配
- ✅ 用户偏好自动保存
- ✅ 移动端完美支持

享受你的全新博客体验吧！🚀

---

如有问题或建议，欢迎在评论区留言！
