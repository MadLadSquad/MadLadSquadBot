package main

import (
	"github.com/MadLadSquad/discordgo"
)

func help(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := NewEmbed().
		SetTitle("Help").
		AddField(prefix+"serverinfo", "Displays info about he current server").
		AddField(prefix+"userinfo <optional ping>", "Displays the info about you or the person you pinged").
		AddField(prefix+"help", "Shows this message").
		InlineAllFields().
		AddField(prefix+"kick <ping>", "Kicks the user from the current server").
		AddField(prefix+"ban <ping> <reason up to 100 words>", "Bans a user from the current server").
		AddField(prefix+"invite", "invites the bot into the server").
		InlineAllFields().
		AddField(prefix+"tos", "Displays the terms of service for the bot").
		AddField(prefix+"privacy", "Displays the privacy policy for the bot").
		AddField(prefix+"about", "Displays information about MadLad Squad and the bot").
		InlineAllFields().
		AddField(prefix+"verify", "If a role with the name \"Member\" or \"Members\" is found it will give the person that role").
		AddField(prefix+"pernik", "Displays a Pernik meme").
		InlineAllFields().
		AddField(prefix+"avatar <optional ping>", "Sends a link to the user's avatar").
		AddField(prefix+"mute <ping>", "Mutes a user by giving him a muted role").
		AddField(prefix+"sus", "Responds with a random sussy message").
		InlineAllFields().
		AddField(prefix+"generate-member-role", "Automatically generates a role for members used for verification").
		AddField(prefix+"set-welcome <ping to channel>", "Sets a given channel as the welcome channel").
		AddField(prefix+"set-event-tracking <ping to channel>", "Sets up event tracking in a channel").
		InlineAllFields().
		AddField(prefix+"set-text-only <ping to channel>", "Sets a channel as text only").
		AddField(prefix+"set-attachments-only <ping to channel>", "Sets a channel as attachments only").
		AddField(prefix+"set-links-only <ping to channel>", "Sets a channel as links only").
		InlineAllFields().
		AddField(prefix+"set-colour-role-channel <ping to channel>", "Sets a channel as one where colours can be selected").
		AddField(prefix+"set-colour-role <colour-name>", "Adds the specific colour role from the colour list to the user").
		AddField(prefix+"list-colour-roles", "Lists all available colours for the colour role").
		InlineAllFields().
		AddField(prefix+"set-meta-role <role-name>", "Gives/Creates a meta role(a role with default perms for customizing your user's profile)").
		AddField(prefix+"set-meta-role-channel <channel-ping>", "Sets a channel as a valid meta role channel").
		InlineAllFields().
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	embed2 := NewEmbed().
		SetTitle("Help2").
		AddField(prefix+"list-aliases", "Lists the aliases stored in the channel, and their respective command").
		AddField(prefix+"remove-meta-role <role name>", "Removes a meta role from your profile").
		AddField(prefix+"alias-help", "Shows a help message on how to construct aliases").
		InlineAllFields().
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	_, err := s.ChannelMessageSendEmbeds(m.ChannelID, []*discordgo.MessageEmbed{embed, embed2})
	if err != nil {
		return
	}
}

func invite(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := NewEmbed().
		SetTitle("Invite").
		AddField("Use this link", "https://discord.com/oauth2/authorize?client_id=697420452712284202&scope=bot&permissions=2134207679").
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		return
	}
}

func aliasHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := NewEmbed().
		SetTitle("Alias help").
		AddField("Introduction to aliases", "Aliases are a way of giving a new name to a pre-existing command on a per-channel basis").
		AddField("Syntax", "To add an alias all you need to do is find a channel and add the following string to the topic \"ubot-macro:\". After the ':' character, follow the following syntax \"alias-name>command;alias-name>command]\"").
		AddField("Symbols", "The \"ubot-macro:\", introduces a new macro sequence. The '>' character can be imagined as an arrow pointing to the command. The ';' symbol is used as a separator between aliases. The ']' symbol is used as a terminator for the expression").
		InlineAllFields().
		AddField("Example", "ubot-macro:set-game-role>set-meta-role;set-city>set-meta-role;set-fav-food-role>set-meta-role]").
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
