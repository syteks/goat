package reddit

import (
	"github.com/bwmarrin/discordgo"
	"goat/packages/handlers/custom_error"
	"math/rand"
)

func GetMeme(session *discordgo.Session, message *discordgo.MessageCreate, parameters []string) {
	subredditName := "memes"
	subredditType := "r"

	for index, value := range parameters {
		if index == 0 {
			subredditName = value
		}

		if index == 1 {
			subredditType = value
		}
	}

	// Gets random memes from the subreddit '/r/memes'.
	posts, err := GetSubredditPosts(subredditName+"/rising", subredditType)

	if err != nil {
		session.ChannelMessageSend("Could not find the subreddit", message.ChannelID)
		return
	}

	if len(posts) <= 0 {
		return
	}

	randNumber := rand.Intn(len(posts)-1) + 1

	post := posts[randNumber]

	if post.IsVideo {
		_, err := session.ChannelMessageSend(message.ChannelID, post.Mp4.Source.Url)

		custom_error.Handle(err, "Failed to send the reddit meme video.")
	} else {
		_, err := session.ChannelMessageSend(message.ChannelID, post.Url)

		custom_error.Handle(err, "Failed to send the reddit meme images.")
	}
}
