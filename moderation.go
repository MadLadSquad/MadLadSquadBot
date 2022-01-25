package main

import (
	"github.com/MadLadSquad/discordgo"
	"strings"
)

func kick(arg string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if arg != "" {
		if checkPerm(s, m, discordgo.PermissionAdministrator) {
			usr, _ := s.User(sanitizePings(arg))

			err := s.GuildMemberDelete(m.GuildID, usr.ID)
			if err != nil {
				return
			}

			embed := NewEmbed().
				SetTitle("Kicked a user!").
				AddField("The following user has been kicked from the server", usr.Mention()).
				SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
				SetColor(0xf1c40f).MessageEmbed

			_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
			if err != nil {
				return
			}
		} else {
			embed := NewEmbed().
				SetTitle("Unable to kick user!").
				AddField("The following user couldn't be kicked from the server! Reason: ", "Author doesn't have sufficient privileges!").
				SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
				SetColor(0xf1c40f).MessageEmbed

			_, sendEmbed := s.ChannelMessageSendEmbed(m.ChannelID, embed)
			if sendEmbed != nil {
				return
			}
		}
	}
}

func ban(arg []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if arg[0] != "" && arg[1] != "" {
		if checkPerm(s, m, discordgo.PermissionAdministrator) {
			usr, err := s.User(sanitizePings(arg[0]))

			reason := ""
			for i := 0; i < len(arg); i++ {
				reason += arg[i] + " "
			}

			err = s.GuildBanCreateWithReason(m.GuildID, usr.ID, arg[1], 7)
			embed := NewEmbed().
				SetTitle("Banned a user!").
				AddField("The following user has been banned from the server", arg[0]).
				AddField("With reason", reason).
				SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
				SetColor(0xf1c40f).MessageEmbed

			_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
			if err != nil {
				return
			}
		} else {
			embed := NewEmbed().
				SetTitle("Unable to ban user!").
				AddField("The user couldn't be banned from the server! Reason: ", "Author doesn't have sufficient privileges!").
				SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
				SetColor(0xf1c40f).MessageEmbed

			_, sendEmbed := s.ChannelMessageSendEmbed(m.ChannelID, embed)
			if sendEmbed != nil {
				return
			}
		}
	}
}

func verify(s *discordgo.Session, m *discordgo.MessageCreate) {
	guild, _ := s.State.Guild(m.GuildID)

	for j := 0; j < len(guild.Roles); j++ {
		if strings.ToLower(guild.Roles[j].Name) == "member" || strings.ToLower(guild.Roles[j].Name) == "members" {
			_ = s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, guild.Roles[j].ID)
			_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
			break
		}
	}
}

func mute(arg string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if checkPerm(s, m, discordgo.PermissionAdministrator) && arg != "" {
		user, _ := s.User(sanitizePings(arg))
		guild, _ := s.State.Guild(m.GuildID)

		for i := 0; i < len(guild.Roles); i++ {
			if strings.ToLower(guild.Roles[i].Name) == "mute" || strings.ToLower(guild.Roles[i].Name) == "muted" {
				_ = s.GuildMemberRoleAdd(m.GuildID, user.ID, guild.Roles[i].ID)
				embed := NewEmbed().
					SetTitle("Muted user!").
					SetThumbnail(user.AvatarURL("")).
					AddField("The following user has been muted", user.Mention()).
					SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
					SetColor(0xf1c40f).MessageEmbed

				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
				break
			}
		}
	}
}

func createMemberRole(s *discordgo.Session, m *discordgo.MessageCreate) {
	role, _ := s.GuildRoleCreate(m.GuildID)
	_, _ = s.GuildRoleEdit(m.GuildID, role.ID, "Members", role.Color, role.Hoist, role.Permissions, role.Mentionable)
}
