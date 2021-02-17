package api

import (
	dgo "github.com/bwmarrin/discordgo"
)

type User struct {
	userID string
	isBot  bool
	c      *Connection
}

func (c *Connection) GetUser(userID string) User {
	user, _ := c.client.User(userID)
	return User{
		userID: userID,
		isBot:  user.Bot,
		c:      c,
	}
}

func (u *User) IsOnline(guildID string) bool {
	userPresence, _ := u.c.client.State.Presence(guildID, u.userID)
	if userPresence != nil {
		if userPresence.Status == dgo.StatusOnline {
			return true
		}
	}
	return false
}

func (u *User) IsBot() bool {
	return u.isBot
}

func (u *User) Mention() string {
	return "<@" + u.userID + ">"
}
