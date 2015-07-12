package bot

import(
	"../config"
	"github.com/Syfaro/telegram-bot-api"
	"log"
)

func RunWithHandler(handler CommandHandler) {
	bot, err := tgbotapi.NewBotAPI(config.Get().Bot.Token)

	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	updates, err := bot.UpdatesChan(u)

	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		cmd := ParseCommand(update.Message.Text)
		handler(bot, update, cmd)
	}
}