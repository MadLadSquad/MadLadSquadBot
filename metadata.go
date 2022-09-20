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
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "welcome",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Sets the welcome channel",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "channel-mention",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The channel to be set",
						Required:    true,
					},
				},
			},
			{
				Name:        "event-tracking",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Sets the event tracking channel",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "channel-mention",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The channel to be set",
						Required:    true,
					},
				},
			},
			{
				Name:        "text-only",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Sets the channel as text only",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "channel-mention",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The channel to be set",
						Required:    true,
					},
				},
			},
			{
				Name:        "attachments-only",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Sets the channel as attachments only",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "channel-mention",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The channel to be set",
						Required:    true,
					},
				},
			},
			{
				Name:        "links-only",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Sets the channel as links only",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "channel-mention",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The channel to be set",
						Required:    true,
					},
				},
			},
			{
				Name:        "colour-role",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Sets the colour role channel",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "channel-mention",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The channel to be set",
						Required:    true,
					},
				},
			},
			{
				Name:        "meta-role",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Sets the meta role channel",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "channel-mention",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The channel to be set",
						Required:    true,
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
