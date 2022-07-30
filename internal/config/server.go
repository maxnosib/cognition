package config

import (
	"fmt"
	"os"
)

type Server struct {
	Port string
}

func NewServer() Server {
	return Server{
		Port: os.Getenv("PORT"),
	}
}

func (s Server) Validate() error {
	if s.Port == "" {
		return fmt.Errorf("port is empty")
	}

	return nil
}
