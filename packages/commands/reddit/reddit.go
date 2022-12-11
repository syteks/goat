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

	channel, err := session.Channel(message.ChannelID)

	custom_error.Handle(err, "There was an error fetching the Channel information.")

	// If there was an error fetching the channel it means
	if err != nil {
		return
	}

	// Gets random memes from the subreddit '/r/memes'.
	posts, err := GetSubredditPosts(subredditName+"/rising", subredditType, channel.NSFW)

	// If there is an error fetching the posts from the subreddit, send a message to the user.
	if err != nil {
		_, _ = session.ChannelMessageSend(message.ChannelID, "Could not find the subreddit")
		return
	}

	// If no posts is found in the subreddit.
	if len(posts) == 0 {
		_, _ = session.ChannelMessageSend(message.ChannelID, "No posts found respecting the channel's age restriction and the given subreddit name.")
		return
	}

	randNumber := 0

	if len(posts) > 1 {
		// Get a random post from the fetched posts.
		randNumber = rand.Intn(len(posts)-1) + 1
	}

	post := posts[randNumber]

	if post.IsVideo {
		_, err := session.ChannelMessageSend(message.ChannelID, post.Mp4.Source.Url)

		custom_error.Handle(err, "Failed to send the reddit meme video.")
	} else {
		_, err := session.ChannelMessageSend(message.ChannelID, post.Url)

		custom_error.Handle(err, "Failed to send the reddit meme images.")
	}
}
