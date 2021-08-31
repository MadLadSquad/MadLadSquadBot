package main

import (
	"github.com/MadLadSquad/discordgo"
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

func ban(arg [2]string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if arg[0] != "" && arg[1] != "" {
		if checkPerm(s, m, discordgo.PermissionAdministrator) {
			usr, err := s.User(sanitizePings(arg[0]))

			err = s.GuildBanCreateWithReason(m.GuildID, usr.ID, arg[1], 7)
			embed := NewEmbed().
				SetTitle("Banned a user!").
				AddField("The following user has been banned from the server", arg[0]).
				AddField("With reason", arg[1]).
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