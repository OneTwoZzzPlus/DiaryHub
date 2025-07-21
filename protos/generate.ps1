Set-Location $PSScriptRoot

$SSO_SERVICE_DIR = "../sso-service/protos/gen/"
$FRONTEND_DIR = "../frontend/protos/gen/"

protoc -I proto ./proto/auth/auth.proto `
    --go_out=$SSO_SERVICE_DIR `
    --go_opt=paths=source_relative `
    --go-grpc_out=$SSO_SERVICE_DIR `
    --go-grpc_opt=paths=source_relative `
    --grpc-gateway_out=$SSO_SERVICE_DIR `
    --grpc-gateway_opt generate_unbound_methods=true `
    --openapiv2_out $SSO_SERVICE_DIR 

# protoc -I proto ./proto/auth/auth.proto `
#     --js_out=import_style=commonjs:$FRONTEND_DIR `
#     --grpc-web_out=import_style=commonjs,mode=grpcwebtext:$FRONTEND_DIR `
