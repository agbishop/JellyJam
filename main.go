package main

import (
	"embed"
	"fmt"
	"github.com/agbishop/JellyJam/handlers"
	"github.com/agbishop/JellyJam/pkgs/env"
	"github.com/joho/godotenv"
)

//go:embed client/out/*
var uiFiles embed.FS

func main() {
	_ = godotenv.Load(".env")
	ip := env.MustEnv("CONTROLLER_IP")
	port := env.MustEnv("CONTROLLER_PORT")
	svc, shutdown := handlers.NewService(fmt.Sprintf("ws://%s:%s/ws/", ip, port), &uiFiles)
	defer shutdown()

	svc.Start()
}
