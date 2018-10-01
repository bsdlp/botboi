package handlers

import "github.com/bwmarrin/discordgo"

var user *discordgo.User

func botUser(session *discordgo.Session) (*discordgo.User, error) {
	if user != nil {
		return user, nil
	}
	var err error
	user, err = session.User("@me")
	return user, err
}
