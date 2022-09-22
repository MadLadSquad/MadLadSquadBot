package main

import (
	"github.com/MadLadSquad/discordgo"
)

func createInvite(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "invite",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Returns a link for inviting the bot",
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB:    "invite",
			discordgo.EnglishUS:    "invite",
			discordgo.Bulgarian:    "покани",
			discordgo.ChineseCN:    "邀请",
			discordgo.ChineseTW:    "邀請",
			discordgo.Croatian:     "pozovi",
			discordgo.Czech:        "pozvat",
			discordgo.Danish:       "inviter",
			discordgo.Dutch:        "uitnodigen",
			discordgo.Finnish:      "kutsu",
			discordgo.French:       "inviter",
			discordgo.German:       "einladen",
			discordgo.Greek:        "Προσκάλεσε",
			discordgo.Hungarian:    "meghívása",
			discordgo.Italian:      "invita",
			discordgo.Japanese:     "招待",
			discordgo.Korean:       "초대하기",
			discordgo.Lithuanian:   "kviesti",
			discordgo.Norwegian:    "inviter",
			discordgo.Polish:       "zaproś",
			discordgo.PortugueseBR: "convidar",
			discordgo.Romanian:     "invită",
			discordgo.Russian:      "пригласить",
			discordgo.SpanishES:    "invitar",
			discordgo.Swedish:      "bjud",
			discordgo.Ukrainian:    "запросити",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "Returns a link for inviting the bot",
			discordgo.EnglishUS: "Returns a link for inviting the bot",
			discordgo.Bulgarian: "Дава линк за покана на бота в друг сървър",
		},
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
