package command

import (
	"github.com/bwmarrin/discordgo"
	"goat/packages/commands/reddit"
	"os"
	"strings"
)

// All the available command that the bot is going to have access to.
var availableCommands = map[string]func(session *discordgo.Session, message *discordgo.MessageCreate, parameters []string){
	"reddit": reddit.GetSubreddit,
}

// Handle it will handle the commands called by the user.
func Handle(session *discordgo.Session, message *discordgo.MessageCreate) {
	explodedSentMessage := strings.Fields(message.Message.Content)

	if len(explodedSentMessage) == 0 || explodedSentMessage[0][0:1] != os.Getenv("BOT_PREFIX") {
		return
	}

	// Gets the parameters that should be passed to the desired command.
	parameters := explodedSentMessage[1:]

	availableCommands[explodedSentMessage[0][1:]](
		session,
		message,
		parameters,
	)
}
