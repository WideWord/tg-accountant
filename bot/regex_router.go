package bot

import(
	"regexp"
	"github.com/Syfaro/telegram-bot-api"
)

type RegexpRouter struct {
	routes map[string]CommandHandler
}

func NewRegexRouter() *RegexpRouter {
	return &RegexpRouter{
		routes: make(map[string]CommandHandler),
	}
}

func (r *RegexpRouter) AddRoute(pattern string, handler CommandHandler) {
	r.routes[pattern] = handler
}

func (r *RegexpRouter) Handler() CommandHandler {
	return func(bot *tgbotapi.BotAPI, update tgbotapi.Update, cmd Command) {
		for pattern, handler := range r.routes {
			matches, err := regexp.MatchString(pattern, update.Message.Text)
			if err == nil && matches {
				handler(bot, update, cmd)
				break
			}
		}
	}
}
