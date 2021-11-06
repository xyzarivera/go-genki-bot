package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(helloWorld)
	dg.AddHandler(standupSession)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func helloWorld(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!helloWorld" {
		s.ChannelMessageSend(m.ChannelID, "Got this! See you later "+m.Author.Username+" :grin:")
	}

}

func standupSession(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// start command
	// sends standup questions
	if m.Content == "!standup" {

		s.ChannelMessageSend(m.ChannelID, `Start of standup 
		What did you do?
		What will you do?
		Do you have any blockers?`)
	}

	// if chatbot is replied
	// needs to check if it replied initial reply
	// needs error checking, ignored for now
	if m.Type == 19 {
		st, _ := s.ChannelMessage(m.ChannelID, m.MessageReference.MessageID)

		if st.Author.ID == s.State.User.ID {
			s.ChannelMessageSend(m.ChannelID, "Got this! See you later "+m.Author.Username+" :grin:")
		}
	}
}
