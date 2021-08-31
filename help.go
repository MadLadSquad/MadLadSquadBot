package main

import "github.com/MadLadSquad/discordgo"

func help(s* discordgo.Session, m* discordgo.MessageCreate) {
	embed := NewEmbed().
		SetTitle("Help").
		AddField(prefix+"serverinfo", "Displays info about he current server").
		AddField(prefix+"userinfo <optional ping>", "Displays the info about you or the person you pinged").
		AddField(prefix+"help", "Shows this message").
		InlineAllFields().
		AddField(prefix+"kick <ping>", "Kicks the user from the current server").
		AddField(prefix+"ban <ping> <reason up to 100 words>", "Bans a user from the current server").
		InlineAllFields().
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		return
	}
}