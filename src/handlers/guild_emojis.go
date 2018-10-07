package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

// `list emotes` returns all emojis in this guild
type guildEmoji struct {
	logger *zap.SugaredLogger
}

func (ge *guildEmoji) handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	ge.handleMessage(s, m.Message)
}

func (ge *guildEmoji) handleMessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	ge.handleMessage(s, m.Message)
}

func formatEmoji(emoji *discordgo.Emoji) string {
	if emoji.Animated {
		return fmt.Sprintf("<a:%s:%s>", emoji.Name, emoji.ID)
	}
	return fmt.Sprintf("<:%s:%s>", emoji.Name, emoji.ID)
}

func (ge *guildEmoji) handleMessage(s *discordgo.Session, m *discordgo.Message) {
	if m.Content != `list emotes` {
		return
	}

	emoji := ge.getEmojiFromChannel(s, m.ChannelID)
	if len(emoji) == 0 {
		return
	}

	msg := m.Author.Mention() + ":"
	for _, emo := range emoji {
		msg += " " + formatEmoji(emo)
	}

	_, err := s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		ge.logger.Errorw(
			"failure", err.Error(),
			"action", "guildEmoji ChannelMessageSend",
		)
	}
}

func (ge *guildEmoji) getEmojiFromChannel(s *discordgo.Session, channelID string) []*discordgo.Emoji {
	channel, err := s.Channel(channelID)
	if err != nil {
		ge.logger.Errorw(
			"failure", err.Error(),
			"action", "guildEmoji get channel",
		)
		return nil
	}

	guild, err := s.Guild(channel.GuildID)
	if err != nil {
		ge.logger.Errorw(
			"failure", err.Error(),
			"action", "guildEmoji get guild",
		)
		return nil
	}

	return ge.usableEmoji(s, guild.Emojis, guild.ID)
}

func checkEmojiRoleWhitelist(whitelist, roles []string) bool {
	if len(whitelist) == 0 {
		return true
	}

	for _, whitelisted := range whitelist {
		for _, role := range roles {
			if role == whitelisted {
				return true
			}
		}
	}
	return false
}

func filterEmojiRoleAccess(emoji []*discordgo.Emoji, roles []string) []*discordgo.Emoji {
	results := make([]*discordgo.Emoji, 0, len(emoji))
	for _, emo := range emoji {
		if checkEmojiRoleWhitelist(emo.Roles, roles) {
			results = append(results, emo)
		}
	}
	return results
}

func (ge *guildEmoji) usableEmoji(s *discordgo.Session, emoji []*discordgo.Emoji, guildID string) []*discordgo.Emoji {
	botUser, err := botUser(s)
	if err != nil {
		ge.logger.Errorw(
			"failure", err.Error(),
			"action", "guildEmoji botUser",
		)
		return nil
	}

	member, err := s.GuildMember(guildID, botUser.ID)
	if err != nil {
		ge.logger.Errorw(
			"failure", err.Error(),
			"action", "guildEmoji GuildMember",
		)
		return nil
	}

	return filterEmojiRoleAccess(emoji, member.Roles)
}

func (ge *guildEmoji) Register(session *discordgo.Session) {
	session.AddHandler(ge.handleMessageCreate)
	session.AddHandler(ge.handleMessageUpdate)
}

func newGuildEmojiHandler(logger *zap.SugaredLogger) *guildEmoji {
	return &guildEmoji{logger: logger}
}
