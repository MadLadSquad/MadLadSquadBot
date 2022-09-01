package main

import (
	"github.com/MadLadSquad/discordgo"
	"strings"
)

var (
	colours = map[string]uint{
		"Alice-Blue":             0xf0f8ff,
		"Antique-White":          0xfaebd7,
		"Aqua":                   0x00ffff,
		"Aquamarine":             0x7fffd4,
		"Azure-Mist":             0xf0ffff,
		"Beige":                  0xf5f5dc,
		"Bisque":                 0xffe4c4,
		"Black":                  0x000000,
		"Blanched-Almond":        0xffebcd,
		"Blue":                   0x0000ff,
		"Blue-violet":            0x8a2be2,
		"Auburn":                 0xa52a2a,
		"Burlywood":              0xdeb887,
		"Cadet-Blue":             0x5f9ea0,
		"Chartreuse":             0x7fff00,
		"Cinnamon":               0xd2691e,
		"Coral":                  0xff7f50,
		"Cornflower-Blue":        0x6495ed,
		"Cornsilk":               0xfff8dc,
		"Crimson":                0xdc143c,
		"Cyan":                   0x00ffff,
		"Dark-Blue":              0x00008b,
		"Dark-Cyan":              0x008b8b,
		"Dark-Goldenrod":         0xb8860b,
		"Dark-Gray":              0xa9a9a9,
		"Dark-Khaki":             0xbdb76b,
		"Dark-Magenta":           0x8b008b,
		"Dark-Olive-Green":       0x556b2f,
		"Dark-Orange":            0xff8c00,
		"Dark-Orchid":            0x9932cc,
		"Dark-Red":               0x8b0000,
		"Dark-Salmon":            0xe9967a,
		"Dark-Sea-Green":         0x8fbc8f,
		"Dark-Slate-Blue":        0x483d8b,
		"Dark-Slate-Gray":        0x2f4f4f,
		"Dark-Turquoise":         0x00ced1,
		"Dark-Violet":            0x9400d3,
		"Deep-Pink":              0xff1493,
		"Capri":                  0x00bfff,
		"Dim-Gray":               0x696969,
		"Dodger-Blue":            0x1e90ff,
		"Firebrick":              0xb22222,
		"Floral-White":           0xfffaf0,
		"Forest-Green":           0x228b22,
		"Magenta":                0xff00ff,
		"Gainsboro":              0xdcdcdc,
		"Ghost-White":            0xf8f8ff,
		"Gold":                   0xffd700,
		"Goldenrod":              0xdaa520,
		"Gray":                   0x808080,
		"Office-Green":           0x008000,
		"Green-yellow":           0xadff2f,
		"Honeydew":               0xf0fff0,
		"Hot-Pink":               0xff69b4,
		"Chestnut":               0xcd5c5c,
		"Indigo":                 0x4b0082,
		"Ivory":                  0xfffff0,
		"Khaki":                  0xf0e68c,
		"Lavender":               0xe6e6fa,
		"Lavender-Blush":         0xfff0f5,
		"Lawn-Green":             0x7cfc00,
		"Lemon-Chiffon":          0xfffacd,
		"Light-Blue":             0xadd8e6,
		"Light-Coral":            0xf08080,
		"Light-Cyan":             0xe0ffff,
		"Light-Goldenrod-Yellow": 0xfafad2,
		"Light-Gray":             0xd3d3d3,
		"Light-Green":            0x90ee90,
		"Light-Pink":             0xffb6c1,
		"Light-Salmon":           0xffa07a,
		"Light-Sea-Green":        0x20b2aa,
		"Light-Sky-Blue":         0x87cefa,
		"Light-Slate-Gray":       0x778899,
		"Light-Steel-Blue":       0xb0c4de,
		"Light-Yellow":           0xffffe0,
		"Lime":                   0x00ff00,
		"Lime-Green":             0x32cd32,
		"Linen":                  0xfaf0e6,
		"Fuchsia":                0xff00ff,
		"Wood-Red":               0x7f0000,
		"Tropical-Green":         0x66cdaa,
		"Medium-Blue":            0x0000cd,
		"Medium-Orchid":          0xba55d3,
		"Medium-Purple":          0x9370db,
		"Medium-Sea-Green":       0x3cb371,
		"Medium-Slate-Blue":      0x7b68ee,
		"Medium-Spring-Green":    0x00fa9a,
		"Medium-Turquoise":       0x48d1cc,
		"Red-violet":             0xc71585,
		"Midnight-Blue":          0x191970,
		"Mint-Cream":             0xf5fffa,
		"Misty-Rose":             0xffe4e1,
		"SandySahara":            0xffe4b5,
		"Navajo-White":           0xffdead,
		"Navy-Blue":              0x000080,
		"Old-Lace":               0xfdf5e6,
		"Olive":                  0x808000,
		"Olive-Drab":             0x6b8e23,
		"Orange":                 0xffa500,
		"Orange-red":             0xff4500,
		"Orchid":                 0xda70d6,
		"Pale-Goldenrod":         0xeee8aa,
		"Pale-Green":             0x98fb98,
		"Pale-Blue":              0xafeeee,
		"Pale-Violet-red":        0xdb7093,
		"Papaya-Whip":            0xffefd5,
		"Peach-Puff":             0xffdab9,
		"Peru":                   0xcd853f,
		"Pink":                   0xffc0cb,
		"Plum":                   0xdda0dd,
		"Powder-Blue":            0xb0e0e6,
		"Sienna":                 0x7f007f,
		"Red":                    0xff0000,
		"Rosy-Brown":             0xbc8f8f,
		"Royal-Blue":             0x4169e1,
		"Saddle-Brown":           0x8b4513,
		"Salmon":                 0xfa8072,
		"Sandy-Brown":            0xf4a460,
		"Sea-Green":              0x2e8b57,
		"Seashell":               0xfff5ee,
		"Dark-Sand":              0xa0522d,
		"Silver":                 0xc0c0c0,
		"Sky-Blue":               0x87ceeb,
		"Slate-Blue":             0x6a5acd,
		"Slate-Gray":             0x708090,
		"Snow":                   0xfffafa,
		"Spring Green":           0x00ff7f,
		"Steel-Blue":             0x4682b4,
		"Tan":                    0xd2b48c,
		"Teal":                   0x008080,
		"Thistle":                0xd8bfd8,
		"Tomato":                 0xff6347,
		"Light-Cyan-2":           0x40e0d0,
		"Violet":                 0xee82ee,
		"Wheat":                  0xf5deb3,
		"White":                  0xffffff,
		"White-Smoke":            0xf5f5f5,
		"Electric-Yellow":        0xffff00,
		"Yellow-green":           0x9acd32,
	}
)

func listColours(s *discordgo.Session, m *discordgo.MessageCreate) {
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

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func giveColour(arg string, s *discordgo.Session, m *discordgo.MessageCreate) {
	channel, _ := s.Channel(m.ChannelID)
	if strings.Contains(strings.ToLower(channel.Topic), "ubot-colour-pick") {
		bFound := false
		colour := 0
		var name = ""

		for key, val := range colours {
			if strings.ToLower(key) == strings.ToLower(arg) {
				bFound = true
				colour = int(val)
				name = key
				break
			}
		}

		if bFound {
			roles, _ := s.GuildRoles(m.GuildID)
			for i := 0; i < len(roles); i++ {
				if name == roles[i].Name {
					_ = s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, roles[i].ID)
					embed := NewEmbed().
						SetTitle("Added you to the "+name+" role!").
						AddField("Role", roles[i].Mention()).
						SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
						SetColor(0xf1c40f).MessageEmbed
					_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
					return
				}
			}

			perms := int64(0)

			roledata := discordgo.RoleParams{}
			roledata.Permissions = &perms
			roledata.Color = &colour

			role, _ := s.GuildRoleCreate(m.GuildID, &roledata)
			_ = s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, role.ID)
			embed := NewEmbed().
				SetTitle("Added you to the "+name+" role!").
				AddField("Role", role.Mention()).
				SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
				SetColor(0xf1c40f).MessageEmbed
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!")
		}
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-colour-role-channel\" command in order to set up colour roles!")
	}
}

func giveMetarole(arg string, s *discordgo.Session, m *discordgo.MessageCreate) {
	channel, _ := s.Channel(m.ChannelID)

	if strings.Contains(strings.ToLower(channel.Topic), "ubot-meta-role-pick") {
		roles, _ := s.GuildRoles(m.GuildID)
		for i := 0; i < len(roles); i++ {
			// That magic number that you see here is the default permissions integer when a new role is created
			if strings.ToLower(roles[i].Name) == strings.ToLower(arg) && roles[i].Permissions == 0 {
				_ = s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, roles[i].ID)
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
		_ = s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, role.ID)
		embed := NewEmbed().
			SetTitle("Added you to the "+role.Name+" role!").
			AddField("Role", role.Mention()).
			SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
			SetColor(0xf1c40f).MessageEmbed
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Channel not marked as \"ubot-meta-role-pick\", contact your server's moderator to run the \"set-meta-role-channel\" command in order to set up colour roles!")
	}
}

func listAliases(s *discordgo.Session, m *discordgo.MessageCreate) {
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
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, "No aliases found in this channel!")
	}
}

func removeMetarole(arg string, s *discordgo.Session, m *discordgo.MessageCreate) {
	g, _ := s.Guild(m.GuildID)
	roles := g.Roles

	for i := 0; i < len(roles); i++ {
		if strings.ToLower(arg) == strings.ToLower(roles[i].Name) && roles[i].Permissions == 0 {
			_ = s.GuildMemberRoleRemove(m.GuildID, m.Author.ID, roles[i].ID)
			embed := NewEmbed().
				SetTitle("Removed metarole!").
				AddField("The following role has been removed", roles[i].Mention()).
				SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
				SetColor(0xf1c40f).MessageEmbed
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
			return
		}
	}
}
