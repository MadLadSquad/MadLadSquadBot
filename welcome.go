package main

import (
	"github.com/MadLadSquad/discordgo"
	"strings"
)

func welcome(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	guild, _ := s.State.Guild(m.GuildID)

	timestamp := m.JoinedAt.String()

	embed := NewEmbed().
		SetTitle("Welcome to "+guild.Name+", "+m.User.Username+"!").
		SetThumbnail(m.User.AvatarURL("")).
		AddField("Joined on", timestamp).
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	for i := 0; i < len(guild.Channels); i++ {
		if strings.ToLower(guild.Channels[i].Name) == "welcome" || strings.Contains(strings.ToLower(guild.Channels[i].Topic), "ubot-welcome") {
			_, err := s.ChannelMessageSendEmbed(guild.Channels[i].ID, embed)
			if err != nil {
				return
			}
		}
	}
}

func bye(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	guild, _ := s.State.Guild(m.GuildID)

	embed := NewEmbed().
		SetTitle("Goodbye to "+m.User.Username+"!").
		SetThumbnail(m.User.AvatarURL("")).
		AddField("The following person has left "+guild.Name, m.User.Mention()).
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	for i := 0; i < len(guild.Channels); i++ {
		if strings.ToLower(guild.Channels[i].Name) == "welcome" || strings.Contains(strings.ToLower(guild.Channels[i].Topic), "ubot-welcome") {
			_, err := s.ChannelMessageSendEmbed(guild.Channels[i].ID, embed)
			if err != nil {
				return
			}
		}
	}
}

func invited(s *discordgo.Session, m *discordgo.GuildCreate) {
	embed := NewEmbed().
		SetTitle("Thanks for inviting me").
		SetThumbnail("https://cdn.discordapp.com/avatars/697420452712284202/924342db89aa1f0acd5239646a835bec.png").
		AddField("To view the commands for the bot type in", "&help").
		AddField("If you have a problem with the bot submit an issue on github!", "https://github.com/MadLadSquad/MadLadSquadBot").
		AddField("Want to help the development, chat, suggest features or other software join our discord", "https://discord.gg/4wgH8ZE").
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	_, err := s.ChannelMessageSendEmbed(m.Channels[0].ID, embed)
	if err != nil {
		return
	}
}
