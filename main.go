package main

import (
	"html/template"
	"os"
)

// Описываем структуру данных для сайта
type Repo struct {
	Title       string
	Description string
	URL         string
}

type PageData struct {
	Name  string
	Bio   string
	Repos []Repo
}

func main() {
	// 1. Данные твоего сайта (потом можно вынести в JSON/YAML)
	data := PageData{
		Name: "my-app-s",
		Bio:  "Go разработчик в процессе обучения. Люблю чистый код и быстрые сайты.",
		Repos: []Repo{
			{"Project-One", "Крутой проект на Go", "https://github.com/my-app-s/go-generator"},
			{"Project-Two", "Go Heart Bot", "https://github.com/my-app-s/go-heart-bot"},
			{"Project-Three", "Go Simple Router with Recovery", "https://github.com/my-app-s/go-custom-router"},
		},
	}

	// 2. Читаем шаблон
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}

	// 3. Создаем папку docs (если нет), чтобы GitHub Pages её подхватил
	os.MkdirAll("docs", 0755)

	// 4. Создаем итоговый файл
	f, err := os.Create("docs/index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 5. Рендерим данные в файл
	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}

	println("Готово! Сайт сгенерирован в папку /docs")
}
