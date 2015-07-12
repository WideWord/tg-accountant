package bot

import(
	"github.com/Syfaro/telegram-bot-api"
)

type CommandHandler func(*tgbotapi.BotAPI, tgbotapi.Update, Command)

type CommandRouter struct {
	handlers map[string]CommandHandler
}

func NewCommandRouter() *CommandRouter {
	return &CommandRouter{
		handlers: make(map[string]CommandHandler),
	}
}

func (r *CommandRouter) AddRoute(route string, handler CommandHandler) {
	r.handlers[route] = handler
}

func (r *CommandRouter) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update, cmd Command) {
	if handler, ok := r.handlers[cmd.Part(0)]; ok {
		handler(bot, update, cmd.Shift())
	}
}

func (r *CommandRouter) Handler() CommandHandler {
	return func(bot *tgbotapi.BotAPI, update tgbotapi.Update, cmd Command) {
		r.Handle(bot, update, cmd)
	}
}
