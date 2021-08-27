package main

import "github.com/SerjLeo/mlf_backend/internal/app"

const configPath = "config"


// @title My Local Financier API
// @version 1.0
// @description API for MLF application
// @license.name MIT

// @host localhost:8000
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app.Run(configPath)
}
