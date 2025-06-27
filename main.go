package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v2"
)

type FrontMatter struct {
	Title   string    `yaml:"title"`
	Date    time.Time `yaml:"date"`
	Summary string    `yaml:"summary"`
}

type Post struct {
	FrontMatter
	Slug    string
	Content template.HTML
}

type Config struct {
	Title       string
	Description string
	About       AboutInfo
}

type AboutInfo struct {
	Content template.HTML
}

func main() {
	config := loadConfig()

	os.RemoveAll("public")
	os.MkdirAll("public", 0755)

	posts := processPosts(config)
	generateListPage(config, posts)
	generateAboutPage(config)
	copyStaticFiles()
	fmt.Println("blog generated! output directory: public")
}

func loadConfig() Config {
	// 默认配置
	defaultConfig := Config{
		Title:       "my simple blog",
		Description: "welcome to my simple blog",
		About: AboutInfo{
			Content: template.HTML(`
				<p>你好！欢迎来到我的博客。</p>
				<p>我是一名热爱技术的开发者，喜欢分享我的学习和经验。</p>
				<p>这个博客是用Go语言构建的静态博客生成器，支持Markdown格式的文章。</p>
				<p>如果你有任何问题或建议，欢迎与我交流！</p>
			`),
		},
	}

	// 尝试读取配置文件
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("配置文件 config.yaml 不存在，使用默认配置\n")
		return defaultConfig
	}

	var configData struct {
		Title       string `yaml:"title"`
		Description string `yaml:"description"`
		About       struct {
			Content string `yaml:"content"`
		} `yaml:"about"`
	}

	if err := yaml.Unmarshal(data, &configData); err != nil {
		fmt.Printf("解析配置文件失败: %v，使用默认配置\n", err)
		return defaultConfig
	}

	config := Config{
		Title:       configData.Title,
		Description: configData.Description,
		About: AboutInfo{
			Content: template.HTML(configData.About.Content),
		},
	}

	// 如果配置文件中的字段为空，使用默认值
	if config.Title == "" {
		config.Title = defaultConfig.Title
	}
	if config.Description == "" {
		config.Description = defaultConfig.Description
	}
	if config.About.Content == "" {
		config.About.Content = defaultConfig.About.Content
	}

	return config
}

func processPosts(config Config) []Post {
	var posts []Post

	err := filepath.Walk("content", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		post, err := parsePost(data, path)
		if err != nil {
			log.Printf("parse failed %s: %v", path, err)
			return nil
		}

		generatePostPage(config, post)
		posts = append(posts, post)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return posts
}

func parsePost(data []byte, path string) (Post, error) {
	parts := bytes.SplitN(data, []byte("---"), 3)
	if len(parts) < 3 {
		return Post{}, fmt.Errorf("无效的 Front Matter")
	}

	var fm FrontMatter
	if err := yaml.Unmarshal(parts[1], &fm); err != nil {
		return Post{}, err
	}

	// 使用文件名作为slug，去掉.md扩展名
	baseName := filepath.Base(path)
	slug := strings.TrimSuffix(baseName, ".md")

	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithExtensions(
			extension.GFM,
			extension.DefinitionList,
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
			),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	if err := md.Convert(parts[2], &buf); err != nil {
		return Post{}, err
	}

	return Post{
		FrontMatter: fm,
		Slug:        slug,
		Content:     template.HTML(buf.String()),
	}, nil
}

func createTemplate(files ...string) (*template.Template, error) {
	funcMap := template.FuncMap{
		"now": time.Now,
	}

	tmpl := template.New(filepath.Base(files[0])).Funcs(funcMap)
	return tmpl.ParseFiles(files...)
}

func generatePostPage(config Config, post Post) {
	tmpl, err := createTemplate("templates/base.html")
	if err != nil {
		log.Fatal(err)
	}

	outputPath := filepath.Join("public", post.Slug+".html")
	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, struct {
		Config
		Post
	}{
		Config: config,
		Post:   post,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func generateListPage(config Config, posts []Post) {
	tmpl, err := createTemplate("templates/list.html")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("public/index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, struct {
		Config
		Posts []Post
	}{
		Config: config,
		Posts:  posts,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func generateAboutPage(config Config) {
	// 首先尝试从Markdown文件读取关于页面内容
	aboutContent, err := processAboutPage()
	if err != nil {
		log.Printf("处理关于页面失败: %v，使用默认内容", err)
		// 使用默认内容
		aboutContent = config.About.Content
	}

	tmpl, err := createTemplate("templates/about.html")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("public/about.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, struct {
		Config
		About AboutInfo
	}{
		Config: config,
		About:  AboutInfo{Content: aboutContent},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func processAboutPage() (template.HTML, error) {
	aboutPath := filepath.Join("pages", "about.md")

	data, err := os.ReadFile(aboutPath)
	if err != nil {
		return "", fmt.Errorf("无法读取关于页面文件: %v", err)
	}

	// 解析Front Matter和内容
	parts := bytes.SplitN(data, []byte("---"), 3)
	if len(parts) < 3 {
		return "", fmt.Errorf("关于页面文件格式无效")
	}

	// 转换Markdown内容为HTML
	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithExtensions(
			extension.GFM,
			extension.DefinitionList,
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
			),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	if err := md.Convert(parts[2], &buf); err != nil {
		return "", fmt.Errorf("转换Markdown失败: %v", err)
	}

	return template.HTML(buf.String()), nil
}

func copyStaticFiles() {
	staticDir := "static"
	publicStaticDir := filepath.Join("public", "static")

	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		fmt.Printf("静态文件目录 %s 不存在，跳过复制\n", staticDir)
		return
	}

	fmt.Printf("开始复制静态文件从 %s 到 %s\n", staticDir, publicStaticDir)

	err := os.MkdirAll(publicStaticDir, 0755)
	if err != nil {
		log.Printf("创建目录失败 %s: %v", publicStaticDir, err)
		return
	}

	err = filepath.Walk(staticDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("访问路径失败 %s: %v", path, err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(staticDir, path)
		if err != nil {
			log.Printf("计算相对路径失败 %s: %v", path, err)
			return err
		}

		destPath := filepath.Join(publicStaticDir, relPath)

		// 确保目标目录存在
		destDir := filepath.Dir(destPath)
		if err := os.MkdirAll(destDir, 0755); err != nil {
			log.Printf("创建目标目录失败 %s: %v", destDir, err)
			return err
		}

		input, err := os.ReadFile(path)
		if err != nil {
			log.Printf("读取文件失败 %s: %v", path, err)
			return err
		}

		err = os.WriteFile(destPath, input, 0644)
		if err != nil {
			log.Printf("写入文件失败 %s: %v", destPath, err)
			return err
		}

		fmt.Printf("复制文件: %s -> %s\n", path, destPath)
		return nil
	})

	if err != nil {
		log.Printf("复制静态文件时出错: %v", err)
	} else {
		fmt.Printf("静态文件复制完成\n")
	}
}
