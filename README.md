# DiaryHub
Учебный проект

## Запуск
- Сборка и запуск `docker-compose up --build`
- Запуск `docker compose up`

## Микросервисы diaryhub
|     Имя     | Технология | Источник  | Описание |
| ----------- | ---------- | --------- | -------- |
| db          | PostgreSQL | DockerHub | Хранилище - реляционная база данных
| sso-service | Go         | self      | Сервис для регистрации, авторизации и управления правами
| front       | TS Next.js | self      | Интерфейс пользователей

## sso-service
Сервис регистрации и авторизации.

- auth *service*
- gRPC server *app*
- REST gateway *app*
- email provider (smtp) *app*
- storage *app*

Запуск sso-service (необходима запущенная БД).
```powershell
cd ./sso-service
$env:CONFIG_PATH="./config/local.yml
go run .
```

## front
Простые страницы с формами ввода и кнопками, отправляющими запросы GRPC-web клиента.

Запуск front веб-сервиса.ы
```powershell
cd ./front
yarn dev
```

## protos
Автогенерация PROTOBUF файлов.
```powershell
cd ./protos
./generate.ps1
```

## TODO
- Create register page (front)
- Gorutine email
- Add confirmation of mail
- ...
