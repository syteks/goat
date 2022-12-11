package reddit

import (
	"github.com/bwmarrin/discordgo"
	"goat/packages/handlers/custom_error"
	"math/rand"
)

// GetSubreddit gets a subreddit post and displays it.
func GetSubreddit(session *discordgo.Session, message *discordgo.MessageCreate, parameters []string) {
	// Default subreddit name.
	subredditName := "memes"

	// Default subreddit type /r.
	subredditType := "r"

	// Go through each one of the parameters and get the subreddit name and type.
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

	// If there is an error fetching the posts from the subreddit, send a message to the user.
	if err != nil {
		_, _ = session.ChannelMessageSend("Could not find the subreddit", message.ChannelID)
		return
	}

	// If no posts is found in the subreddit.
	if len(posts) <= 0 {
		return
	}

	// Get a random post from the fetched posts.
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
