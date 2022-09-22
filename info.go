package main

import (
	"github.com/MadLadSquad/discordgo"
	"strconv"
)

func createUserInfo(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "userinfo",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Returns information about the caller or the mentioned user",
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "userinfo",
			discordgo.EnglishUS: "userinfo",
			discordgo.Bulgarian: "потребителска-информация",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "Returns information about the caller or the mentioned user",
			discordgo.EnglishUS: "Returns information about the caller or the mentioned user",
			discordgo.Bulgarian: "Връща информация за потребителя, извикал командата, или споменатият потребител",
		},
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "The user's mention",
				Type:        discordgo.ApplicationCommandOptionString,
				Description: "Provide a mention to get the info of a user rather than yourself",
				Required:    false,
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "The user's mention",
					discordgo.EnglishUS: "The user's mention",
					discordgo.Bulgarian: "Пинг към друг потребител",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "Provide a mention to get the info of a user rather than yourself",
					discordgo.EnglishUS: "Provide a mention to get the info of a user rather than yourself",
					discordgo.Bulgarian: "Споменете друг потребител, за да вземете инфомацията му, вместо вашата",
				},
			},
		},
	}
	_, _ = s.ApplicationCommandCreate(s.State.User.ID, "", command)
}

func createServerInfo(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "serverinfo",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Returns information about the current server",
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "serverinfo",
			discordgo.EnglishUS: "serverinfo",
			discordgo.Bulgarian: "сървърна-информация",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "Returns information about the current server",
			discordgo.EnglishUS: "Returns information about the current server",
			discordgo.Bulgarian: "Връща информация за даденият сървър",
		},
	}
	_, _ = s.ApplicationCommandCreate(s.State.User.ID, "", command)
}

func createPrivacy(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "privacy",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Returns the privacy policy statement",
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "privacy",
			discordgo.EnglishUS: "privacy",
			discordgo.Bulgarian: "поверителност",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "Returns the privacy policy statement",
			discordgo.EnglishUS: "Returns the privacy policy statement",
			discordgo.Bulgarian: "Връща информация, за поверителността и обработката на данни, която извършваме",
		},
	}
	_, _ = s.ApplicationCommandCreate(s.State.User.ID, "", command)
}

func createTos(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "tos",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Returns the terms of service statement",
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "tos",
			discordgo.EnglishUS: "tos",
			discordgo.Bulgarian: "правила-за-употреба",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "Returns the terms of service statement",
			discordgo.EnglishUS: "Returns the terms of service statement",
			discordgo.Bulgarian: "Връща правилата за употреба на продукта",
		},
	}
	_, _ = s.ApplicationCommandCreate(s.State.User.ID, "", command)
}

func createAbout(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "about",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Returns information about the bot and its developers",
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "about",
			discordgo.EnglishUS: "about",
			discordgo.Bulgarian: "за-бота",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "Returns information about the bot and its developers",
			discordgo.EnglishUS: "Returns information about the bot and its developers",
			discordgo.Bulgarian: "Връща информация, за бота и неговите разработчици",
		},
	}
	_, _ = s.ApplicationCommandCreate(s.State.User.ID, "", command)
}

func createAvatar(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "avatar",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Returns the avatar image of the caller or the mentioned user",
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "avatar",
			discordgo.EnglishUS: "avatar",
			discordgo.Bulgarian: "аватар",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "Returns the avatar image of the caller or the mentioned user",
			discordgo.EnglishUS: "Returns the avatar image of the caller or the mentioned user",
			discordgo.Bulgarian: "Връща аватара на, потребителя, извикал командата или на потребителя споменат като аргумент",
		},
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "mention",
				Type:        discordgo.ApplicationCommandOptionString,
				Description: "Provide a mention to get another person's avatar rather than yours",
				Required:    false,
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "mention",
					discordgo.EnglishUS: "mention",
					discordgo.Bulgarian: "споменат-потребител",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "Provide a mention to get another person's avatar rather than yours",
					discordgo.EnglishUS: "Provide a mention to get another person's avatar rather than yours",
					discordgo.Bulgarian: "Споменете друг потребител, за да вземете неговият аватар, вместо вашият",
				},
			},
		},
	}
	_, _ = s.ApplicationCommandCreate(s.State.User.ID, "", command)
}

func showUserInfo(arg string, s *discordgo.Session, m *discordgo.InteractionCreate) {
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
			SetTitle(usr.Username+"'s user information").
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

		_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				TTS:     false,
				Content: "",
				Flags:   discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{
					embed,
				},
			},
		})
	} else {
		premiumType := "None"

		switch m.Member.User.PremiumType {
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
			SetTitle(m.Member.User.Username+"'s user information").
			SetThumbnail(m.Member.User.AvatarURL("")).
			AddField("Bot", strconv.FormatBool(m.Member.User.Bot)).
			AddField("Mention", m.Member.User.Mention()).
			AddField("Nitro type", premiumType).
			InlineAllFields().
			AddField("User ID", m.Member.User.ID).
			AddField("System", strconv.FormatBool(m.Member.User.System)).
			InlineAllFields().
			SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
			SetColor(0xf1c40f).MessageEmbed

		_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				TTS:     false,
				Content: "",
				Flags:   discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{
					embed,
				},
			},
		})
	}
}

func showServerInfo(s *discordgo.Session, m *discordgo.InteractionCreate) {
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

	guildDescription := "None"
	if len(guild.Description) > 2 {
		guildDescription = guild.Description
	}
	embed := NewEmbed().
		SetTitle(guild.Name+" server information").
		SetThumbnail(guild.IconURL()).
		AddField("Member Count", strconv.Itoa(len(guild.Members))).
		AddField("Region", guild.Region).
		AddField("Channel Count", strconv.Itoa(len(guild.Channels))).
		InlineAllFields().
		AddField("Server ID", guild.ID).
		AddField("Large", strconv.FormatBool(guild.Large)).
		AddField("AFK Timeout", strconv.Itoa(guild.AfkTimeout)+" seconds").
		InlineAllFields().
		AddField("Emoji Count", strconv.Itoa(len(guild.Emojis))).
		AddField("Owner", owner.Mention()).
		AddField("Role Count", strconv.Itoa(len(guild.Roles))).
		InlineAllFields().
		AddField("MFA Level", mfaLevel).
		AddField("Boost Level", premiumTier).
		AddField("Verification Level", verificationLevel).
		InlineAllFields().
		AddField("Guild Description", guildDescription).
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

func privacyPolicy(s *discordgo.Session, m *discordgo.InteractionCreate) {
	embed := NewEmbed().
		SetTitle("Privacy policy").
		AddField("This bot is completely open source under the MIT permissive license", "https://github.com/MadLadSquad/MadLadSquadBot/blob/master/LICENSE").
		AddField("This means that you can freely read or modify the code!", "This is important because you have the freedom and power to see which data is collected and possibly try to change that!").
		AddField("That being said this bot stores no user data!", "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!").
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

func termsOfService(s *discordgo.Session, m *discordgo.InteractionCreate) {
	embed := NewEmbed().
		SetTitle("Terms of service").
		AddField("We don't have any!", "Use this bot for whatever you like as long as you respect it's licence: https://github.com/MadLadSquad/MadLadSquadBot/blob/master/LICENSE").
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

func about(s *discordgo.Session, m *discordgo.InteractionCreate) {
	embed := NewEmbed().
		SetTitle("About").
		SetThumbnail("https://cdn.discordapp.com/avatars/697420452712284202/924342db89aa1f0acd5239646a835bec.png").
		AddField("This bot is developed and maintained by MadLad Squad", "https://github.com/MadLadSquad/ , https://madladsquad.com/").
		AddField("If you have any problems, want to suggest a feature or just chat join our discord server", "https://discord.com/invite/heA4FpTaTj").
		AddField("If you want to submit bug reports report them here", "https://github.com/MadLadSquad/MadLadSquadBot/issues").
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

func avatar(arg string, s *discordgo.Session, m *discordgo.InteractionCreate) {
	if arg != "" {
		usr, _ := s.User(sanitizePings(arg))

		embed := NewEmbed().
			SetTitle(usr.Username+"'s avatar!").
			SetImage(usr.AvatarURL("")).
			SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
			SetColor(0xf1c40f).MessageEmbed

		_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				TTS:     false,
				Content: "",
				Flags:   discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{
					embed,
				},
			},
		})
	} else {
		embed := NewEmbed().
			SetTitle(m.Member.User.Username+"'s avatar!").
			SetImage(m.Member.User.AvatarURL("")).
			SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
			SetColor(0xf1c40f).MessageEmbed

		_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				TTS:     false,
				Content: "",
				Flags:   discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{
					embed,
				},
			},
		})
	}
}
