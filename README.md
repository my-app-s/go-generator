Инструкцию в этом `README.md` нужно обновить, так как мы перешли с удаленной загрузки README по ссылке на **локальное чтение файла** прямо из папки проекта.

Вот готовый исправленный текст для твоего `README.md`. Замени им содержимое файла, и документация снова станет актуальной:

```markdown
# HTML Generator in Go

> [!NOTE]
> Простой и быстрый генератор статических страниц на Go с автоподгрузкой локального `README.md`.
> 
> ![Go Version](https://img.shields.io/badge/go-1.22%2B-blue.svg)
> ![License](https://img.shields.io/badge/license-GNU%20AGPLv3-red.svg)
> ![status: dev](https://img.shields.io/badge/status-dev-orange)

## 🛠 Features
- **Local README Parsing:** Автоматически читает и конвертирует локальный `README.md` в чистый HTML для лендинга.
- **Fast performance:** Генерация страницы занимает микросекунды благодаря Go.
- **GitHub Pages Ready:** Автоматическая сборка через GitHub Actions и вывод в папку `/dist`.

## 🎨 Визуализация процесса
- **Input:** Конфигурация в `config.json` и локальный `README.md` репозитория.
- **Processing:** Чтение файлов, парсинг Markdown (`gomarkdown`) и рендеринг через шаблоны Go.
- **Output:** Готовый оптимизированный `.html` файл в директории `/dist`.

## 🚀 Быстрый старт
### 📦 Installation
```Bash
git clone [https://github.com/my-app-s/go-generator.git](https://github.com/my-app-s/go-generator.git)
cd go-generator
go mod tidy

```

### ⚙️ Configuration

Перед запуском убедись, что в корне проекта созданы файлы `config.json` и `README.md`:

#### Example (`config.json`)

```json
{
  "name_repository": "go-generate",
  "name_author": "my-app-s(M.R.E)",
  "url_avatar": "[https://avatars.githubusercontent.com/u/94853425?v=4](https://avatars.githubusercontent.com/u/94853425?v=4)",
  "url_repository": "[https://github.com/my-app-s/go-generator](https://github.com/my-app-s/go-generator)",
  "stack": [
    {"name": "Go"},
    {"name": "HTML5"},
    {"name": "CSS3"}
  ],
  "links": [
    {"name": "GitHub", "url": "[https://github.com/my-app-s](https://github.com/my-app-s)"},
    {"name": "Landing", "url": "[https://my-app-s.github.io/web-welcome](https://my-app-s.github.io/web-welcome)"},
    {"name": "LinkedIn", "url": "[https://www.linkedin.com/in/rustem-m-692916334](https://www.linkedin.com/in/rustem-m-692916334)"},
    {"name": "HH", "url": "[https://hh.kz/resume/82ec45adff0f0ff5f60039ed1f6f3448515845](https://hh.kz/resume/82ec45adff0f0ff5f60039ed1f6f3448515845)"}
  ]
}

```

### 💻 Usage

```bash
go run main.go

```

## 🚀 Demo

Generated page view here:

[👉 Demo на GitHub Pages](https://my-app-s.github.io/go-generator/)

---

## 📜 Disclaimer

**English**: Materials are provided ***as is*** under the LICENSE file. No warranties, no rights granted unless explicitly stated. Authors are not liable for damages. No partnership or obligations created.  

**Русский**: Материалы предоставляются ***как есть*** и регулируются LICENSE. Гарантий нет, права не передаются без явного указания. Автор(ы) не несут ответственности. Партнёрство или обязательства не создаются.  

📌 See full disclaimer in [DISCLAIMER.md](https://github.com/my-app-s/my-app-s/blob/main/DISCLAIMER.md)

---

## 📜 License
Проект распространяется под лицензией **GNU AGPLv3**. Подробности в файле LICENSE.
