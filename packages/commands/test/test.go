package test

import (
	"github.com/bwmarrin/discordgo"
	"goat/packages/handlers/custom_error"
)

func Handle(session *discordgo.Session, message *discordgo.MessageCreate, parameters []string) {
	_, err := session.ChannelMessageSend(message.ChannelID, "Stop testing me you fucktard.")

	custom_error.Handle(err, "There was an error sending the message.")
}
