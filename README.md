# URL Shortener Service (Go)

Сервис сокращения ссылок с поддержкой двух типов хранилища: PostgreSQL и in-memory.

## 🚀 Возможности

- Генерация коротких ссылок длиной 10 символов
- Поддержка символов: `a-z`, `A-Z`, `0-9`, `_`
- Одна оригинальная ссылка → одна сокращённая
- Два типа хранилища:
    - PostgreSQL
    - In-memory 
- HTTP API (Gin)
- Unit-тесты (service, repository, handler)
- Graceful Shutdown 
- Docker-ready

---

## ⚙️ Запуск

Запуск производится через docker-compose
```bash
docker-compose up
```

Для смены типа хранилища программы, следует отредактировать строку запуска контейнера приложения в файле Dockerfile

Команда для запуска контейнера с PostgreSQL:
```bash
CMD ["./server", "-n", "postgres"]
```
Команда для запуска с in memory хранилищем:
```bash
CMD ["./server", "-n", "memory"]
```

Приложение принимает два вида запросов:

1. POST запрос такого шаблона:
```bash
curl -X POST http://localhost:8080/api/shorter \
  -H "Content-Type: application/json" \
  -d '{"url":"https://google.com"}'
  ```

2. GET запрос такого шаблона:
```bash
curl -X GET http://localhost:8080/api/link/DoaqYGPx5l
```
Где ```DoaqYGPx5l``` является сокращенной ссылкой