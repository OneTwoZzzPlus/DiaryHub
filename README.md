# DiaryHub
Учебный проект

Сборка и запуск `docker-compose up --build`

Запуск `docker compose up`

Микросервисы diaryhub:
- db (PostgreSQL)
- sso-service (go)

## sso-service
- auth *service*
- gRPC server *app*
- storage *app*


Отдельный запуск sso-service
```powershell
cd ./sso-service
$env:CONFIG_PATH="./config/local.yml
go run .
```

PROTOC
```powershell
cd ./sso-service/protos
protoc -I proto ./proto/auth/auth.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative
```