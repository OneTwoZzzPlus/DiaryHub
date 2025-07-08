FROM golang:1.24-alpine
WORKDIR /app
COPY . .
RUN go build -o sso-service .
CMD ["/app/sso-service"]