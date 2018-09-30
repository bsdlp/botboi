package handlers

import "github.com/bwmarrin/discordgo"

var userID string

func botUserID(session *discordgo.Session) (string, error) {
	if userID != "" {
		return userID, nil
	}
	user, err := session.User("@me")
	if err != nil {
		return "", err
	}

	userID = user.ID
	return userID, nil
}
