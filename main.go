package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MadLadSquad/discordgo"
)

var (
	Token string
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

	dg.AddHandler(onMessageCreate)

	dg.AddHandler(onApplicationCommand)

	dg.AddHandler(onReady)

	dg.AddHandler(onChannelCreate)
	dg.AddHandler(onChannelUpdate)
	dg.AddHandler(onChannelRemove)
	dg.AddHandler(onPinUpdate)

	dg.AddHandler(onBanRemove)

	dg.AddHandler(onGuildUpdate)
	dg.AddHandler(onGuildMemberUpdate)
	dg.AddHandler(onWelcome)
	dg.AddHandler(onBye)

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

// Main command handler here for usability reasons
func onApplicationCommand(s *discordgo.Session, m *discordgo.InteractionCreate) {
	if m.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := m.ApplicationCommandData()
	str := ""
	checkVal := data.Name
recheck:
	switch checkVal {
	case "userinfo":
		if data.Options != nil && data.Options[0].Value != nil {
			str = data.Options[0].StringValue()
		}
		showUserInfo(str, s, m)
		break
	case "serverinfo":
		showServerInfo(s, m)
		break
	case "privacy":
		privacyPolicy(s, m)
		break
	case "tos":
		termsOfService(s, m)
		break
	case "about":
		about(s, m)
		break
	case "avatar":
		if data.Options != nil && data.Options[0].Value != nil {
			str = data.Options[0].StringValue()
		}
		avatar(str, s, m)
		break
	case "invite":
		invite(s, m)
		break
	case "verify":
		verify(s, m)
		break
	case "generate-member-role":
		createMemberRole(s, m)
	case "sus":
		sus(s, m)
		break
	case "list-colour-roles":
		listColours(s, m)
		break
	case "set-channel":
		switch data.Options[0].Name {
		case "welcome":
			channelChangeMetadata(data.Options[0].Options[0].StringValue(), s, m, " ubot-welcome", "welcome")
			break
		case "event-tracking":
			channelChangeMetadata(data.Options[0].Options[0].StringValue(), s, m, " ubot-event-log", "event logging")
			break
		case "text-only":
			channelChangeMetadata(data.Options[0].Options[0].StringValue(), s, m, " ubot-restrict-text-only", "text only")
			break
		case "attachments-only":
			channelChangeMetadata(data.Options[0].Options[0].StringValue(), s, m, " ubot-restrict-attachments-only", "attachments only")
			break
		case "links-only":
			channelChangeMetadata(data.Options[0].Options[0].StringValue(), s, m, " ubot-restrict-links-only", "links only")
			break
		case "colour-role":
			channelChangeMetadata(data.Options[0].Options[0].StringValue(), s, m, " ubot-colour-pick", "colour role")
			break
		case "meta-role":
			channelChangeMetadata(data.Options[0].Options[0].StringValue(), s, m, " ubot-meta-role-pick", "meta role")
			break
		}
		break
	case "meta-role":
		switch data.Options[0].Name {
		case "give":
			giveMetarole(data.Options[0].Options[0].StringValue(), s, m)
			break
		case "remove":
			removeMetarole(data.Options[0].Options[0].StringValue(), s, m)
			break
		}
		break
	case "set-colour-role":
		giveColour(data.Options[0].StringValue(), s, m)
		break
	case "alias":
		switch data.Options[0].Name {
		case "list":
			listAliases(s, m)
			break
		case "help":
			aliasHelp(s, m)
			break
		case "refresh":
			refresh(s, m)
			break
		}
		break
	default:
		// The default behaviour after everything is done is to check if the command is a macro.
		// Here we iterate all commands, check if the name is equal to one of the guild only commands,
		// and if it is we change the check value to the value and recheck by doing a quick goto call
		cmds, _ := s.ApplicationCommands(s.State.User.ID, m.GuildID)
		for i := 0; i < len(cmds); i++ {
			if data.Name == cmds[i].Name {
				// Remove the first part from the description and the last symbol, assign checkVal and recheck
				checkVal = cmds[i].Description[len("Custom command that calls is an alias for "):]
				checkVal = checkVal[:len(checkVal)-1]
				goto recheck
			}
		}
	}

}
