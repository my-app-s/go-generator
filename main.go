package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
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
	// 1. Читаем конфигурационный файл config.json
	configFile, err := os.Open("config.json")
	if err != nil {
		panic(fmt.Sprintf("не удалось открыть config.json: %v", err))
	}
	defer configFile.Close()

	var data PageData
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&data)
	if err != nil {
		panic(fmt.Sprintf("ошибка парсинга config.json: %v", err))
	}

	// Читаем локальный README.md, который лежит рядом с config.json
	mdBytes, err := os.ReadFile("README.md")
	var mdText string
	if err != nil {
		fmt.Println("Предупреждение: не удалось найти локальный README.md:", err)
		mdText = "Описание временно недоступно."
	} else {
		mdText = string(mdBytes)
	}
	
	htmlDescription := convertMarkdownToHTML(mdText)
	data.Description = template.HTML(htmlDescription) // Дописываем динамический HTML в структуру

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

	println("Готово! Сайт сгенерирован в папку /dist с данными из config.json")
}
