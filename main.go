// Copyright (C) 2026 my-app-s
// Licensed under the GNU AGPLv3

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// Теги `json:"..."` говорят Go, как сопоставлять ключи из файла с полями структуры
type Tools struct {
	Name string `json:"name"`
}

type Urls struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PageData struct {
	NameRepository string        `json:"name_repository"`
	NameAuthor     string        `json:"name_author"`
	Description    template.HTML // Заполняется динамически из README.md
	URLAvatar      string        `json:"url_avatar"`
	URLRepository  string        `json:"url_repository"`
	Stack          []Tools       `json:"stack"`
	Links          []Urls        `json:"links"`
	CopyrightYear  int           `json:"copyright_year"`
}

// Функция для скачивания README.md с GitHub по сырой ссылке (Raw)
func fetchReadme(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ошибка загрузки README: статус %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Функция для конвертации Markdown-текста в HTML
func convertMarkdownToHTML(mdContent string) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(mdContent))

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	output := markdown.Render(doc, renderer)
	return output
}

func main() {
	// Disable because Github Actions itself adds a timestamp
	log.SetFlags(0)

	configPath := flag.String("config", "config.json", "Path to config file")
	readmePath := flag.String("readme", "README.md", "Path to README file")
	flag.Parse()

	// 1. Reading the configuration file config.json
	fmt.Println("Открываю config.json:", *configPath)
	configFile, err := os.Open(*configPath)
	if err != nil {
		log.Fatalf("::error:: failed to open config.json: %v", err)
    }
	defer configFile.Close()

	var data PageData
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatalf("::error:: error parsing config.json: %v", err)
    }

	// We read the local README.md, which lies next to config.json
	fmt.Println("Reading README.md:", *readmePath)
	mdBytes, err := os.ReadFile(*readmePath)
	var mdText string
	if err != nil {
		fmt.Println("Warning: Could not find local README.md:", err)
		mdText = "Description temporarily unavailable."
	} else {
		mdText = string(mdBytes)
	}

	htmlDescription := convertMarkdownToHTML(mdText)
	data.Description = template.HTML(htmlDescription) // Adding dynamic HTML to the structure

	// 3. Reading the template
	fmt.Println("I'm reading the template templates/index.html")
	tmplHTML, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("::error:: template templates/index.html not found or invalid: %v", err)
    }

	// 4. Create a dist folder
	err = os.MkdirAll("dist", 0755)
    if err != nil {
        log.Fatalf("::error:: failed to create dist directory: %v", err)
    }

	// 5. Creating the final file
	fmt.Println("I create dist/index.html")
	fHTML, err := os.Create("dist/index.html")
	if err != nil {
		log.Fatalf("::error:: failed to create dist/index.html: %v", err)
    }
	defer fHTML.Close()

	// 6. Rendering data to a file
	err = tmplHTML.Execute(fHTML, data)
	if err != nil {
		log.Fatalf("::error:: failed to execute template: %v", err)
    }

	// 7. Copy input.css
	fmt.Println("Создаю dist/input.css")
	tmplCSS, err := os.Open("templates/input.css")
	if err != nil {
		log.Fatalf("::error:: template input.css not found: %v", err)
	}
	defer tmplCSS.Close()

	fCSS, err := os.Create("dist/input.css")
	if err != nil {
		log.Fatalf("::error:: file input.css not created: %v", err)
	}

	// 1. Copying data
	_, err = io.Copy(fCSS, tmplCSS)
	if err != nil {
		fCSS.Close() // Close before falling so as not to leave the handle open
		log.Fatalf("::error:: failed to copy input.css: %v", err)
	}

	// 2. Close and check the error ONLY AFTER copying
	if err := fCSS.Close(); err != nil {
		log.Fatalf("::error:: failed to close input.css: %v", err)
	}

	println("Ready! The site is generated in the /dist folder with data from config.json")
}
