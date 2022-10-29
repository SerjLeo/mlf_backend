package main

import "github.com/SerjLeo/mlf_backend/internal/bot"

const configPath = "config"

func main() {
	bot.RunBot(configPath)
}
