# Sub Service

Микросервис для управления подписками, разработанный на Go с использованием Echo framework и PostgreSQL.
🚀 Технологии

    Go 1.21+

    Echo - высокопроизводительный веб-фреймворк

    PostgreSQL - реляционная база данных

    Logrus - структурированное логирование

    Swagger - документация API

    Docker - контейнеризация

📦 Установка и запуск
## Требования

    Go 1.21 или выше

    PostgreSQL 12+

    Docker (опционально)

## Локальная установка

Клонируйте репозиторий:

bash

git clone https://github.com/Darkhan-Sol0/sub_service.git
cd sub_service

## Установите зависимости:

bash

go mod download

## Настройте базу данных:

sql

CREATE DATABASE sub_service;
CREATE USER sub_user WITH PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE sub_service TO sub_user;

## Создайте файл конфигурации config.yaml:

yaml

log_level: "info"

server_http:
  address: "0.0.0.0:8080"
  session_timeout: 4s
  idle_timeout: 60s

database:
  database_env: postgres
  port: 5432
  host: localhost
  database_name: sub_service
  username: sub_user
  password: password

## Запустите приложение:

bash

go run cmd/main.go

Запуск через Docker

## Соберите и запустите контейнеры:

bash

docker-compose up -d

## Приложение будет доступно по адресу: http://localhost:8080

📖 API Документация

После запуска приложения документация Swagger доступна по адресу:
text

http://localhost:8080/swagger/index.html

## 🔧 API Endpoints
Подписки

    GET / - Тестовый endpoint

    POST /add_sub - Добавление новой подписки

    GET /get_sub_by_id/:id - Получение подписки по ID

    GET /get_list - Получение всех подписок

    GET /get_list_by_user/:uuid - Получение подписок пользователя

    GET /get_price_subs - Получение стоимости подписки по фильтрам

    PATCH /update_sub - Обновление подписки

    DELETE /delete_sub - Удаление подписки

Примеры запросов

Добавление подписки:
bash

curl -X POST http://localhost:8080/add_sub \
  -H "Content-Type: application/json" \
  -d '{
    "service_name": "Netflix",
    "price": 999,
    "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
    "start_date": "2024-01-01",
    "month": 12
  }'

Получение подписки по ID:
bash

curl http://localhost:8080/get_sub_by_id/1

Получение стоимости подписки:
bash

curl "http://localhost:8080/get_price_subs?serv=YandexGold&uuid=60601fee-2bf1-4721-ae6f-7636e79a0cba&sdate=01-2024&edate=12-2024"

⚙️ Конфигурация

Сервис использует YAML-конфигурацию. Основные параметры:

    log_level - Уровень логирования (trace, debug, info, warn, error, fatal, panic)

    server_http.address - Адрес и порт для HTTP сервера

    database - Параметры подключения к PostgreSQL

📊 Логирование

Сервис использует структурированное логирование через Logrus. Логи выводятся в формате JSON и включают:

    Временные метки

    Уровни логирования

    Контекстные поля для трассировки запросов

🗄️ Структура проекта
text

sub_service/<br>
├── cmd/<br>
│   └── main.go              # Точка входа<br>
├── internal/<br>
│   ├── config/              # Конфигурация приложения<br>
│   ├── dto/                 # Data Transfer Objects<br>
│   ├── logger/              # Настройка логгера<br>
│   ├── service/             # Бизнес-логика<br>
│   └── web/                 # HTTP handlers и роутинг<br>
├── docs/                    # Swagger документация<br>
├── config.yaml              # Файл конфигурации<br>
├── Dockerfile               # Конфигурация Docker<br>
├── docker-compose.yml       # Docker Compose<br>
└── go.mod                   # Зависимости Go<br>

🐛 Troubleshooting
common problems

    Ошибка подключения к базе данных

        Проверьте параметры подключения в config.yaml

        Убедитесь, что PostgreSQL запущен

    Swagger не отображается

        Убедитесь, что аннотации добавлены ко всем обработчикам

        Выполните swag init для обновления документации

    Проблемы с миграциями

        Проверьте права доступа к базе данных

        Убедитесь, что все необходимые таблицы созданы

🤝 Разработка
Добавление нового endpoint

    Добавьте обработчик в internal/web/routing.go

    Добавьте Swagger аннотации

    Обновите документацию: swag init

    Добавьте соответствующий метод в сервисный слой

Локальная разработка
bash

# Запуск с горячей перезагрузкой
go run cmd/main.go

# Генерация Swagger документации
swag init

# Запуск тестов
go test ./...

📝 Лицензия

Этот проект лицензирован под MIT License - смотрите файл LICENSE для деталей.
👨‍💻 Автор

Darkhan Solovyev - GitHub

Для дополнительной информации обращайтесь к документации API через Swagger UI после запуска приложения.