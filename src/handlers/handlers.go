package handlers

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

// Handler describes a handler
type Handler interface {
	Register(*discordgo.Session)
}

// Loader holds all of the handlers
type Loader struct {
	Logger *zap.SugaredLogger
}

// Register registers all handlers
func (l *Loader) Register(session *discordgo.Session) {
	newGuildEmojiHandler(l.Logger).Register(session)
}
