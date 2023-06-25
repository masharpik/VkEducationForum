FROM golang:1.20.5-alpine3.18 as builder

WORKDIR /app

# Копируем все файлы из текущей директории хоста в рабочую директорию в контейнере
COPY . .

# Скачиваем зависимости
RUN go mod download

RUN go build -o main ./forum

FROM postgres:15.3-alpine3.18

WORKDIR /app

ENV POSTGRES_USER=forum_user
ENV POSTGRES_PASSWORD=forum_pass
ENV POSTGRES_DB=forum_name

COPY db/db.sql /docker-entrypoint-initdb.d/

# Копируем исполняемый файл из предыдущего образа
COPY --from=builder /app/main main

COPY scripts/start.sh start.sh
RUN chmod +x start.sh

COPY .env .env

# Запускаем исполняемый файл
CMD ["./start.sh"]
