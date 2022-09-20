package main

import (
	"github.com/MadLadSquad/discordgo"
)

func createInvite(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "invite",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Returns a link for inviting the bot",
	}
	_, _ = s.ApplicationCommandCreate(s.State.User.ID, "", command)
}

func invite(s *discordgo.Session, m *discordgo.InteractionCreate) {
	embed := NewEmbed().
		SetTitle("Invite").
		AddField("Use this link", "https://discord.com/oauth2/authorize?client_id=697420452712284202&scope=bot&permissions=2134207679").
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
