package main

import "github.com/MadLadSquad/discordgo"

func welcome(s* discordgo.Session, m* discordgo.MessageCreate) {

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

func privacyPolicy(s* discordgo.Session, m* discordgo.MessageCreate) {
	embed := NewEmbed().
		SetTitle("Privacy policy").
		AddField("This bot is completely open source under the MIT permissive license", "https://github.com/MadLadSquad/MadLadSquadBot/blob/master/LICENSE").
		AddField("This means that you can freely read or modify the code!", "This is important because you have the freedom and power to see which data is collected and possibly try to change that!").
		AddField("That being said this bot stores no user data!", "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!").
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		return
	}
}

func termsOfService(s* discordgo.Session, m* discordgo.MessageCreate) {
	embed := NewEmbed().
		SetTitle("Terms of service").
		AddField("We don't have any!", "Use this bot for whatever you like as long as you respect it's licence: https://github.com/MadLadSquad/MadLadSquadBot/blob/master/LICENSE").
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		return
	}
}

func about(s* discordgo.Session, m* discordgo.MessageCreate) {
	embed := NewEmbed().
		SetTitle("About").
		SetThumbnail("https://cdn.discordapp.com/avatars/697420452712284202/924342db89aa1f0acd5239646a835bec.png").
		AddField("This bot is developed and maintained by MadLad Squad", "https://github.com/MadLadSquad/ , https://madladsquad.com/").
		AddField("If you have any problems, want to suggest a feature or just chat join our discord server", "https://discord.com/invite/heA4FpTaTj").
		AddField("If you want to support bug reports report them here", "https://github.com/MadLadSquad/MadLadSquadBot/issues").
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		return
	}
}