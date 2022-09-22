// This file provides functions to create the commands for metadata change functions
package main

import (
	"fmt"
	"github.com/MadLadSquad/discordgo"
)

// Should be const but the functions need the address
var (
	dmperm       = false
	perm   int64 = 0
)

func createSetChannel(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:                     "set-channel",
		Type:                     discordgo.ChatApplicationCommand,
		Description:              "Sets up a channel to have certain permissions",
		DMPermission:             &dmperm,
		DefaultMemberPermissions: &perm,
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "set-channel",
			discordgo.EnglishUS: "set-channel",
			discordgo.Bulgarian: "променяне-на-канал",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "Sets up a channel to have certain permissions",
			discordgo.EnglishUS: "Sets up a channel to have certain permissions",
			discordgo.Bulgarian: "Променя даден канал да има допълнителни права/правила",
		},
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "welcome",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Sets the welcome channel",
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "welcome",
					discordgo.EnglishUS: "welcome",
					discordgo.Bulgarian: "добре-дошли",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "Sets the welcome channel",
					discordgo.EnglishUS: "Sets the welcome channel",
					discordgo.Bulgarian: "Прави даден канал от тип \"добре дошли\"",
				},
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "channel-mention",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The channel to be set",
						Required:    true,
						NameLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "channel-metion",
							discordgo.EnglishUS: "channel-metion",
							discordgo.Bulgarian: "пинг-към-канала",
						},
						DescriptionLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "The channel to be set",
							discordgo.EnglishUS: "The channel to be set",
							discordgo.Bulgarian: "Канала, който трбва да бъде променен",
						},
					},
				},
			},
			{
				Name:        "event-tracking",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Sets the event tracking channel",
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "event-tracking",
					discordgo.EnglishUS: "event-tracking",
					discordgo.Bulgarian: "проследяване-на-промени",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "Sets the event tracking channel",
					discordgo.EnglishUS: "Sets the event tracking channel",
					discordgo.Bulgarian: "Прави даден канал да проследява промени в състоянието на групата",
				},
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "channel-mention",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The channel to be set",
						Required:    true,
						NameLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "channel-metion",
							discordgo.EnglishUS: "channel-metion",
							discordgo.Bulgarian: "пинг-към-канала",
						},
						DescriptionLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "The channel to be set",
							discordgo.EnglishUS: "The channel to be set",
							discordgo.Bulgarian: "Канала, който трбва да бъде променен",
						},
					},
				},
			},
			{
				Name:        "text-only",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Sets the channel as text only",
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "text-only",
					discordgo.EnglishUS: "text-only",
					discordgo.Bulgarian: "само-текст",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "Sets the channel as text only",
					discordgo.EnglishUS: "Sets the channel as text only",
					discordgo.Bulgarian: "Променя даден канал, така че, единствените позволен съобщения са тези, съдържащи само нормален текст",
				},
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "channel-mention",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The channel to be set",
						Required:    true,
						NameLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "channel-metion",
							discordgo.EnglishUS: "channel-metion",
							discordgo.Bulgarian: "пинг-към-канала",
						},
						DescriptionLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "The channel to be set",
							discordgo.EnglishUS: "The channel to be set",
							discordgo.Bulgarian: "Канала, който трбва да бъде променен",
						},
					},
				},
			},
			{
				Name:        "attachments-only",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Sets the channel as attachments only",
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "attachments-only",
					discordgo.EnglishUS: "attachments-only",
					discordgo.Bulgarian: "само-прикачени-файлове",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "Sets the channel as attachments only",
					discordgo.EnglishUS: "Sets the channel as attachments only",
					discordgo.Bulgarian: "Променя даден канал, така че, единствените позволени съобщения са тези, съдъджащи прикачени файлове",
				},
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "channel-mention",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The channel to be set",
						Required:    true,
						NameLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "channel-metion",
							discordgo.EnglishUS: "channel-metion",
							discordgo.Bulgarian: "пинг-към-канала",
						},
						DescriptionLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "The channel to be set",
							discordgo.EnglishUS: "The channel to be set",
							discordgo.Bulgarian: "Канала, който трбва да бъде променен",
						},
					},
				},
			},
			{
				Name:        "links-only",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Sets the channel as links only",
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "links-only",
					discordgo.EnglishUS: "links-only",
					discordgo.Bulgarian: "само-линкове",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "Sets the channel as links only",
					discordgo.EnglishUS: "Sets the channel as links only",
					discordgo.Bulgarian: "Променя даден канал, така че, единствените позволени съобщения са тези, съдържащи линкове",
				},
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "channel-mention",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The channel to be set",
						Required:    true,
						NameLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "channel-metion",
							discordgo.EnglishUS: "channel-metion",
							discordgo.Bulgarian: "пинг-към-канала",
						},
						DescriptionLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "The channel to be set",
							discordgo.EnglishUS: "The channel to be set",
							discordgo.Bulgarian: "Канала, който трбва да бъде променен",
						},
					},
				},
			},
			{
				Name:        "colour-role",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Sets the colour role channel",
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "colour-role",
					discordgo.EnglishUS: "color-role",
					discordgo.Bulgarian: "цветови-роли",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "Sets the colour role channel",
					discordgo.EnglishUS: "Sets the colour role channel",
					discordgo.Bulgarian: "Прави даден канал, канал за слагане на цветови роли",
				},
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "channel-mention",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The channel to be set",
						Required:    true,
						NameLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "channel-metion",
							discordgo.EnglishUS: "channel-metion",
							discordgo.Bulgarian: "пинг-към-канала",
						},
						DescriptionLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "The channel to be set",
							discordgo.EnglishUS: "The channel to be set",
							discordgo.Bulgarian: "Канала, който трбва да бъде променен",
						},
					},
				},
			},
			{
				Name:        "meta-role",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Sets the meta role channel",
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "meta-role",
					discordgo.EnglishUS: "meta-role",
					discordgo.Bulgarian: "мета-роля",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "Sets the meta role channel",
					discordgo.EnglishUS: "Sets the meta role channel",
					discordgo.Bulgarian: "Прави даден канал, канал за слагане на мета роли",
				},
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "channel-mention",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The channel to be set",
						NameLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "channel-metion",
							discordgo.EnglishUS: "channel-metion",
							discordgo.Bulgarian: "пинг-към-канала",
						},
						DescriptionLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "The channel to be set",
							discordgo.EnglishUS: "The channel to be set",
							discordgo.Bulgarian: "Канала, който трбва да бъде променен",
						},
						Required: true,
					},
				},
			},
		},
	}
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
	if err != nil {
		fmt.Println(err)
	}
}
