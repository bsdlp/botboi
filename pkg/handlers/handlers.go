package handlers

import "github.com/bwmarrin/discordgo"

// Handler describes a handler
type Handler interface {
	Register(*discordgo.Session)
}

// Loader holds all of the handlers
type Loader struct {
	Handlers []Handler
}

// Register registers all handlers
func (l *Loader) Register(session *discordgo.Session) {
	for _, handler := range l.Handlers {
		session.AddHandler(handler)
	}
}

// New loads all the handlers
func New() (*Loader, error) {
	loader := &Loader{
		Handlers: []Handler{},
	}
	return loader, nil
}
