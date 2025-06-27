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
}

func main() {
	config := Config{
		Title:       "my simple blog",
		Description: "welcome to my simple blog",
	}

	os.RemoveAll("public")
	os.MkdirAll("public", 0755)

	posts := processPosts(config)
	generateListPage(config, posts)
	copyStaticFiles()
	fmt.Println("blog generated! output directory: public")
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
