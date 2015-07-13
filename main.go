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

	r := bot.NewCommandRouter()

	r.AddRoute("/d", func(bot *tgbotapi.BotAPI, update tgbotapi.Update, cmd bot.Command) {
		user := update.Message.From.ID
		debtor := cmd.Part(0)
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

	r.AddRoute("/s", func(bot *tgbotapi.BotAPI, update tgbotapi.Update, cmd bot.Command) {
		user := update.Message.From.ID
		debtor := cmd.Part(0)
		newDebt, err := strconv.ParseFloat(cmd.Part(1), 64)

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

	r.AddRoute("/a", func(bot *tgbotapi.BotAPI, update tgbotapi.Update, cmd bot.Command) {
		user := update.Message.From.ID
		res, err := accountant.UserDebtors(strconv.FormatInt(int64(user), 10))

		if err != nil {
			log.Println(err.Error())
			bot.SendMessage(tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка"))
			return
		}

		bot.SendMessage(tgbotapi.NewMessage(update.Message.Chat.ID, "Вам должны:"))
		for debtor, debt := range res {
			bot.SendMessage(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Долг %v - %.2f", debtor, debt)))
		}
	})

	r.AddRoute("/about", func(bot *tgbotapi.BotAPI, update tgbotapi.Update, cmd bot.Command) {
		bot.SendMessage(tgbotapi.NewMessage(update.Message.Chat.ID, "/d [name] [money] - добавить долг\n/s [name] [money] - установить долг\n/a - список должников"))
	})

	bot.RunWithHandler(r.Handler())

}