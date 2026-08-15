// Copyright (C) 2026 my-app-s (M.R.E)

package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// Описываем структуру данных для сайта
type Tools struct {
	Name string
}

type Urls struct {
	Name string
	URL  string
}

type PageData struct {
	NameRepository string
	NameAuthor     string
	Description    template.HTML // Используем template.HTML для вывода сырого HTML
	URLAvatar      string
	URLRepository  string
	Stack          []Tools
	Links          []Urls
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
	// 1. Создаем набор расширений для парсера (заголовки, списки, ссылки и т.д.)
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)

	// 2. Парсим Markdown в документ
	doc := p.Parse([]byte(mdContent))

	// 3. Создаем HTML-рендерер с настройками по умолчанию
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	// 4. Рендерим документ в байты HTML
	output := markdown.Render(doc, renderer)

	return output
}

func main() {
	// 1. Используем Raw-ссылку на README.md
	rawReadmeURL := "https://raw.githubusercontent.com/my-app-s/go-generator/main/README.md"

	// Скачиваем сырой текст README
	mdText, err := fetchReadme(rawReadmeURL)
	if err != nil {
		fmt.Println("Предупреждение: не удалось загрузить README, используем текст по умолчанию:", err)
		mdText = "Описание временно недоступно."
	}

	// Конвертируем Markdown в HTML
	htmlDescription := convertMarkdownToHTML(mdText)

	// 2. Данные сайта
	data := PageData{
		NameRepository: "go-generate",
		NameAuthor:     "my-app-s(M.R.E)",
		Description:    template.HTML(htmlDescription), // Передаем готовый HTML
		URLAvatar:      "https://avatars.githubusercontent.com/u/94853425?v=4",
		URLRepository:  "https://github.com/my-app-s/go-generator",
		Stack: []Tools{
			{"Go"},
			{"HTML5"},
			{"CSS3"},
		},
		Links: []Urls{
			{"GitHub", "https://github.com/my-app-s"},
			{"Landing", "https://my-app-s.github.io/web-welcome"},
			{"LinkedIn", "https://www.linkedin.com/in/rustem-m-692916334"},
			{"HH", "https://hh.kz/resume/82ec45adff0f0ff5f60039ed1f6f3448515845"},
		},
	}

	// 3. Читаем шаблон
	tmpl, err := template.ParseFiles("templates/my-app-s.html")
	if err != nil {
		panic(err)
	}

	// 4. Создаем папку dist
	os.MkdirAll("dist", 0755)

	// 5. Создаем итоговый файл
	f, err := os.Create("dist/index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 6. Рендерим данные в файл
	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}

	println("Готово! Сайт сгенерирован в папку /dist с подгруженным README")
}
