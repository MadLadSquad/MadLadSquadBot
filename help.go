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

func createAliasHelp(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "alias-help",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Shows the user how to use command aliases",
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

func aliasHelp(s *discordgo.Session, m *discordgo.InteractionCreate) {
	embed := NewEmbed().
		SetTitle("Alias help").
		AddField("Introduction to aliases", "Aliases are a way of giving a new name to a pre-existing command on a per-channel basis").
		AddField("Syntax", "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"").
		AddField("Symbols", "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression").
		InlineAllFields().
		AddField("Example", "ubot-macro:set-game-role>set-meta-role;set-city>set-meta-role;set-fav-food-role>set-meta-role]").
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
