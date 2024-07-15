
# go-ml-vk-test-task

Этот репозиторий содержит проект `go-ml-vk-test-task`, который является сервисом для обработки документов, получаемых из очереди. Проект написан на языке Go и использует PostgreSQL в качестве базы данных.

## Структура проекта

- **main.go**: Исполняемый файл `main.go`, который является точкой входа в приложение.
- **internal/**: Внутренние пакеты приложения.
- **schema/**: Скрипты для инициализации и заполнения базы данных.
- **.gitignore**: Файл, указывающий Git игнорировать определенные файлы и директории.
- **Dockerfile**: Определяет образ Docker для сборки приложения.
- **docker-compose.yaml**: Файл для определения и запуска многоконтейнерных Docker приложений.

## Зависимости

- Go 1.22.4
- PostgreSQL
- Docker

## Установка и запуск

1. **Клонирование репозитория**:
```bash
    git clone https://github.com/DedMokus/go-ml-vk-test-task.git cd go-ml-vk-test-task
```

2. **Запуск приложения с помощью Docker Compose**:

```bash 
    docker-compose up --build
```

### Дополнительные ресурсы

- [Официальный сайт PostgreSQL](https://www.postgresql.org/)
- [Docker Hub для PostgreSQL](https://hub.docker.com/_/postgres)
