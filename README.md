## Стек технологий

#### API
- Fiber
#### Брокер сообщений
- NATS Streaming
#### Хранилище:
- PostgreSQL
- TTLCache
#### Логирование:
- slog

---

## Конечные точки

### Стартовая страница
```http request
GET http://localhost:8080/order
```

### Страница отображения данных заказа по UID
```http request
GET http://localhost:8080/order?uid=b563feb7b2b84b6test
```

---

## Развертывание

**Собрать** приложение

```shell
docker compose build
```

**Поднять** приложение

```shell
docker compose up -d
```

---

## Конфигурации

### Все параметры загружаются из файта **[.env](.env)**

```dotenv
LOGGER_LEVEL=debug

SERVER_ADDR=:8080

POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=wb

NATS_ADDR=nats://nats-streaming:4222

CLUSTER_ID=order-data
CLIENT_SUB_ID=order-sub
```