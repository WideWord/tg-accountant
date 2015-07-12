package main

import(
	"./bot"
	"./config"
	"github.com/Syfaro/telegram-bot-api"
)

func main() {
	config.Read()

	r := bot.NewCommandRouter()

	r.AddRoute("/about", func(bot *tgbotapi.BotAPI, update tgbotapi.Update, cmd bot.Command) {
		bot.SendMessage(tgbotapi.NewMessage(update.Message.Chat.ID, "HW!"))
	})

	bot.RunWithRouter(r)

}