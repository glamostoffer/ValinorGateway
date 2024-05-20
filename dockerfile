# Установка базового образа
FROM golang:1.22

# Установка рабочей директории в контейнере
WORKDIR /app

# Копирование файлов проекта в контейнер
ADD . /app

# Установка зависимостей
RUN go mod download

# Компиляция проекта
RUN go build -o main ./cmd/main.go

# Создание директории config и файла config.yaml
RUN mkdir -p ./config && echo "\
env: 'local'\n\
start_timeout: 5s\n\
stop_timeout: 5s\n\
http_server:\n\
  port: 8080\n\
  host: 'localhost'\n\
  stop_timeout: 5s\n\
auth: \n\
  Address: 'localhost:6000'\n\
  DontRun: false\n\
chat: \n\
  Address: 'localhost:6001'\n\
  DontRun: false\n\
route_config: \n\
  access_token_ttl: 1h\n\
" > ./config/config.yaml

# Установка переменной окружения CONFIG_PATH
ENV CONFIG_PATH=./config/config.yaml

# Запуск приложения
CMD ["./main"]