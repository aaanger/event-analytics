# Приложение для сбора и аналитики просмотров веб-страниц
- Общение между клиентом и сервером по gRPC
- База данных Clickhouse
- В качестве драйвера для БД использутся [clickhouse-go](github.com/ClickHouse/clickhouse-go/v2)
- Protobuf файлы для методов обработчика запросов по gRPC
- Образ gRPC сервера создается в Dockerfile и поднимается вместе с Clickhouse в Docker Compose
