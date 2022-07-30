# основной образ их может быть несколько
FROM golang:1.18.4

# описание что это за образ и др метаданные
LABEL author="max"
LABEL version="1.0"
LABEL description="Реализация zettelkasten — система ведения заметок и управления личными знаниями, используемая в исследованиях и учебе"

# Выставляем переменную для тестов
ENV PORT=:8080
ENV PGPASSWORD=postgres
ENV PG_HOST=db_cognition
ENV PG_PORT=5432
ENV PG_USER=postgres
ENV PG_PWD=postgres
ENV PG_DB=cognition

RUN mkdir -p /app/
COPY . /app
WORKDIR /app


# RUN go get -u github.com/librun/migrago@v1.1.1

RUN make build build

EXPOSE 8080

ENTRYPOINT [ "sh", "scripts/start_app.sh" ]
