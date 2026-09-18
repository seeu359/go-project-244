# Вычислитель отличий на Go

[![hexlet-check](https://github.com/seeu359/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/seeu359/go-project-244/actions)
[![ci](https://github.com/seeu359/go-project-244/actions/workflows/ci.yml/badge.svg)](https://github.com/seeu359/go-project-244/actions/workflows/ci.yml)

Консольная утилита для сравнения вложенных структур (JSON, YAML)

Учебный проект Хекслета: https://ru.hexlet.io/programs/go


## Стек

- Go

## Установка

<!-- Опишите установку: клонирование, зависимости, переменные окружения -->

```bash
git clone https://github.com/seeu359/go-project-244.git
cd go-project-244
```

## Использование

```bash
# сборка
make build

# сравнение файлов
./bin/gendiff file1.json file2.json
```

Вывод:

```
{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}
```

<!-- TODO: добавить запись asciinema — [![asciicast](https://asciinema.org/a/ID.svg)](https://asciinema.org/a/ID) -->

---

<details>
<summary>Автоматические тесты Хекслета</summary>

Тесты запускаются на каждый коммит. За запуск отвечает файл `.github/workflows/hexlet-check.yml` — не удаляйте и не переименовывайте ни его, ни репозиторий.

</details>

## О Хекслете

[Хекслет](https://ru.hexlet.io/) — школа программирования: авторские программы обучения с практикой, поддержкой наставников и реальными проектами, которые остаются в резюме. Этот репозиторий — один из таких проектов.
