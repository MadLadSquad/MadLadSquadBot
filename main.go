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
	prefix = "!"
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
	dg.AddHandler(welcome)

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

	if strings.Contains(content, prefix+"userinfo") {
		showUserInfo(message[1], s, m)
	} else if strings.Contains(content, prefix+"serverinfo") {
		showServerInfo(s, m)
	} else if strings.Contains(content, prefix+"help") {
		help(s, m)
	} else if strings.Contains(content, prefix+"kick") {
		kick(message[1], s, m)
	} else if strings.Contains(content, prefix+"ban") {
		arg := [2]string{ message[1] }

		for i := 2; i < 102; i++ {
			if message[i] != "" {
				arg[1] += message[i] + " "
			} else {
				break
			}
		}

		ban(arg, s, m)
	} else if strings.Contains(content, prefix+"invite") {
		invite(s, m)
	} else if strings.Contains(content, prefix+"privacy") {
		privacyPolicy(s, m)
	} else if strings.Contains(content, prefix+"tos") {
		termsOfService(s, m)
	} else if strings.Contains(content, prefix+"about") {
		about(s, m)
	} else if strings.Contains(content, prefix+"verify") {
		verify(s, m)
	} else if strings.Contains(content, prefix+"avatar") {
		avatar(message[1], s, m)
	} else if strings.Contains(content, prefix+"mute") {
		mute(message[1], s, m)
	} else if strings.Contains(content, prefix+"sus") {
		sus(s, m)
	}
}