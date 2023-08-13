package bot

import "botyard/internal/entities/bot"

type createBody struct {
	bot.Bot
	Id struct{} `json:"-"`
}

type editBody struct {
	bot.Bot
	Commands struct{} `json:"-"`
	Id       struct{} `json:"-"`
}

type commandsBody struct {
	Commands []bot.Command
}

type commandBody struct {
	Alias string
}
