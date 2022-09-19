package main

import (
	"fmt"
	"github.com/MadLadSquad/discordgo"
	"strconv"
	"strings"
)

var (
	coloursHex = []uint64{
		0xf0f8ff,
		0xfaebd7,
		0x00ffff,
		0x7fffd4,
		0xf0ffff,
		0xf5f5dc,
		0xffe4c4,
		0x000000,
		0xffebcd,
		0x0000ff,
		0x8a2be2,
		0xa52a2a,
		0xdeb887,
		0x5f9ea0,
		0x7fff00,
		0xd2691e,
		0xff7f50,
		0x6495ed,
		0xfff8dc,
		0xdc143c,
		0x00ffff,
		0x00008b,
		0x008b8b,
		0xb8860b,
		0xa9a9a9,
		0xbdb76b,
		0x8b008b,
		0x556b2f,
		0xff8c00,
		0x9932cc,
		0x8b0000,
		0xe9967a,
		0x8fbc8f,
		0x483d8b,
		0x2f4f4f,
		0x00ced1,
		0x9400d3,
		0xff1493,
		0x00bfff,
		0x696969,
		0x1e90ff,
		0xb22222,
		0xfffaf0,
		0x228b22,
		0xff00ff,
		0xdcdcdc,
		0xf8f8ff,
		0xffd700,
		0xdaa520,
		0x808080,
		0x008000,
		0xadff2f,
		0xf0fff0,
		0xff69b4,
		0xcd5c5c,
		0x4b0082,
		0xfffff0,
		0xf0e68c,
		0xe6e6fa,
		0xfff0f5,
		0x7cfc00,
		0xfffacd,
		0xadd8e6,
		0xf08080,
		0xe0ffff,
		0xfafad2,
		0xd3d3d3,
		0x90ee90,
		0xffb6c1,
		0xffa07a,
		0x20b2aa,
		0x87cefa,
		0x778899,
		0xb0c4de,
		0xffffe0,
		0x00ff00,
		0x32cd32,
		0xfaf0e6,
		0xff00ff,
		0x7f0000,
		0x66cdaa,
		0x0000cd,
		0xba55d3,
		0x9370db,
		0x3cb371,
		0x7b68ee,
		0x00fa9a,
		0x48d1cc,
		0xc71585,
		0x191970,
		0xf5fffa,
		0xffe4e1,
		0xffe4b5,
		0xffdead,
		0x000080,
		0xfdf5e6,
		0x808000,
		0x6b8e23,
		0xffa500,
		0xff4500,
		0xda70d6,
		0xeee8aa,
		0x98fb98,
		0xafeeee,
		0xdb7093,
		0xffefd5,
		0xffdab9,
		0xcd853f,
		0xffc0cb,
		0xdda0dd,
		0xb0e0e6,
		0x7f007f,
		0xff0000,
		0xbc8f8f,
		0x4169e1,
		0x8b4513,
		0xfa8072,
		0xf4a460,
		0x2e8b57,
		0xfff5ee,
		0xa0522d,
		0xc0c0c0,
		0x87ceeb,
		0x6a5acd,
		0x708090,
		0xfffafa,
		0x00ff7f,
		0x4682b4,
		0xd2b48c,
		0x008080,
		0xd8bfd8,
		0xff6347,
		0x40e0d0,
		0xee82ee,
		0xf5deb3,
		0xffffff,
		0xf5f5f5,
		0xffff00,
		0x9acd32,
	}

	colours = []string{
		"Alice-Blue",
		"Antique-White",
		"Aqua",
		"Aquamarine",
		"Azure-Mist",
		"Beige",
		"Bisque",
		"Black",
		"Blanched-Almond",
		"Blue",
		"Blue-violet",
		"Auburn",
		"Burlywood",
		"Cadet-Blue",
		"Chartreuse",
		"Cinnamon",
		"Coral",
		"Cornflower-Blue",
		"Cornsilk",
		"Crimson",
		"Cyan",
		"Dark-Blue",
		"Dark-Cyan",
		"Dark-Goldenrod",
		"Dark-Gray",
		"Dark-Khaki",
		"Dark-Magenta",
		"Dark-Olive-Green",
		"Dark-Orange",
		"Dark-Orchid",
		"Dark-Red",
		"Dark-Salmon",
		"Dark-Sea-Green",
		"Dark-Slate-Blue",
		"Dark-Slate-Gray",
		"Dark-Turquoise",
		"Dark-Violet",
		"Deep-Pink",
		"Capri",
		"Dim-Gray",
		"Dodger-Blue",
		"Firebrick",
		"Floral-White",
		"Forest-Green",
		"Magenta",
		"Gainsboro",
		"Ghost-White",
		"Gold",
		"Goldenrod",
		"Gray",
		"Office-Green",
		"Green-yellow",
		"Honeydew",
		"Hot-Pink",
		"Chestnut",
		"Indigo",
		"Ivory",
		"Khaki",
		"Lavender",
		"Lavender-Blush",
		"Lawn-Green",
		"Lemon-Chiffon",
		"Light-Blue",
		"Light-Coral",
		"Light-Cyan",
		"Light-Goldenrod-Yellow",
		"Light-Gray",
		"Light-Green",
		"Light-Pink",
		"Light-Salmon",
		"Light-Sea-Green",
		"Light-Sky-Blue",
		"Light-Slate-Gray",
		"Light-Steel-Blue",
		"Light-Yellow",
		"Lime",
		"Lime-Green",
		"Linen",
		"Fuchsia",
		"Wood-Red",
		"Tropical-Green",
		"Medium-Blue",
		"Medium-Orchid",
		"Medium-Purple",
		"Medium-Sea-Green",
		"Medium-Slate-Blue",
		"Medium-Spring-Green",
		"Medium-Turquoise",
		"Red-violet",
		"Midnight-Blue",
		"Mint-Cream",
		"Misty-Rose",
		"SandySahara",
		"Navajo-White",
		"Navy-Blue",
		"Old-Lace",
		"Olive",
		"Olive-Drab",
		"Orange",
		"Orange-red",
		"Orchid",
		"Pale-Goldenrod",
		"Pale-Green",
		"Pale-Blue",
		"Pale-Violet-red",
		"Papaya-Whip",
		"Peach-Puff",
		"Peru",
		"Pink",
		"Plum",
		"Powder-Blue",
		"Sienna",
		"Red",
		"Rosy-Brown",
		"Royal-Blue",
		"Saddle-Brown",
		"Salmon",
		"Sandy-Brown",
		"Sea-Green",
		"Seashell",
		"Dark-Sand",
		"Silver",
		"Sky-Blue",
		"Slate-Blue",
		"Slate-Gray",
		"Snow",
		"Spring Green",
		"Steel-Blue",
		"Tan",
		"Teal",
		"Thistle",
		"Tomato",
		"Light-Cyan-2",
		"Violet",
		"Wheat",
		"White",
		"White-Smoke",
		"Electric-Yellow",
		"Yellow-green",
	}
)

func createGiveColour(s *discordgo.Session) {
	var choices []*discordgo.ApplicationCommandOptionChoice

	for i := 0; i < len(colours); i++ {
		_ = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  strings.ToLower(colours[i]),
			Value: strconv.Itoa(i),
		})
	}

	command := &discordgo.ApplicationCommand{
		Name:        "set-colour-role",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Gives the user a colour role",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "colour",
				Description: "The name of your colour, run the /list-colour-roles command for more info",
				Type:        discordgo.ApplicationCommandOptionString,
				Choices:     choices,
			},
		},
	}
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
	if err != nil {
		fmt.Println(err)
	}
}

func createListColourRoles(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "list-colour-roles",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Gives info on the colour roles",
	}
	_, _ = s.ApplicationCommandCreate(s.State.User.ID, "", command)
}

func createListAliases(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "list-aliases",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Lists the aliases in the given channel",
	}
	_, _ = s.ApplicationCommandCreate(s.State.User.ID, "", command)
}

func createMetaRole(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "meta-role",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Manages meta roles",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "give",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Creates and/or gives the user a certain meta role",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "role-name",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The name of the role(case-sensitive)",
						Required:    true,
					},
				},
			},
			{
				Name:        "remove",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Removes a meta role",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "role-name",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The name of the role(case-sensitive)",
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

func listColours(s *discordgo.Session, m *discordgo.InteractionCreate) {
	embed := NewEmbed().
		SetTitle("Available colour names").
		AddField("A colour/name table can be found here", "https://www.spycolor.com/w3c-colors").
		AddField("How to use", "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it").
		InlineAllFields().
		AddField("The following colours don't have a name so we made up a name for them", "#40e0d0, #a0522d, #fa8072, #7f007f, #ffe4b5, #66cdaa").
		AddField("#40e0d0", "Light-Cyan-2").
		AddField("#a0522d", "Dark-Sand").
		AddField("#fa8072", "Salmon").
		AddField("#7f007f", "Sienna").
		AddField("#ffe4b5", "SandySahara").
		AddField("#66cdaa", "Tropical-Green").
		InlineAllFields().
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").SetColor(0xf1c40f).MessageEmbed

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

func giveColour(arg string, s *discordgo.Session, m *discordgo.InteractionCreate) {
	channel, _ := s.Channel(m.ChannelID)
	if strings.Contains(strings.ToLower(channel.Topic), "ubot-colour-pick") {
		bFound := false
		colour := 0
		var name = ""

		for i := 0; i < len(colours); i++ {
			if strings.ToLower(colours[i]) == strings.ToLower(arg) {
				bFound = true
				colour = int(coloursHex[i])
				name = colours[i]
				break
			}
		}

		if bFound {
			roles, _ := s.GuildRoles(m.GuildID)
			for i := 0; i < len(roles); i++ {
				if name == roles[i].Name {
					_ = s.GuildMemberRoleAdd(m.GuildID, m.Member.User.ID, roles[i].ID)
					embed := NewEmbed().
						SetTitle("Added you to the "+name+" role!").
						AddField("Role", roles[i].Mention()).
						SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
						SetColor(0xf1c40f).MessageEmbed
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
					return
				}
			}

			perms := int64(0)

			roledata := discordgo.RoleParams{}
			roledata.Permissions = &perms
			roledata.Color = &colour
			roledata.Name = name

			role, _ := s.GuildRoleCreate(m.GuildID, &roledata)
			_ = s.GuildMemberRoleAdd(m.GuildID, m.Member.User.ID, role.ID)
			embed := NewEmbed().
				SetTitle("Added you to the "+name+" role!").
				AddField("Role", role.Mention()).
				SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
				SetColor(0xf1c40f).MessageEmbed
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
		} else {
			_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					TTS:     false,
					Content: "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
		}
	} else {
		_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				TTS:     false,
				Content: "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-colour-role-channel\" command in order to set up colour roles!",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
}

func giveMetarole(arg string, s *discordgo.Session, m *discordgo.InteractionCreate) {
	channel, _ := s.Channel(m.ChannelID)

	if strings.Contains(strings.ToLower(channel.Topic), "ubot-meta-role-pick") {
		roles, _ := s.GuildRoles(m.GuildID)
		for i := 0; i < len(roles); i++ {
			// That magic number that you see here is the default permissions integer when a new role is created
			if strings.ToLower(roles[i].Name) == strings.ToLower(arg) && roles[i].Permissions == 0 {
				_ = s.GuildMemberRoleAdd(m.GuildID, m.Member.User.ID, roles[i].ID)
				embed := NewEmbed().
					SetTitle("Added you to the "+roles[i].Name+" role!").
					AddField("Role", roles[i].Mention()).
					SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
					SetColor(0xf1c40f).MessageEmbed
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
				return
			}
		}
		perms := int64(0)

		roleinfo := discordgo.RoleParams{}
		roleinfo.Permissions = &perms
		roleinfo.Name = arg

		role, _ := s.GuildRoleCreate(m.GuildID, &roleinfo)
		_ = s.GuildMemberRoleAdd(m.GuildID, m.Member.User.ID, role.ID)
		embed := NewEmbed().
			SetTitle("Added you to the "+role.Name+" role!").
			AddField("Role", role.Mention()).
			SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
			SetColor(0xf1c40f).MessageEmbed
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
	} else {
		_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				TTS:     false,
				Content: "Couldn't add meta role, as the current channel is not marked for meta roles",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
}

func listAliases(s *discordgo.Session, m *discordgo.InteractionCreate) {
	channel, _ := s.Channel(m.ChannelID)
	topic := strings.ToLower(channel.Topic)

	if strings.Contains(topic, "ubot-macro:") {
		i := strings.Index(topic, "ubot-macro:")

		if i != -1 {
			i += len("ubot-macro:")
		} else {
			return
		}
		var fields []*discordgo.MessageEmbedField
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

					fields = append(fields, &discordgo.MessageEmbedField{
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
	} else {
		_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				TTS:     false,
				Content: "No aliases found in the channel",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
}

func removeMetarole(arg string, s *discordgo.Session, m *discordgo.InteractionCreate) {
	g, _ := s.Guild(m.GuildID)
	roles := g.Roles

	for i := 0; i < len(roles); i++ {
		if strings.ToLower(arg) == strings.ToLower(roles[i].Name) && roles[i].Permissions == 0 {
			_ = s.GuildMemberRoleRemove(m.GuildID, m.Member.User.ID, roles[i].ID)
			embed := NewEmbed().
				SetTitle("Removed metarole!").
				AddField("The following role has been removed", roles[i].Mention()).
				SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
				SetColor(0xf1c40f).MessageEmbed
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
			return
		}
	}
}
