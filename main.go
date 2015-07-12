package main

import(
	"./bot"
	"./config"
	//"./db"
	"github.com/Syfaro/telegram-bot-api"
)

func main() {
	config.Read()
	//db.Connect()

	r := bot.NewRegexRouter()

	r.AddRoute("%[A-Za-z0-9_+=#$@]+ [+\\-]?[0-9]+", func(bot *tgbotapi.BotAPI, update tgbotapi.Update, cmd bot.Command) {
		bot.SendMessage(tgbotapi.NewMessage(update.Message.Chat.ID, cmd.Part(1)))
	})

	bot.RunWithHandler(r.Handler())

}