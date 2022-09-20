package main

import (
	"fmt"
	"github.com/MadLadSquad/discordgo"
	"strings"
)

func createAlias(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "alias",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Returns information about the caller or the mentioned user",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "list",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Lists all aliases",
				Options:     nil,
			},
			{
				Name:        "help",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Shows the user how to use command aliases",
				Options:     nil,
			},
			{
				Name:        "refresh",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Refreshes the aliases for the server, might take a couple of minutes to update",
				Options:     nil,
			},
		},
	}
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
	if err != nil {
		fmt.Println(err)
	}
}

func parseMacro(fields *[]*discordgo.MessageEmbedField, channel *discordgo.Channel) {
	topic := strings.ToLower(channel.Topic)

	if strings.Contains(topic, "ubot-macro:") {
		i := strings.Index(topic, "ubot-macro:")

		if i != -1 {
			i += len("ubot-macro:")
		} else {
			return
		}
		bAccumulateAlias := true
		currentAlias := ""
		currentCmd := ""

		for j := i; j < len(topic); j++ {
			if bAccumulateAlias {
				if topic[j] == '>' {
					bAccumulateAlias = false
				} else {
					currentAlias += string(topic[j])
				}
			} else {
				if topic[j] == ';' || topic[j] == ' ' || topic[j] == '[' || topic[j] == ']' {
					bAccumulateAlias = true

					*fields = append(*fields, &discordgo.MessageEmbedField{
						Name:   currentAlias,
						Value:  currentCmd,
						Inline: true,
					})

					currentCmd = ""
					currentAlias = ""
				} else {
					currentCmd += string(topic[j])
				}
			}
		}
	}
}

func listAliases(s *discordgo.Session, m *discordgo.InteractionCreate) {
	initialField := discordgo.MessageEmbedField{
		Name:   "Global commands",
		Value:  "Default",
		Inline: true,
	}
	fields := []*discordgo.MessageEmbedField{&initialField}
	channels, _ := s.GuildChannels(m.GuildID)
	for i := 0; i < len(channels); i++ {
		parseMacro(&fields, channels[i])
	}

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0xf1c40f,
		Fields: fields,
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4",
			Text:    "Message delivered using Untitled Technology",
		},
		Title: "Aliases List",
	}
	_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: "",
			Flags:   discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				embed,
			},
		},
	})
}

func refresh(s *discordgo.Session, m *discordgo.InteractionCreate) {
	_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	// Delete the previous commands
	cmds, _ := s.ApplicationCommands(s.State.User.ID, m.GuildID)
	for i := 0; i < len(cmds); i++ {
		_ = s.ApplicationCommandDelete(s.State.User.ID, m.GuildID, cmds[i].ID)
	}

	// Find the new commands and add them
	var fields []*discordgo.MessageEmbedField
	channels, _ := s.GuildChannels(m.GuildID)
	for i := 0; i < len(channels); i++ {
		parseMacro(&fields, channels[i])
	}
	for i := 0; i < len(fields); i++ {
		command := &discordgo.ApplicationCommand{
			Name:        fields[i].Name,
			Type:        discordgo.ChatApplicationCommand,
			Description: "Custom command that calls is an alias for " + fields[i].Value + "!",
		}
		_, _ = s.ApplicationCommandCreate(s.State.User.ID, m.GuildID, command)
	}
}

func aliasHelp(s *discordgo.Session, m *discordgo.InteractionCreate) {
	embed := NewEmbed().
		SetTitle("Alias help").
		AddField("Introduction to aliases", "Aliases are a way of giving a new name to a pre-existing command for your server").
		AddField("Syntax", "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"").
		AddField("Symbols", "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression").
		InlineAllFields().
		AddField("Example", "ubot-macro:set-game-role>set-meta-role;set-city>set-meta-role;set-fav-food-role>set-meta-role]").
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: "",
			Flags:   discordgo.MessageFlagsEphemeral,
			Embeds:  []*discordgo.MessageEmbed{embed},
		},
	})
}
