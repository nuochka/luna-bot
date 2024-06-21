package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the token from the environment variable
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatalf("No token provided. Please set DISCORD_BOT_TOKEN in the .env file.")
	}

	// Create a new Discord session
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	// Add message handler
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		fmt.Printf("Message received: %+v\n", m)

		if m.Content == "hello" {
			_, err := s.ChannelMessageSend(m.ChannelID, "world!")
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// Open the session
	err = sess.Open()
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}

	defer sess.Close()
	fmt.Println("the bot is online")

	// Wait for a termination signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
