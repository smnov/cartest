package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/smnov/cartest/docs"
)

func main() {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))
	if err := godotenv.Load(); err != nil {
		l.Info("No .env file found")
	}
	db, err := NewPostgresStore()
	if err != nil {
		panic(err.Error())
	}
	parsePort, exists := os.LookupEnv("SERVER_PORT")
	port := string(":" + parsePort)
	if !exists {
		l.Info("port variable not found, using default instead")
		port = ":8080"
	}
	s := NewServer(port, db, l)
	s.Start()
}
