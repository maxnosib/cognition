package repository

import (
	"os"
	"strconv"

	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/v4"
)

type Model struct {
	*pgx.Conn
}

func New() (Model, error) {
	str := os.Getenv("PG_PORT")
	port, err := strconv.Atoi(str)
	if err != nil {
		return Model{}, err
	}

	cfg := pgx.ConnConfig{
		Host:     os.Getenv("PG_HOST"),
		Port:     uint16(port),
		Database: os.Getenv("PG_DB"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PWD"),
	}

	db, err := pgx.Connect(cfg)
	if err != nil {
		return Model{}, err
	}

	return Model{db}, nil
}
