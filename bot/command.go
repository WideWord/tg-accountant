package bot

import(
	"strings"
)

type Command struct {
	parts []string
	fullMessage string
}

func ParseCommand(message string) Command {
	return Command{
		parts: strings.Split(message, " "),
		fullMessage: message,
	}
}

func (cmd Command) Part(index int) string {
	if index >= len(cmd.parts) {
		return ""
	} else {
		return cmd.parts[index]
	}
}

func (cmd Command) FullMessage() string {
	return cmd.fullMessage
}

func (cmd Command) Shift() Command {
	if len(cmd.parts) >= 1 {
		return Command{
			parts: cmd.parts[1:],
			fullMessage: cmd.fullMessage,
		}
	} else {
		return Command{
			parts: make([]string, 0),
			fullMessage: cmd.fullMessage,
		}
	}
}
