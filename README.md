# DiaryHub
Учебный проект

## Запуск
- Сборка и запуск `docker-compose up --build`
- Запуск `docker compose up`

## Микросервисы diaryhub
|     Имя     | Технология | Источник  | Описание |
| ----------- | ---------- | --------- | -------- |
| db          | PostgreSQL | DockerHub | Хранилище - реляционная база данных
| envoy-proxy | Envoy      | DockerHub | GRPC-web адаптер и API gateway
| sso-service | Go         | self      | Сервис для регистрации, авторизации и управления правами
| frontend    | JavaScript | self      | Интерфейс пользователей

## sso-service
- auth *service*
- gRPC server *app*
- storage *app*

Отдельный запуск sso-service (необходима запущенная БД).
```powershell
cd ./sso-service
$env:CONFIG_PATH="./config/local.yml
go run .
```

## frontend
На текущей стадии это простейший скрипт GRPC-web клиента.
```powershell
cd ./frontend
node client.js
```

## protos
Автогенерация PROTOBUF файлов.
```powershell
./protos/generate.ps1
```