# DiaryHub
Учебный проект

- Сборка `docker compose build`
- Запуск `docker compose up`

Отдельный запуск auth-service
```
cd ./auth-service
$env:CONFIG_PATH="./config/local.yml
go run .
```