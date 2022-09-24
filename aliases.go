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
		Description: "Returns information about command aliases",
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "alias",
			discordgo.EnglishUS: "alias",
			discordgo.Bulgarian: "алиас",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "Returns information about command aliases",
			discordgo.EnglishUS: "Returns information about command aliases",
			discordgo.Bulgarian: "Връща информация за алиасите на команди",
		},
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "list",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Lists all aliases",
				Options:     nil,
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "list",
					discordgo.EnglishUS: "list",
					discordgo.Bulgarian: "покажи",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "Lists all aliases",
					discordgo.EnglishUS: "Lists all aliases",
					discordgo.Bulgarian: "Показва всички алиаси",
				},
			},
			{
				Name:        "help",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Shows the user how to use command aliases",
				Options:     nil,
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "help",
					discordgo.EnglishUS: "help",
					discordgo.Bulgarian: "помощ",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "Shows the user how to use command aliases",
					discordgo.EnglishUS: "Shows the user how to use command aliases",
					discordgo.Bulgarian: "Показва на потребителя, как да прави алиаси на команди",
				},
			},
			{
				Name:        "refresh",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Refreshes the aliases for the server, might take a couple of minutes to update",
				Options:     nil,
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "refresh",
					discordgo.EnglishUS: "refresh",
					discordgo.Bulgarian: "актуализация",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "Refreshes the aliases for the server, might take a couple of minutes to update",
					discordgo.EnglishUS: "Refreshes the aliases for the server, might take a couple of minutes to update",
					discordgo.Bulgarian: "Актуализира алиасите за сървъра, може да отнеме няколко минути за да завърши актуализацията",
				},
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
	embedStrings := []map[discordgo.Locale]string{
		{
			discordgo.EnglishGB:    "Global commands",
			discordgo.EnglishUS:    "Global commands",
			discordgo.Bulgarian:    "Глобални команди",
			discordgo.ChineseCN:    "Global commands",
			discordgo.ChineseTW:    "Global commands",
			discordgo.Croatian:     "Global commands",
			discordgo.Czech:        "Global commands",
			discordgo.Danish:       "Global commands",
			discordgo.Dutch:        "Global commands",
			discordgo.Finnish:      "Global commands",
			discordgo.French:       "Global commands",
			discordgo.German:       "Global commands",
			discordgo.Greek:        "Global commands",
			discordgo.Hungarian:    "Global commands",
			discordgo.Italian:      "Global commands",
			discordgo.Japanese:     "Global commands",
			discordgo.Korean:       "Global commands",
			discordgo.Lithuanian:   "Global commands",
			discordgo.Norwegian:    "Global commands",
			discordgo.Polish:       "Global commands",
			discordgo.PortugueseBR: "Global commands",
			discordgo.Romanian:     "Global commands",
			discordgo.Russian:      "Global commands",
			discordgo.SpanishES:    "Global commands",
			discordgo.Swedish:      "Global commands",
			discordgo.Ukrainian:    "Global commands",
			discordgo.Hindi:        "Global commands",
			discordgo.Thai:         "Global commands",
			discordgo.Turkish:      "Global commands",
			discordgo.Vietnamese:   "Global commands",
			discordgo.Unknown:      "Global commands",
		},
		{
			discordgo.EnglishGB:    "Default",
			discordgo.EnglishUS:    "Default",
			discordgo.Bulgarian:    "По подразбиране",
			discordgo.ChineseCN:    "Default",
			discordgo.ChineseTW:    "Default",
			discordgo.Croatian:     "Default",
			discordgo.Czech:        "Default",
			discordgo.Danish:       "Default",
			discordgo.Dutch:        "Default",
			discordgo.Finnish:      "Default",
			discordgo.French:       "Default",
			discordgo.German:       "Default",
			discordgo.Greek:        "Default",
			discordgo.Hungarian:    "Default",
			discordgo.Italian:      "Default",
			discordgo.Japanese:     "Default",
			discordgo.Korean:       "Default",
			discordgo.Lithuanian:   "Default",
			discordgo.Norwegian:    "Default",
			discordgo.Polish:       "Default",
			discordgo.PortugueseBR: "Default",
			discordgo.Romanian:     "Default",
			discordgo.Russian:      "Default",
			discordgo.SpanishES:    "Default",
			discordgo.Swedish:      "Default",
			discordgo.Ukrainian:    "Default",
			discordgo.Hindi:        "Default",
			discordgo.Thai:         "Default",
			discordgo.Turkish:      "Default",
			discordgo.Vietnamese:   "Default",
			discordgo.Unknown:      "Default",
		},
		{
			discordgo.EnglishGB:    "List aliases",
			discordgo.EnglishUS:    "List aliases",
			discordgo.Bulgarian:    "Покажи алиасите",
			discordgo.ChineseCN:    "List aliases",
			discordgo.ChineseTW:    "List aliases",
			discordgo.Croatian:     "List aliases",
			discordgo.Czech:        "List aliases",
			discordgo.Danish:       "List aliases",
			discordgo.Dutch:        "List aliases",
			discordgo.Finnish:      "List aliases",
			discordgo.French:       "List aliases",
			discordgo.German:       "List aliases",
			discordgo.Greek:        "List aliases",
			discordgo.Hungarian:    "List aliases",
			discordgo.Italian:      "List aliases",
			discordgo.Japanese:     "List aliases",
			discordgo.Korean:       "List aliases",
			discordgo.Lithuanian:   "List aliases",
			discordgo.Norwegian:    "List aliases",
			discordgo.Polish:       "List aliases",
			discordgo.PortugueseBR: "List aliases",
			discordgo.Romanian:     "List aliases",
			discordgo.Russian:      "List aliases",
			discordgo.SpanishES:    "List aliases",
			discordgo.Swedish:      "List aliases",
			discordgo.Ukrainian:    "List aliases",
			discordgo.Hindi:        "List aliases",
			discordgo.Thai:         "List aliases",
			discordgo.Turkish:      "List aliases",
			discordgo.Vietnamese:   "List aliases",
			discordgo.Unknown:      "List aliases",
		},
	}

	initialField := discordgo.MessageEmbedField{
		Name:   embedStrings[0][m.Locale],
		Value:  embedStrings[1][m.Locale],
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
			Text:    footerTranslations[m.Locale],
		},
		Title: embedStrings[2][m.Locale],
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
	embedStrings := []map[discordgo.Locale]string{
		{
			discordgo.EnglishGB:    "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.EnglishUS:    "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Bulgarian:    "Актуализирахме командите, може да отнеме няколко минути за да се появят промените, моля не спамете командата!",
			discordgo.ChineseCN:    "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.ChineseTW:    "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Croatian:     "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Czech:        "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Danish:       "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Dutch:        "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Finnish:      "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.French:       "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.German:       "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Greek:        "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Hungarian:    "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Italian:      "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Japanese:     "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Korean:       "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Lithuanian:   "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Norwegian:    "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Polish:       "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.PortugueseBR: "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Romanian:     "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Russian:      "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.SpanishES:    "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Swedish:      "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Ukrainian:    "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Hindi:        "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Thai:         "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Turkish:      "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Vietnamese:   "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
			discordgo.Unknown:      "Refreshed, this may take a couple of minutes to take affect, please don't spam the command!",
		},
	}
	_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: embedStrings[0][m.Locale],
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
	embedStrings := []map[discordgo.Locale]string{
		{
			discordgo.EnglishGB:    "Alias help",
			discordgo.EnglishUS:    "Alias help",
			discordgo.Bulgarian:    "Помощ за алиасите",
			discordgo.ChineseCN:    "Alias help",
			discordgo.ChineseTW:    "Alias help",
			discordgo.Croatian:     "Alias help",
			discordgo.Czech:        "Alias help",
			discordgo.Danish:       "Alias help",
			discordgo.Dutch:        "Alias help",
			discordgo.Finnish:      "Alias help",
			discordgo.French:       "Alias help",
			discordgo.German:       "Alias help",
			discordgo.Greek:        "Alias help",
			discordgo.Hungarian:    "Alias help",
			discordgo.Italian:      "Alias help",
			discordgo.Japanese:     "Alias help",
			discordgo.Korean:       "Alias help",
			discordgo.Lithuanian:   "Alias help",
			discordgo.Norwegian:    "Alias help",
			discordgo.Polish:       "Alias help",
			discordgo.PortugueseBR: "Alias help",
			discordgo.Romanian:     "Alias help",
			discordgo.Russian:      "Alias help",
			discordgo.SpanishES:    "Alias help",
			discordgo.Swedish:      "Alias help",
			discordgo.Ukrainian:    "Alias help",
			discordgo.Hindi:        "Alias help",
			discordgo.Thai:         "Alias help",
			discordgo.Turkish:      "Alias help",
			discordgo.Vietnamese:   "Alias help",
			discordgo.Unknown:      "Alias help",
		},
		{
			discordgo.EnglishGB:    "Introduction to aliases",
			discordgo.EnglishUS:    "Introduction to aliases",
			discordgo.Bulgarian:    "Въведение в алиасите",
			discordgo.ChineseCN:    "Introduction to aliases",
			discordgo.ChineseTW:    "Introduction to aliases",
			discordgo.Croatian:     "Introduction to aliases",
			discordgo.Czech:        "Introduction to aliases",
			discordgo.Danish:       "Introduction to aliases",
			discordgo.Dutch:        "Introduction to aliases",
			discordgo.Finnish:      "Introduction to aliases",
			discordgo.French:       "Introduction to aliases",
			discordgo.German:       "Introduction to aliases",
			discordgo.Greek:        "Introduction to aliases",
			discordgo.Hungarian:    "Introduction to aliases",
			discordgo.Italian:      "Introduction to aliases",
			discordgo.Japanese:     "Introduction to aliases",
			discordgo.Korean:       "Introduction to aliases",
			discordgo.Lithuanian:   "Introduction to aliases",
			discordgo.Norwegian:    "Introduction to aliases",
			discordgo.Polish:       "Introduction to aliases",
			discordgo.PortugueseBR: "Introduction to aliases",
			discordgo.Romanian:     "Introduction to aliases",
			discordgo.Russian:      "Introduction to aliases",
			discordgo.SpanishES:    "Introduction to aliases",
			discordgo.Swedish:      "Introduction to aliases",
			discordgo.Ukrainian:    "Introduction to aliases",
			discordgo.Hindi:        "Introduction to aliases",
			discordgo.Thai:         "Introduction to aliases",
			discordgo.Turkish:      "Introduction to aliases",
			discordgo.Vietnamese:   "Introduction to aliases",
			discordgo.Unknown:      "Introduction to aliases",
		},
		{
			discordgo.EnglishGB:    "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.EnglishUS:    "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Bulgarian:    "Алиасите ви позволяват да дадете ново име на съществуваща команда",
			discordgo.ChineseCN:    "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.ChineseTW:    "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Croatian:     "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Czech:        "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Danish:       "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Dutch:        "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Finnish:      "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.French:       "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.German:       "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Greek:        "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Hungarian:    "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Italian:      "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Japanese:     "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Korean:       "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Lithuanian:   "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Norwegian:    "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Polish:       "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.PortugueseBR: "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Romanian:     "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Russian:      "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.SpanishES:    "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Swedish:      "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Ukrainian:    "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Hindi:        "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Thai:         "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Turkish:      "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Vietnamese:   "Aliases are a way of giving a new name to a pre-existing command for your server",
			discordgo.Unknown:      "Aliases are a way of giving a new name to a pre-existing command for your server",
		},
		{
			discordgo.EnglishGB:    "Syntax",
			discordgo.EnglishUS:    "Синтаксис",
			discordgo.Bulgarian:    "Syntax",
			discordgo.ChineseCN:    "Syntax",
			discordgo.ChineseTW:    "Syntax",
			discordgo.Croatian:     "Syntax",
			discordgo.Czech:        "Syntax",
			discordgo.Danish:       "Syntax",
			discordgo.Dutch:        "Syntax",
			discordgo.Finnish:      "Syntax",
			discordgo.French:       "Syntax",
			discordgo.German:       "Syntax",
			discordgo.Greek:        "Syntax",
			discordgo.Hungarian:    "Syntax",
			discordgo.Italian:      "Syntax",
			discordgo.Japanese:     "Syntax",
			discordgo.Korean:       "Syntax",
			discordgo.Lithuanian:   "Syntax",
			discordgo.Norwegian:    "Syntax",
			discordgo.Polish:       "Syntax",
			discordgo.PortugueseBR: "Syntax",
			discordgo.Romanian:     "Syntax",
			discordgo.Russian:      "Syntax",
			discordgo.SpanishES:    "Syntax",
			discordgo.Swedish:      "Syntax",
			discordgo.Ukrainian:    "Syntax",
			discordgo.Hindi:        "Syntax",
			discordgo.Thai:         "Syntax",
			discordgo.Turkish:      "Syntax",
			discordgo.Vietnamese:   "Syntax",
			discordgo.Unknown:      "Syntax",
		},
		{
			discordgo.EnglishGB:    "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.EnglishUS:    "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Bulgarian:    "За да добавите алиас, влезнете в канал, и сложете следният текст в темата на канала \"ubot-macro:\". След символа ':' , следвайте следният синтаксис \"alias-name>command;alias-name>command]\"",
			discordgo.ChineseCN:    "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.ChineseTW:    "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Croatian:     "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Czech:        "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Danish:       "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Dutch:        "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Finnish:      "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.French:       "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.German:       "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Greek:        "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Hungarian:    "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Italian:      "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Japanese:     "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Korean:       "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Lithuanian:   "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Norwegian:    "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Polish:       "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.PortugueseBR: "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Romanian:     "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Russian:      "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.SpanishES:    "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Swedish:      "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Ukrainian:    "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Hindi:        "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Thai:         "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Turkish:      "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Vietnamese:   "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
			discordgo.Unknown:      "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"",
		},
		{
			discordgo.EnglishGB:    "Symbols",
			discordgo.EnglishUS:    "Symbols",
			discordgo.Bulgarian:    "Символи",
			discordgo.ChineseCN:    "Symbols",
			discordgo.ChineseTW:    "Symbols",
			discordgo.Croatian:     "Symbols",
			discordgo.Czech:        "Symbols",
			discordgo.Danish:       "Symbols",
			discordgo.Dutch:        "Symbols",
			discordgo.Finnish:      "Symbols",
			discordgo.French:       "Symbols",
			discordgo.German:       "Symbols",
			discordgo.Greek:        "Symbols",
			discordgo.Hungarian:    "Symbols",
			discordgo.Italian:      "Symbols",
			discordgo.Japanese:     "Symbols",
			discordgo.Korean:       "Symbols",
			discordgo.Lithuanian:   "Symbols",
			discordgo.Norwegian:    "Symbols",
			discordgo.Polish:       "Symbols",
			discordgo.PortugueseBR: "Symbols",
			discordgo.Romanian:     "Symbols",
			discordgo.Russian:      "Symbols",
			discordgo.SpanishES:    "Symbols",
			discordgo.Swedish:      "Symbols",
			discordgo.Ukrainian:    "Symbols",
			discordgo.Hindi:        "Symbols",
			discordgo.Thai:         "Symbols",
			discordgo.Turkish:      "Symbols",
			discordgo.Vietnamese:   "Symbols",
			discordgo.Unknown:      "Symbols",
		},
		{
			discordgo.EnglishGB:    "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.EnglishUS:    "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Bulgarian:    "\"ubot-macro:\", маркира началото на нова поредица от макрота. Символа '>' може да си го представите като стрелка която сочи, от новото име до командата. Символа ';' е разделител между алиаси. Символа ']' спира поредицата макрота",
			discordgo.ChineseCN:    "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.ChineseTW:    "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Croatian:     "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Czech:        "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Danish:       "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Dutch:        "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Finnish:      "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.French:       "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.German:       "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Greek:        "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Hungarian:    "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Italian:      "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Japanese:     "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Korean:       "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Lithuanian:   "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Norwegian:    "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Polish:       "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.PortugueseBR: "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Romanian:     "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Russian:      "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.SpanishES:    "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Swedish:      "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Ukrainian:    "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Hindi:        "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Thai:         "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Turkish:      "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Vietnamese:   "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
			discordgo.Unknown:      "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression",
		},
		{
			discordgo.EnglishGB:    "Example",
			discordgo.EnglishUS:    "Example",
			discordgo.Bulgarian:    "На пример",
			discordgo.ChineseCN:    "Example",
			discordgo.ChineseTW:    "Example",
			discordgo.Croatian:     "Example",
			discordgo.Czech:        "Example",
			discordgo.Danish:       "Example",
			discordgo.Dutch:        "Example",
			discordgo.Finnish:      "Example",
			discordgo.French:       "Example",
			discordgo.German:       "Example",
			discordgo.Greek:        "Example",
			discordgo.Hungarian:    "Example",
			discordgo.Italian:      "Example",
			discordgo.Japanese:     "Example",
			discordgo.Korean:       "Example",
			discordgo.Lithuanian:   "Example",
			discordgo.Norwegian:    "Example",
			discordgo.Polish:       "Example",
			discordgo.PortugueseBR: "Example",
			discordgo.Romanian:     "Example",
			discordgo.Russian:      "Example",
			discordgo.SpanishES:    "Example",
			discordgo.Swedish:      "Example",
			discordgo.Ukrainian:    "Example",
			discordgo.Hindi:        "Example",
			discordgo.Thai:         "Example",
			discordgo.Turkish:      "Example",
			discordgo.Vietnamese:   "Example",
			discordgo.Unknown:      "Example",
		},
	}
	embed := NewEmbed().
		SetTitle(embedStrings[0][m.Locale]).
		AddField(embedStrings[1][m.Locale], embedStrings[2][m.Locale]).
		AddField(embedStrings[3][m.Locale], embedStrings[4][m.Locale]).
		AddField(embedStrings[5][m.Locale], embedStrings[6][m.Locale]).
		InlineAllFields().
		AddField(embedStrings[7][m.Locale], "ubot-macro:set-game-role>set-meta-role;set-city>set-meta-role;set-fav-food-role>set-meta-role]").
		SetFooter(footerTranslations[m.Locale], "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
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
