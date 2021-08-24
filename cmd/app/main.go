package main

import "github.com/SerjLeo/mlf_backend/internal/app"

const configPath = "configs"

func main() {
	app.Run(configPath)
}