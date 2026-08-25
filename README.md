# HTML Generator in Go

> [!NOTE]
> 
> ![Go Version](https://img.shields.io/badge/Go-1.25%2B-blue.svg)
> ![License](https://img.shields.io/badge/License-GNU%20AGPLv3-red.svg)
> ![Status](https://img.shields.io/badge/Status-Dev-orange)
> ![CI](https://img.shields.io/badge/CI-GitHub%20Actions-green)
> ![CI](https://github.com/my-app-s/go-generator/actions/workflows/deploy.yml/badge.svg)
![Latest Tag](https://img.shields.io/github/v/tag/my-app-s/go-generator)


## Описание

Простой и быстрый локальный генератор статического лендинга(статической страницы) на Go так же реализован как CI для GitHub Actions для автоматического деплоя в GitHub Pages.

### 🛠 Features
- **Local README Parsing:** Автоматически читает и конвертирует `README.md` в чистый HTML.
- **Fast performance:** Генерация страницы занимает микросекунды благодаря Go.
- **GitHub Pages Ready:** Автоматическая сборка с деплоем CI через GitHub Actions.

### 🎨 Визуализация процесса
- **Input:** Конфигурация в `config.json` и `README.md` репозитория.
- **Processing:** Чтение файлов, парсинг Markdown (`gomarkdown`) и рендеринг через шаблоны Go.
- **Output:** Готовый оптимизированный `.html` файл в директории `/dist`.

## 🚀 Инструкции

Примеры есть по пути `templates/examples/` в репозитории.

### 📦 Локальная генерация (OS Linux)

Для локальный генерации лендинга необходимо:
- скачать репозиторий командой `git clone https://github.com/my-app-s/go-generator.git`
- перейти в деректорию скачаного `go-generator`
- обновить зависимости командой `go mod tidy`
- запустить `main.go`

Команды для выполнения в терминале:

```Bash
# скачать репозиторий
git clone https://github.com/my-app-s/go-generator.git
# перейти в деректорию репозитория
cd go-generator
# обновить зависимости
go mod tidy
# запустить
go run main.go
```

- после создатся локальная директория `/dist` в которой будет сгенерированый файл `index.html`
- перейти в директории `/dist` можно командой `cd dist`
- посмотреть содержимое директории можно командой `ls`

Команды для выполнения в терминале:

```Bash
# перейти в деректорию dist
cd dist
# посмотреть содержимое директории
ls
```

### 📦 CI деплой для GitHub Page (GitHub Actions)

Для автоматического деплоя с помощью GitHub Actions как CI для GitHub Page необходимо следующее:
- создать в репозитории файл `deploy.yml` по пути `.github/workflows`

Команды для выполнения в терминале:

```Bash
# созлание деректорий .github/workflows (ключ -p позволяет создать полный путь директорий)
mkdir -p `.github/workflows`
# созлание deploy.yml
touch .github/workflows/deploy.yml
```

- скопировать пример **deploy** ниже в созданный файл `deploy.yml`

```yml
name: Deploy to GitHub Pages

on:
  push:
    # Запускаем workflows при пуше в ветку main или вручную
    branches: [ main ]
    # Игнорируем файлы при изменении которых деплой не будет запускаться
    paths-ignore:
      - 'LICENSE'
      - '.gitignore'
      - '.github/workflows/**'
  workflow_dispatch:

# Устанавливаем права
permissions:
  contents: read
  pages: write
  id-token: write

# Разрешаем только один одновременный деплой
concurrency:
  group: "pages"
  cancel-in-progress: true

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Landing Repository (конфиг и редми)
        uses: actions/checkout@v7
        with:
          path: landing-source

      - name: Checkout Generator Repository (исходники генератора на Go)
        uses: actions/checkout@v7
        with:
          repository: 'my-app-s/go-generator'
          path: generator-source

      - name: Set up Go
        uses: actions/setup-go@v7
        with:
          go-version: '1.25'
          cache: false

      - name: Install dependencies
        run: |
          cd generator-source
          go mod tidy

      - name: Build binaries
        run: |
          cd generator-source
          CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o go-generator .
          chmod +x ./go-generator

      - name: Copy binaries
        run: cp ./go-generator ../landing-source

      - name: Run Go Generator
        run: |
          cd landing-source
          ./go-generator

      - name: Setup Pages
        uses: actions/configure-pages@v6

      - name: Upload artifact
        uses: actions/upload-pages-artifact@v5
        with:
          path: 'landing-source/dist'
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

Альтернативный вариант использование как Action:

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
      - name: Checkout Landing Repository (конфиг и редми)
        uses: actions/checkout@v7

      # Использование my-app-s/go-generator@1 action
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

- создать в корне репозитория файл `config.json`

Команды для выполнения в терминале:

```Bash
# создание config.json
touch config.json
```

- скопировать пример **config** ниже в созданный файл `config.json`

> [!IMPORTANT]
>
> В файле config.json обязательно заполнить
> Редактировать по правилам синтаксиса JSON

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
  ]
}
```

- в корне должен быть `README.md` (не обязательно но на сайте будет выведено `Описание временно недоступно`)
- выполнить коммит
- выполнить push

## Realization Action

### Example use in deploy:

Если стандартно по инструкции:

```yaml
- name: Generate Static Site
  uses: my-app-s/go-generator@v1
```

Или используется кастомный подход:

```yaml
- name: Generate Static Site
  uses: my-app-s/go-generator@v1
  with:
    config: 'custom-config.json'
    readme: 'docs/MAIN_README.md'
```

### 🔄 Обновление генератора в лендингах

Если обновлен код в репозиторий `go-generator` (например, изменил шаблон HTML), то созданный лендинг подтянет изменения при следующем деплое. Чтобы принудительно запустить пересборку без изменения файлов лендинга, необходимо выполнить пустой коммит:

> [!IMPORTANT]
>
> Но нужно из `paths-ignore` в `deploy.yml` удалить `'.github/workflows/**'` если этого не сделать **обновление генератора в лендингах** не сработает.

```bash
git commit --allow-empty -m "ci: trigger rebuild with latest generator template"
git push

```

## Disclaimer & License

- **Short Disclaimer (EN)**: Materials are provided ***as is*** under the LICENSE file. No warranties, no rights granted unless explicitly stated. Authors are not liable for damages. No partnership or obligations created.
- **Short Disclaimer (RU)**: Материалы предоставляются ***как есть*** и регулируются LICENSE. Гарантий нет, права не передаются без явного указания. Автор(ы) не несут ответственности. Партнёрство или обязательства не создаются.
- **Full Disclaimer**: Read the full text in the [DISCLAIMER.md](https://github.com/my-app-s/my-app-s/blob/main/DISCLAIMER.md) (Available in EN/RU).
- **License**: Distributed under the [GNU AGPLv3](https://github.com/my-app-s/go-generator/blob/main/LICENSE) license.

## Author & Contacts

- **GitHub**: [@my-app-s](https://github.com/my-app-s)
- **LinkedIn**: [In/my-app-s](https://www.linkedin.com/in/my-app-s)
