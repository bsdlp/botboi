package bot

import (
	"github.com/bsdlp/botboi/pkg/cfg"
	"github.com/bsdlp/botboi/pkg/handlers"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

// Bot implements the discord bot
type Bot struct {
	Config   cfg.Config
	Logger   *zap.Logger
	Handlers *handlers.Loader

	discord *discordgo.Session
}

// Run starts up the bot
func (bt *Bot) Run() error {
	if bt.discord == nil {
		dg, err := discordgo.New("Bot " + bt.Config.DiscordBotToken)
		if err != nil {
			return err
		}
		bt.discord = dg
	}

	bt.Handlers.Register(bt.discord)

	return bt.discord.Open()
}

// Stop closes the connection and resets the discord session
func (bt *Bot) Stop() error {
	if bt.discord == nil {
		return nil
	}
	err := bt.discord.Close()
	if err != nil {
		return err
	}

	bt.discord = nil
	return nil
}
