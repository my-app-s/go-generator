# HTML Generator in Go

![Go Version](https://img.shields.io/badge/Go-1.25.6-blue.svg)
![Tailwind Version](https://img.shields.io/badge/Tailwind-v4.3.3-blue.svg)
![License](https://img.shields.io/badge/License-GNU%20AGPLv3-red.svg)
![Status](https://img.shields.io/badge/Status-Stable-green)
![Latest Tag](https://img.shields.io/github/v/tag/my-app-s/go-generator)

> Status Github Actions
> 
> ![Status GitHub Pages](https://github.com/my-app-s/go-generator/actions/workflows/deploy-pages.yml/badge.svg)
[![GitHub Actions Status](https://img.shields.io/github/actions/workflow/status/my-app-s/go-generator/deploy-pages.yml?label=Action%3A%20uses%20v1&logo=github)](https://github.com/my-app-s/go-generator/actions)

## Описание

Простой и быстрый локальный генератор статического лендинга на Go с поддержкой Tailwind CSS v4, реализованный также как композитный экшен для GitHub Actions для автоматического деплоя в GitHub Pages без лишних Node.js зависимостей.

### 🛠 Features
- **Node.js-free Tailwind v4:** Компиляция стилей через standalone-бинарник Tailwind CLI.
- **Local README Parsing:** Автоматически читает и конвертирует `README.md` в чистый HTML.
- **Fast performance:** Генерация страницы занимает микросекунды благодаря Go.
- **GitHub Pages Ready:** Автоматическая сборка и деплой через GitHub Actions.

### 🎨 Визуализация процесса
- **Input:** Конфигурация в `config.json`, `README.md` и разметка с Tailwind v4.
- **Processing:** Чтение файлов, парсинг Markdown (`gomarkdown`), рендеринг через шаблоны Go и компиляция стилей Tailwind v4.
- **Output:** Готовый оптимизированный `.html` и стилизованный `output.css` в директории `/dist`.

## 🚀 Performance & Core Web Vitals

Проект развернут на GitHub Pages и оптимизирован для максимальной скорости работы без лишнего клиентского JavaScript. Метрики по результатам контрольного замера продакшена (Chrome DevTools):

| Метрика | Значение | Оценка |
| :--- | :--- | :--- |
| **Largest Contentful Paint (LCP)** | 0.33s | Отлично |
| **Cumulative Layout Shift (CLS)** | 0.00 | Идеально |
| **Interaction to Next Paint (INP)** | 32ms | Отлично |

![LCP](https://img.shields.io/badge/LCP-0.33s-brightgreen?style=flat-square&logo=googlechrome)
![CLS](https://img.shields.io/badge/CLS-0.00-brightgreen?style=flat-square&logo=googlechrome)
![INP](https://img.shields.io/badge/INP-32ms-brightgreen?style=flat-square&logo=googlechrome)

## 🛠 Tech Stack

- **Core:** Go (Golang) — высокая скорость компиляции и генерации статики.
- **Styles:** Tailwind CSS v4 — современная утилитарная стилизация без тяжелого раздувания.
- **CI/CD:** GitHub Actions — полностью автоматизированный пайплайн сборки и деплоя на GitHub Pages.

## 🚀 Инструкции

Примеры есть по пути `templates/examples/` в репозитории.

### 📦 Локальная генерация (OS Linux)

Для локальной генерации лендинга необходимо:
- скачать репозиторий командой `git clone https://github.com/my-app-s/go-generator.git`
- перейти в директорию скачанного `go-generator`
- обновить зависимости командой `go mod tidy`
- запустить `main.go`

Команды для выполнения в терминале:

```Bash
# скачать репозиторий
git clone [https://github.com/my-app-s/go-generator.git](https://github.com/my-app-s/go-generator.git)
# перейти в директорию репозитория
cd go-generator
# обновить зависимости
go mod tidy
# запустить
go run main.go

```

* после создастся локальная директория `/dist` в которой будет сгенерированный файл `index.html`
* перейти в директорию `/dist` можно командой `cd dist`
* посмотреть содержимое директории можно командой `ls`

Команды для выполнения в терминале:

```Bash
# перейти в директорию dist
cd dist
# посмотреть содержимое директории
ls

```

### 📦 CI деплой для GitHub Page (GitHub Actions)

Для автоматического деплоя с помощью GitHub Actions в качестве CI для GitHub Pages рекомендуется использовать готовый композитный экшен:

```yml
name: Deploy to GitHub Pages

on:
  push:
    branches: [ main ]
    paths-ignore:
      - 'LICENSE'
      - '.gitignore'
      - '.github/workflows/**'
  workflow_dispatch:

permissions:
  contents: read
  pages: write
  id-token: write

concurrency:
  group: "pages"
  cancel-in-progress: true

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Landing Repository
        uses: actions/checkout@v7

      - name: Generate Static Site
        uses: my-app-s/go-generator@v1

      - name: Setup Pages
        uses: actions/configure-pages@v6

      - name: Upload artifact
        uses: actions/upload-pages-artifact@v5
        with:
          path: 'dist'
          retention-days: 1

  deploy:
    environment:
      name: github-pages
      url: ${{ steps.deployment.outputs.page_url }}
    runs-on: ubuntu-latest
    needs: build
    steps:
      - name: Deploy to GitHub Pages
        id: deployment
        uses: actions/deploy-pages@v5

```

* создать в корне репозитория файл `config.json`

Команды для выполнения в терминале:

```Bash
# создание config.json
touch config.json

```

* скопировать пример **config** ниже в созданный файл `config.json`

> [!IMPORTANT]
> В файле config.json обязательно заполнить данные.
> Редактировать по правилам синтаксиса JSON.

```json
{
  "name_repository": "название репозитория",
  "name_author": "имя автора",
  "url_avatar": "ссылка на аватар",
  "url_repository": "ссылка на репозиторий",
  "stack": [
    {"name": "имя навыка"},
    {"name": "имя навыка"},
    {"name": "имя навыка"}
  ],
  "links": [
    {"name": "название", "url": "ссылка"},
    {"name": "название", "url": "ссылка"},
    {"name": "название", "url": "ссылка"},
    {"name": "название", "url": "ссылка"}
  ],
  "copyright_year": 2026
}

```

* в корне должен быть `README.md` (не обязательно, но на сайте будет выведено *«Описание временно недоступно»*)
* выполнить коммит
* выполнить push

## Realization Action

### Example use in deploy:

Стандартное использование:

```yaml
- name: Generate Static Site
  uses: my-app-s/go-generator@v1

```

Или использование с кастомными путями:

```yaml
- name: Generate Static Site
  uses: my-app-s/go-generator@v1
  with:
    config: 'custom-config.json'
    readme: 'docs/MAIN_README.md'

```

### 🔄 Обновление генератора в лендингах

Если обновлен код в репозитории `go-generator` (например, изменился шаблон или логика Tailwind), то созданный лендинг подтянет изменения при следующем деплое. Чтобы принудительно запустить пересборку без изменения файлов лендинга, необходимо выполнить пустой коммит:

> [!IMPORTANT]
> Предварительно нужно удалить `'.github/workflows/**'` из строки `paths-ignore` в файле `deploy.yml`, иначе пуш изменений воркфлоу не триггерит сборку лендинга.

```bash
git commit --allow-empty -m "ci: trigger rebuild with latest generator template"
git push

```

## Disclaimer & License

* **Short Disclaimer (EN)**: Materials are provided ***as is*** under the LICENSE file. No warranties. Authors are not liable for damages. No partnership or obligations created.
* **Short Disclaimer (RU)**: Материалы предоставляются ***как есть*** и регулируются файлом LICENSE. Гарантий нет. Автор(ы) не несут ответственности за убытки. Партнёрство или обязательства не создаются.
* **Full Disclaimer**: Read the full text in the [DISCLAIMER](https://www.google.com/search?q=./DISCLAIMER.md) (Available in EN/RU).
* **License**: This project is dual-licensed:
* **Open Source**: Licensed under the [GNU AGPLv3](https://www.google.com/search?q=./LICENSE).
* **Commercial**: A separate proprietary commercial license is required for proprietary, closed-source, or enterprise use that does not comply with AGPLv3 terms. Contact the copyright holder for commercial licensing.



## Author & Contacts

* **GitHub**: [@my-app-s](https://github.com/my-app-s)
* **LinkedIn**: [In/my-app-s](https://www.linkedin.com/in/my-app-s)
* **Mail**: [myapps.mre.dev@gmail.com](mailto:myapps.mre.dev@gmail.com)
