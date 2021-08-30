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
	}

	if strings.Contains(content, prefix+"serverinfo") {
		showServerInfo(s, m)
	}

	if strings.Contains(content, prefix+"help") {
		help(s, m)
	}

	if strings.Contains(content, prefix+"kick") {
		kick(message[1], s, m)
	}

	if strings.Contains(content, prefix+"ban") {
		arg := [3]string{ message[1], message[2], message[3] }
		ban(arg, s, m)
	}
}

func parseMessage(str string) [50]string {
	ret := [50]string{}

	index := 0

	for i := 0; i < len(str); i++ {
		if index < 50 {
			if str[i] == ' ' {
				index++
			} else {
				ret[index] += string(str[i])
			}
		} else {
			break
		}
	}

	return ret
}