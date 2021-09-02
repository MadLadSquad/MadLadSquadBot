package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/MadLadSquad/discordgo"
)

var (
	Token  string
	prefix = "&"
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.AddHandler(onReady)

	dg.AddHandler(onChannelCreate)
	dg.AddHandler(onChannelUpdate)
	dg.AddHandler(onChannelRemove)
	dg.AddHandler(onPinUpdate)

	dg.AddHandler(onBanRemove)

	dg.AddHandler(onGuildUpdate)
	dg.AddHandler(onGuildMemberUpdate)
	dg.AddHandler(welcome)
	dg.AddHandler(bye)

	dg.AddHandler(onRoleCreate)
	dg.AddHandler(onRoleUpdate)
	dg.AddHandler(onRoleRemove)

	if err != nil {
		return
	}

	dg.Identify.Intents = discordgo.IntentsAll

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Starting the avanta!")
	time.Sleep(2 * time.Second)
	fmt.Println("Configuring the avanta process!")
	time.Sleep(3 * time.Second)
	fmt.Println("Recruiting the avantajii!")
	time.Sleep(5 * time.Second)
	fmt.Println("We are ready to avantim people!")
	fmt.Println("Press CTRL+C to exit!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	err2 := dg.Close()
	if err2 != nil {
		return
	}
}

// This is the main event meaning it won't go into events.go
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.ToLower(m.Content)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Author.Bot {
		return
	}

	message := parseMessage(content)

	channel, _ := s.State.Channel(m.ChannelID)
	if strings.Contains(strings.ToLower(channel.Topic), "ubot-restrict-text-only") && (len(m.Attachments) > 0 || strings.Contains(content, "http://" ) || strings.Contains(content, "https://")) {
		_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Message deleted due to it containing links or attachments in a text only channel!")
	} else if strings.Contains(strings.ToLower(channel.Topic), "ubot-restrict-attachments-only") && len(m.Attachments) == 0 {
		_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Message deleted due to it containing no attachments in an attachment only channel!")
	} else if strings.Contains(strings.ToLower(channel.Topic), "ubot-restrict-links-only") && !(strings.Contains(content, "http://" ) || strings.Contains(content, "https://")) {
		_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Message deleted due to it containing no links in link only channel!")
	} else {
		onMessageCreate(s, m, message, content)
	}
}