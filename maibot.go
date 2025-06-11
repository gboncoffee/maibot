package main

import (
	"log"
	"os"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func main() {
	channelID := flag.String("channelID", "", "Channel ID to send message.")
	message := flag.String("message", "", "The message to send.")

	flag.Parse()

	if *channelID == "" {
		flag.Usage()
		os.Exit(1)
	}

	discord, err := Login()
	if err != nil {
		log.Fatalf("Cannot login to Discord: %v", err)
	}

	_, err = discord.ChannelMessageSend(
		*channelID,
		*message,
		func(_ *discordgo.RequestConfig) {},
	)

	if err != nil {
		log.Fatalf("Cannot send Discord message: %v", err)
	}
}

func Login() (*discordgo.Session, error) {
	token := os.Getenv("DISCORD_TOKEN")
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("could not connect to Discord: %v", err)
	}

	return discord, nil
}
