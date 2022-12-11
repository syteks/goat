package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"goat/packages/handlers/command"
	"goat/packages/handlers/custom_error"
	"log"
	"os"
)

// main is the start of our discord bot life.
// Loads the .env file and starts the discord heartbeat.
func main() {
	loadEnvFile()
	startDiscord()
}

// startDiscord will start the discord bot heartbeat.
// It will initiate the discord bot thread.
func startDiscord() {
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	custom_error.Handle(err, "Could not find the bots ID.")

	discord.AddHandler(discordReadyHandler)
	discord.AddHandler(command.Handle)

	// Start the discord heartbeat.
	err = discord.Open()

	custom_error.Handle(err, "Could not open the discord bot heart beat.")

	defer func(discord *discordgo.Session) {
		err := discord.Close()

		if err != nil {
			custom_error.Handle(err, "Could not stop the discord bot heart beat.")
		}
	}(discord)

	<-make(chan struct{})
}

// discordReadyHandler will be triggered when the bot is ready to be started.
// Changes the bot status is order
func discordReadyHandler(discord *discordgo.Session, ready *discordgo.Ready) {
	err := discord.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status: "Abuse me with -",
		AFK:    false,
	})

	custom_error.Handle(err, "Could not set the bot's status.")
}

// loadEnvFile loads the bot's env file.
func loadEnvFile() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Could not laod the .env file. Err: %s", err)
	}
}
