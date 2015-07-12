package main

import(
	"./bot"
	"./config"
	"./db"
	"./accountant"
	"github.com/Syfaro/telegram-bot-api"
	"strconv"
	"log"
	"fmt"
)

func main() {
	config.Read()
	db.Connect()

	r := bot.NewRegexRouter()

	r.AddRoute("%[A-Za-z0-9_+=#$@]+ [+\\-]?[0-9]+", func(bot *tgbotapi.BotAPI, update tgbotapi.Update, cmd bot.Command) {
		user := update.Message.From.ID
		debtor := cmd.Part(0)[1:]
		change, err := strconv.ParseFloat(cmd.Part(1), 64)

		if err != nil {
			log.Println(err.Error())
			bot.SendMessage(tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка"))
			return
		}

		newDebt, err := accountant.AddDebt(strconv.FormatInt(int64(user), 10), debtor, change)

		if err != nil {
			log.Println(err.Error())
			bot.SendMessage(tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка"))
			return
		}

		bot.SendMessage(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Новый долг %v - %.2f", debtor, newDebt)))
	})

	r.AddRoute("%[A-Za-z0-9_+=#$@]+ =[+\\-]?[0-9]+", func(bot *tgbotapi.BotAPI, update tgbotapi.Update, cmd bot.Command) {
		user := update.Message.From.ID
		debtor := cmd.Part(0)[1:]
		newDebt, err := strconv.ParseFloat(cmd.Part(1)[1:], 64)

		if err != nil {
			log.Println(err.Error())
			bot.SendMessage(tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка"))
			return
		}

		err = accountant.SetDebt(strconv.FormatInt(int64(user), 10), debtor, newDebt)

		if err != nil {
			log.Println(err.Error())
			bot.SendMessage(tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка"))
			return
		}

		bot.SendMessage(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Новый долг %v - %.2f", debtor, newDebt)))
	})

	bot.RunWithHandler(r.Handler())

}