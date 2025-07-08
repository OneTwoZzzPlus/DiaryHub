# DiaryHub
Учебный проект

Сборка `docker compose build`

Запуск `docker compose up`

Микросервисы diaryhub:
- sso-service


# sso-service
Состав sso-service:
- auth *service*
- gRPC server *app*
- storage *app*


Отдельный запуск sso-service
```
cd ./sso-service
$env:CONFIG_PATH="./config/local.yml
go run .
```

PROTOC
```
protoc -I proto ./proto/auth/auth.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative
```