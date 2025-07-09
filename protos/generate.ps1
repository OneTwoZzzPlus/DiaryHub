Set-Location $PSScriptRoot

$SSO_SERVICE_DIR = "../sso-service/protos/gen/"
$FRONTEND_DIR = "../frontend/protos/gen/"

protoc -I proto ./proto/auth/auth.proto --go_out=$SSO_SERVICE_DIR --go_opt=paths=source_relative --go-grpc_out=$SSO_SERVICE_DIR --go-grpc_opt=paths=source_relative

protoc -I proto ./proto/auth/auth.proto --js_out=import_style=commonjs:$FRONTEND_DIR --grpc-web_out=import_style=commonjs,mode=grpcwebtext:$FRONTEND_DIR