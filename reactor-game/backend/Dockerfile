# Базовый образ с Go для сборки
FROM golang:1.23-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и загружаем зависимости
COPY go.mod ./
RUN go mod download

# Копируем весь код
COPY . .

# Собираем приложение
RUN go build -o reactor-game-backend main.go
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
# Финальный образ (лёгкий, без Go)
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем скомпилированный бинарник из builder
COPY --from=builder /app/reactor-game-backend .

# Указываем порт, который будет открыт
EXPOSE 8080

# Команда для запуска приложения
CMD ["./reactor-game-backend"]