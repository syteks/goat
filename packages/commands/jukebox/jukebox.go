package jukebox

import (
	"encoding/binary"
	"github.com/bwmarrin/discordgo"
	"github.com/kkdai/youtube/v2"
	"io"
)

func Play(session *discordgo.Session, message *discordgo.MessageCreate, parameters []string) {
	// Gets the user's voice state (the voice channel that he is in the given guild)
	voiceState, err := session.State.VoiceState(message.GuildID, message.Author.ID)

	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "You must be in a voice channel.")
		return
	}

	client := youtube.Client{}

	video, _ := client.GetVideo("https://www.youtube.com/watch?v=QohH89Eu5iM&ab_channel=AEdicas-Tutoriaisemportugu%C3%AAs")

	format := video.Formats[0]

	reader, _, _ := client.GetStream(video, &format)

	voiceConnection, err := session.ChannelVoiceJoin(message.GuildID, voiceState.ChannelID, false, true)
	voiceConnection.Speaking(true)

	defer voiceConnection.Speaking(false)

	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "There was an error joinning the channel.")
		return
	}

	buffer := make([]byte, 960*2)

	for {
		err = binary.Read(reader, binary.LittleEndian, &buffer)

		if err == io.EOF {
			break
		}
	}
	voiceConnection.OpusSend <- buffer
}
