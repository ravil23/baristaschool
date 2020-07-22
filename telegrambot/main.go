package main

import (
	"github.com/ravil23/baristaschool/telegrambot/telegram"
)

func main() {
	bot := telegram.NewBot()
	bot.Init()
	bot.HealthCheck()
	bot.Run()
}
