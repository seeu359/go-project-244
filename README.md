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

Вывод (формат по умолчанию — stylish):

```
{
    common: {
      + follow: false
        setting1: Value 1
      - setting2: 200
      - setting3: true
      + setting3: null
      + setting4: blah blah
      + setting5: {
            key5: value5
        }
        setting6: {
            doge: {
              - wow:
              + wow: so much
            }
            key: value
          + ops: vops
        }
    }
    group1: {
      - baz: bas
      + baz: bars
        foo: bar
      - nest: {
            key: value
        }
      + nest: str
    }
  - group2: {
        abc: 12345
        deep: {
            id: 45
        }
    }
  + group3: {
        deep: {
            id: {
                number: 45
            }
        }
        fee: 100500
    }
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
