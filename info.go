package main

import (
	"github.com/MadLadSquad/discordgo"
	"strconv"
)

func showUserInfo(arg string, s* discordgo.Session, m* discordgo.MessageCreate) {
	if arg != "" {
		usr, _ := s.User(sanitizePings(arg))
		premiumType := "None"
		switch usr.PremiumType {
		case 0:
			premiumType = "None"
			break
		case 1:
			premiumType = "Nitro Classic"
			break
		case 2:
			premiumType = "Nitro"
			break
		}

		embed := NewEmbed().
			SetTitle(usr.Username + "'s user information").
			SetThumbnail(usr.AvatarURL("")).
			AddField("Bot", strconv.FormatBool(usr.Bot)).
			AddField("Mention", usr.Mention()).
			AddField("Nitro type", premiumType).
			InlineAllFields().
			AddField("User ID", usr.ID).
			AddField("System", strconv.FormatBool(usr.System)).
			InlineAllFields().
			SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
			SetColor(0xf1c40f).MessageEmbed

		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
		if err != nil {
			return
		}
	} else {
		premiumType := "None"

		switch m.Author.PremiumType {
		case 0:
			premiumType = "None"
			break
		case 1:
			premiumType = "Nitro Classic"
			break
		case 2:
			premiumType = "Nitro"
			break
		}

		embed := NewEmbed().
			SetTitle(m.Author.Username + "'s user information").
			SetThumbnail(m.Author.AvatarURL("")).
			AddField("Bot", strconv.FormatBool(m.Author.Bot)).
			AddField("Mention", m.Author.Mention()).
			AddField("Nitro type", premiumType).
			InlineAllFields().
			AddField("User ID", m.Author.ID).
			AddField("System", strconv.FormatBool(m.Author.System)).
			InlineAllFields().
			SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
			SetColor(0xf1c40f).MessageEmbed

		_, sendEmbed := s.ChannelMessageSendEmbed(m.ChannelID, embed)
		if sendEmbed != nil {
			return
		}
	}
}

func showServerInfo(s* discordgo.Session, m* discordgo.MessageCreate) {
	guild, _ := s.State.Guild(m.GuildID)
	owner, _ := s.User(guild.OwnerID)

	mfaLevel := "None"
	premiumTier := "None"
	switch guild.MfaLevel {
	case 0:
		mfaLevel = "None"
		break
	case 1:
		mfaLevel = "Elevated"
		break
	}

	switch guild.PremiumTier {
	case 0:
		premiumTier = "None"
		break
	case 1:
		premiumTier = "Level One"
		break
	case 2:
		premiumTier = "Level Two"
		break
	case 3:
		premiumTier = "Level Three"
		break
	}

	verificationLevel := "None"
	switch guild.VerificationLevel {
	case 0:
		verificationLevel = "None"
		break
	case 1:
		verificationLevel = "Low"
		break
	case 2:
		verificationLevel = "Medium"
		break
	case 3:
		verificationLevel = "High"
		break
	case 4:
		verificationLevel = "Highest"
		break
	}

	//guildDescription := "None"
	if len(guild.Description) > 2 {
		//guildDescription = guild.Description
	}
	embed := NewEmbed().
		SetTitle(guild.Name + " server information").
		SetThumbnail(guild.IconURL()).
		AddField("Member Count", strconv.Itoa(len(guild.Members))).
		AddField("Region", guild.Region).
		AddField("Channel Count", strconv.Itoa(len(guild.Channels))).
		InlineAllFields().
		AddField("Server ID", guild.ID).
		AddField("Large", strconv.FormatBool(guild.Large)).
		AddField("AFK Timeout", strconv.Itoa(guild.AfkTimeout) + " seconds").
		InlineAllFields().
		AddField("Emoji Count", strconv.Itoa(len(guild.Emojis))).
		AddField("Owner", owner.Mention()).
		AddField("Role Count", strconv.Itoa(len(guild.Roles))).
		InlineAllFields().
		AddField("MFA Level", mfaLevel).
		AddField("Boost Level", premiumTier).
		AddField("Verification Level", verificationLevel).
		InlineAllFields().
		//AddField("Guild Description", guildDescription).
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
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
		AddField("If you want to submit bug reports report them here", "https://github.com/MadLadSquad/MadLadSquadBot/issues").
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		return
	}
}

func avatar(arg string, s* discordgo.Session, m* discordgo.MessageCreate) {
	if arg != "" {
		usr, _ := s.User(sanitizePings(arg))

		embed := NewEmbed().
			SetTitle(usr.Username + "'s avatar!").
			SetImage(usr.AvatarURL("")).
			SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
			SetColor(0xf1c40f).MessageEmbed

		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
		if err != nil {
			return
		}
	} else {
		embed := NewEmbed().
			SetTitle(m.Author.Username + "'s avatar!").
			SetImage(m.Author.AvatarURL("")).
			SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
			SetColor(0xf1c40f).MessageEmbed

		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
		if err != nil {
			return
		}
	}
}