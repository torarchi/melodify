# Melodify - Music Generation Service

Melodify - это сервис для генерации музыки, использующий AI-модель Riffusion через API Replicate. Сервис построен с использованием принципов Clean Architecture и Domain-Driven Design.

## 🎵 Функциональность

- Генерация музыки на основе текстового описания
- Мониторинг статуса генерации
- Получение сгенерированного аудио
- RESTful API интерфейс

### Локальный запуск

1. Клонируйте репозиторий:
```bash
git clone https://github.com/torarchi/melodify.git
cd melodify
```

2. Создайте файл .env:
```bash
REPLICATE_API_TOKEN=
SERVER_PORT=8080
```

3. Установите зависимости:
```bash
go mod download
```

4. Запустите приложение:
```bash
go run cmd/main.go
```

### Запуск через Docker

1. Соберите образ:
```bash
docker-compose build
```

2. Запустите контейнер:
```bash
docker-compose up -d
```

## 📡 API Endpoints

### Создание предсказания
```bash
POST /predictions
Content-Type: application/json

{
    "prompt_b": "rock guitar solo"
}
```

### Получение статуса
```bash
GET /predictions/get?id=PREDICTION_ID
```

## 🎼 Примеры использования

### Генерация рок-соло
```bash
curl -X POST http://localhost:8080/predictions \
  -H "Content-Type: application/json" \
  -d '{"prompt_b": "energetic rock guitar solo"}'
```

### Проверка статуса
```bash
curl http://localhost:8080/predictions/get?id=YOUR_PREDICTION_ID
```

## ⚙️ Конфигурация

Конфигурация осуществляется через переменные окружения:

- `REPLICATE_API_TOKEN` - токен для доступа к API Replicate
- `SERVER_PORT` - порт для HTTP сервера (по умолчанию 8080)
