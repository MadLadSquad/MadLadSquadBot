package main

import (
	"fmt"
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
				Name:        "mention",
				Type:        discordgo.ApplicationCommandOptionString,
				Description: "Provide a mention to get the info of a user rather than yourself",
				Required:    false,
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "mention",
					discordgo.EnglishUS: "mention",
					discordgo.Bulgarian: "пинг-към-потребител",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "Provide a mention to get the info of a user rather than yourself",
					discordgo.EnglishUS: "Provide a mention to get the info of a user rather than yourself",
					discordgo.Bulgarian: "Споменете друг потребител, за да вземете инфомацията му, вместо вашата",
				},
			},
		},
	}
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
	if err != nil {
		fmt.Println(err)
	}
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
			SetFooter(footerTranslations[m.Locale], "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
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
			SetFooter(footerTranslations[m.Locale], "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
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
		SetThumbnail(guild.IconURL("256")).
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
		SetFooter(footerTranslations[m.Locale], "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
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
	embedStrings := []map[discordgo.Locale]string{
		{
			discordgo.EnglishGB:    "Privacy policy",
			discordgo.EnglishUS:    "Privacy policy",
			discordgo.Bulgarian:    "Правила за поверителност",
			discordgo.ChineseCN:    "Privacy policy",
			discordgo.ChineseTW:    "Privacy policy",
			discordgo.Croatian:     "Privacy policy",
			discordgo.Czech:        "Privacy policy",
			discordgo.Danish:       "Privacy policy",
			discordgo.Dutch:        "Privacy policy",
			discordgo.Finnish:      "Privacy policy",
			discordgo.French:       "Privacy policy",
			discordgo.German:       "Privacy policy",
			discordgo.Greek:        "Privacy policy",
			discordgo.Hungarian:    "Privacy policy",
			discordgo.Italian:      "Privacy policy",
			discordgo.Japanese:     "Privacy policy",
			discordgo.Korean:       "Privacy policy",
			discordgo.Lithuanian:   "Privacy policy",
			discordgo.Norwegian:    "Privacy policy",
			discordgo.Polish:       "Privacy policy",
			discordgo.PortugueseBR: "Privacy policy",
			discordgo.Romanian:     "Privacy policy",
			discordgo.Russian:      "Privacy policy",
			discordgo.SpanishES:    "Privacy policy",
			discordgo.Swedish:      "Privacy policy",
			discordgo.Ukrainian:    "Privacy policy",
			discordgo.Hindi:        "Privacy policy",
			discordgo.Thai:         "Privacy policy",
			discordgo.Turkish:      "Privacy policy",
			discordgo.Vietnamese:   "Privacy policy",
			discordgo.Unknown:      "Privacy policy",
		},
		{
			discordgo.EnglishGB:    "This bot is completely open source under the MIT permissive license",
			discordgo.EnglishUS:    "This bot is completely open source under the MIT permissive license",
			discordgo.Bulgarian:    "Този бот е напълно направен с технологии с отворен код и е пуснат под свободният лиценз MIT",
			discordgo.ChineseCN:    "This bot is completely open source under the MIT permissive license",
			discordgo.ChineseTW:    "This bot is completely open source under the MIT permissive license",
			discordgo.Croatian:     "This bot is completely open source under the MIT permissive license",
			discordgo.Czech:        "This bot is completely open source under the MIT permissive license",
			discordgo.Danish:       "This bot is completely open source under the MIT permissive license",
			discordgo.Dutch:        "This bot is completely open source under the MIT permissive license",
			discordgo.Finnish:      "This bot is completely open source under the MIT permissive license",
			discordgo.French:       "This bot is completely open source under the MIT permissive license",
			discordgo.German:       "This bot is completely open source under the MIT permissive license",
			discordgo.Greek:        "This bot is completely open source under the MIT permissive license",
			discordgo.Hungarian:    "This bot is completely open source under the MIT permissive license",
			discordgo.Italian:      "This bot is completely open source under the MIT permissive license",
			discordgo.Japanese:     "This bot is completely open source under the MIT permissive license",
			discordgo.Korean:       "This bot is completely open source under the MIT permissive license",
			discordgo.Lithuanian:   "This bot is completely open source under the MIT permissive license",
			discordgo.Norwegian:    "This bot is completely open source under the MIT permissive license",
			discordgo.Polish:       "This bot is completely open source under the MIT permissive license",
			discordgo.PortugueseBR: "This bot is completely open source under the MIT permissive license",
			discordgo.Romanian:     "This bot is completely open source under the MIT permissive license",
			discordgo.Russian:      "This bot is completely open source under the MIT permissive license",
			discordgo.SpanishES:    "This bot is completely open source under the MIT permissive license",
			discordgo.Swedish:      "This bot is completely open source under the MIT permissive license",
			discordgo.Ukrainian:    "This bot is completely open source under the MIT permissive license",
			discordgo.Hindi:        "This bot is completely open source under the MIT permissive license",
			discordgo.Thai:         "This bot is completely open source under the MIT permissive license",
			discordgo.Turkish:      "This bot is completely open source under the MIT permissive license",
			discordgo.Vietnamese:   "This bot is completely open source under the MIT permissive license",
			discordgo.Unknown:      "This bot is completely open source under the MIT permissive license",
		},
		{
			discordgo.EnglishGB:    "This means that you can freely read or modify the code!",
			discordgo.EnglishUS:    "This means that you can freely read or modify the code!",
			discordgo.Bulgarian:    "Това означава, че можете свободно да четете, използвате, копирате и променяте кода!",
			discordgo.ChineseCN:    "This means that you can freely read or modify the code!",
			discordgo.ChineseTW:    "This means that you can freely read or modify the code!",
			discordgo.Croatian:     "This means that you can freely read or modify the code!",
			discordgo.Czech:        "This means that you can freely read or modify the code!",
			discordgo.Danish:       "This means that you can freely read or modify the code!",
			discordgo.Dutch:        "This means that you can freely read or modify the code!",
			discordgo.Finnish:      "This means that you can freely read or modify the code!",
			discordgo.French:       "This means that you can freely read or modify the code!",
			discordgo.German:       "This means that you can freely read or modify the code!",
			discordgo.Greek:        "This means that you can freely read or modify the code!",
			discordgo.Hungarian:    "This means that you can freely read or modify the code!",
			discordgo.Italian:      "This means that you can freely read or modify the code!",
			discordgo.Japanese:     "This means that you can freely read or modify the code!",
			discordgo.Korean:       "This means that you can freely read or modify the code!",
			discordgo.Lithuanian:   "This means that you can freely read or modify the code!",
			discordgo.Norwegian:    "This means that you can freely read or modify the code!",
			discordgo.Polish:       "This means that you can freely read or modify the code!",
			discordgo.PortugueseBR: "This means that you can freely read or modify the code!",
			discordgo.Romanian:     "This means that you can freely read or modify the code!",
			discordgo.Russian:      "This means that you can freely read or modify the code!",
			discordgo.SpanishES:    "This means that you can freely read or modify the code!",
			discordgo.Swedish:      "This means that you can freely read or modify the code!",
			discordgo.Ukrainian:    "This means that you can freely read or modify the code!",
			discordgo.Hindi:        "This means that you can freely read or modify the code!",
			discordgo.Thai:         "This means that you can freely read or modify the code!",
			discordgo.Turkish:      "This means that you can freely read or modify the code!",
			discordgo.Vietnamese:   "This means that you can freely read or modify the code!",
			discordgo.Unknown:      "This means that you can freely read or modify the code!",
		},
		{
			discordgo.EnglishGB:    "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.EnglishUS:    "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Bulgarian:    "Това е важно, защото ви позволява по напълно прозрачен начин да разберете, какви данни използваме, запазваме и обработваме, и ви позволява сами да направите копие на програмата и да я промените",
			discordgo.ChineseCN:    "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.ChineseTW:    "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Croatian:     "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Czech:        "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Danish:       "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Dutch:        "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Finnish:      "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.French:       "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.German:       "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Greek:        "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Hungarian:    "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Italian:      "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Japanese:     "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Korean:       "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Lithuanian:   "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Norwegian:    "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Polish:       "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.PortugueseBR: "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Romanian:     "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Russian:      "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.SpanishES:    "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Swedish:      "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Ukrainian:    "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Hindi:        "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Thai:         "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Turkish:      "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Vietnamese:   "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
			discordgo.Unknown:      "This is important because you have the freedom and power to see which data is collected and possibly try to change that!",
		},
		{
			discordgo.EnglishGB:    "That being said this bot stores no user data!",
			discordgo.EnglishUS:    "That being said this bot stores no user data!",
			discordgo.Bulgarian:    "Ние, обаче не събираме никакви потребителски данни!",
			discordgo.ChineseCN:    "That being said this bot stores no user data!",
			discordgo.ChineseTW:    "That being said this bot stores no user data!",
			discordgo.Croatian:     "That being said this bot stores no user data!",
			discordgo.Czech:        "That being said this bot stores no user data!",
			discordgo.Danish:       "That being said this bot stores no user data!",
			discordgo.Dutch:        "That being said this bot stores no user data!",
			discordgo.Finnish:      "That being said this bot stores no user data!",
			discordgo.French:       "That being said this bot stores no user data!",
			discordgo.German:       "That being said this bot stores no user data!",
			discordgo.Greek:        "That being said this bot stores no user data!",
			discordgo.Hungarian:    "That being said this bot stores no user data!",
			discordgo.Italian:      "That being said this bot stores no user data!",
			discordgo.Japanese:     "That being said this bot stores no user data!",
			discordgo.Korean:       "That being said this bot stores no user data!",
			discordgo.Lithuanian:   "That being said this bot stores no user data!",
			discordgo.Norwegian:    "That being said this bot stores no user data!",
			discordgo.Polish:       "That being said this bot stores no user data!",
			discordgo.PortugueseBR: "That being said this bot stores no user data!",
			discordgo.Romanian:     "That being said this bot stores no user data!",
			discordgo.Russian:      "That being said this bot stores no user data!",
			discordgo.SpanishES:    "That being said this bot stores no user data!",
			discordgo.Swedish:      "That being said this bot stores no user data!",
			discordgo.Ukrainian:    "That being said this bot stores no user data!",
			discordgo.Hindi:        "That being said this bot stores no user data!",
			discordgo.Thai:         "That being said this bot stores no user data!",
			discordgo.Turkish:      "That being said this bot stores no user data!",
			discordgo.Vietnamese:   "That being said this bot stores no user data!",
			discordgo.Unknown:      "That being said this bot stores no user data!",
		},
		{
			discordgo.EnglishGB:    "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.EnglishUS:    "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Bulgarian:    "Всичката информация, която е нужна за правилното функциониране на бота, е събрана чрез публичният API на Discord и бива запазвана за не повече, от крайният живот на дадената команда!",
			discordgo.ChineseCN:    "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.ChineseTW:    "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Croatian:     "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Czech:        "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Danish:       "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Dutch:        "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Finnish:      "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.French:       "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.German:       "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Greek:        "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Hungarian:    "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Italian:      "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Japanese:     "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Korean:       "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Lithuanian:   "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Norwegian:    "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Polish:       "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.PortugueseBR: "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Romanian:     "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Russian:      "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.SpanishES:    "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Swedish:      "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Ukrainian:    "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Hindi:        "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Thai:         "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Turkish:      "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Vietnamese:   "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
			discordgo.Unknown:      "All data necessary for the operation of this bot is acquired publicly using the Discord Bot API and is stored for no longer than the lifetime of the given command!",
		},
	}
	embed := NewEmbed().
		SetTitle(embedStrings[0][m.Locale]).
		AddField(embedStrings[1][m.Locale], "https://github.com/MadLadSquad/MadLadSquadBot/blob/master/LICENSE").
		AddField(embedStrings[2][m.Locale], embedStrings[3][m.Locale]).
		AddField(embedStrings[4][m.Locale], embedStrings[5][m.Locale]).
		SetFooter(footerTranslations[m.Locale], "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
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
	embedStrings := []map[discordgo.Locale]string{
		{
			discordgo.EnglishGB:    "Terms of service",
			discordgo.EnglishUS:    "Terms of service",
			discordgo.Bulgarian:    "Правила за употребата",
			discordgo.ChineseCN:    "Terms of service",
			discordgo.ChineseTW:    "Terms of service",
			discordgo.Croatian:     "Terms of service",
			discordgo.Czech:        "Terms of service",
			discordgo.Danish:       "Terms of service",
			discordgo.Dutch:        "Terms of service",
			discordgo.Finnish:      "Terms of service",
			discordgo.French:       "Terms of service",
			discordgo.German:       "Terms of service",
			discordgo.Greek:        "Terms of service",
			discordgo.Hungarian:    "Terms of service",
			discordgo.Italian:      "Terms of service",
			discordgo.Japanese:     "Terms of service",
			discordgo.Korean:       "Terms of service",
			discordgo.Lithuanian:   "Terms of service",
			discordgo.Norwegian:    "Terms of service",
			discordgo.Polish:       "Terms of service",
			discordgo.PortugueseBR: "Terms of service",
			discordgo.Romanian:     "Terms of service",
			discordgo.Russian:      "Terms of service",
			discordgo.SpanishES:    "Terms of service",
			discordgo.Swedish:      "Terms of service",
			discordgo.Ukrainian:    "Terms of service",
			discordgo.Hindi:        "Terms of service",
			discordgo.Thai:         "Terms of service",
			discordgo.Turkish:      "Terms of service",
			discordgo.Vietnamese:   "Terms of service",
			discordgo.Unknown:      "Terms of service",
		},
		{
			discordgo.EnglishGB:    "We don't have any!",
			discordgo.EnglishUS:    "We don't have any!",
			discordgo.Bulgarian:    "Нямаме такива!",
			discordgo.ChineseCN:    "We don't have any!",
			discordgo.ChineseTW:    "We don't have any!",
			discordgo.Croatian:     "We don't have any!",
			discordgo.Czech:        "We don't have any!",
			discordgo.Danish:       "We don't have any!",
			discordgo.Dutch:        "We don't have any!",
			discordgo.Finnish:      "We don't have any!",
			discordgo.French:       "We don't have any!",
			discordgo.German:       "We don't have any!",
			discordgo.Greek:        "We don't have any!",
			discordgo.Hungarian:    "We don't have any!",
			discordgo.Italian:      "We don't have any!",
			discordgo.Japanese:     "We don't have any!",
			discordgo.Korean:       "We don't have any!",
			discordgo.Lithuanian:   "We don't have any!",
			discordgo.Norwegian:    "We don't have any!",
			discordgo.Polish:       "We don't have any!",
			discordgo.PortugueseBR: "We don't have any!",
			discordgo.Romanian:     "We don't have any!",
			discordgo.Russian:      "We don't have any!",
			discordgo.SpanishES:    "We don't have any!",
			discordgo.Swedish:      "We don't have any!",
			discordgo.Ukrainian:    "We don't have any!",
			discordgo.Hindi:        "We don't have any!",
			discordgo.Thai:         "We don't have any!",
			discordgo.Turkish:      "We don't have any!",
			discordgo.Vietnamese:   "We don't have any!",
			discordgo.Unknown:      "We don't have any!",
		},
		{
			discordgo.EnglishGB:    "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.EnglishUS:    "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Bulgarian:    "Може да използвате този бот по какъвто начин искате, стига да уважавате неговият лиценз",
			discordgo.ChineseCN:    "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.ChineseTW:    "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Croatian:     "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Czech:        "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Danish:       "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Dutch:        "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Finnish:      "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.French:       "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.German:       "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Greek:        "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Hungarian:    "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Italian:      "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Japanese:     "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Korean:       "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Lithuanian:   "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Norwegian:    "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Polish:       "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.PortugueseBR: "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Romanian:     "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Russian:      "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.SpanishES:    "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Swedish:      "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Ukrainian:    "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Hindi:        "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Thai:         "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Turkish:      "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Vietnamese:   "Use this bot for whatever you like as long as you respect it's licence",
			discordgo.Unknown:      "Use this bot for whatever you like as long as you respect it's licence",
		},
	}
	embed := NewEmbed().
		SetTitle(embedStrings[0][m.Locale]).
		AddField(embedStrings[1][m.Locale], embedStrings[2][m.Locale]+": https://github.com/MadLadSquad/MadLadSquadBot/blob/master/LICENSE").
		SetFooter(footerTranslations[m.Locale], "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
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
	embedStrings := []map[discordgo.Locale]string{
		{
			discordgo.EnglishGB:    "About",
			discordgo.EnglishUS:    "About",
			discordgo.Bulgarian:    "За бота",
			discordgo.ChineseCN:    "About",
			discordgo.ChineseTW:    "About",
			discordgo.Croatian:     "About",
			discordgo.Czech:        "About",
			discordgo.Danish:       "About",
			discordgo.Dutch:        "About",
			discordgo.Finnish:      "About",
			discordgo.French:       "About",
			discordgo.German:       "About",
			discordgo.Greek:        "About",
			discordgo.Hungarian:    "About",
			discordgo.Italian:      "About",
			discordgo.Japanese:     "About",
			discordgo.Korean:       "About",
			discordgo.Lithuanian:   "About",
			discordgo.Norwegian:    "About",
			discordgo.Polish:       "About",
			discordgo.PortugueseBR: "About",
			discordgo.Romanian:     "About",
			discordgo.Russian:      "About",
			discordgo.SpanishES:    "About",
			discordgo.Swedish:      "About",
			discordgo.Ukrainian:    "About",
			discordgo.Hindi:        "About",
			discordgo.Thai:         "About",
			discordgo.Turkish:      "About",
			discordgo.Vietnamese:   "About",
			discordgo.Unknown:      "About",
		},
		{
			discordgo.EnglishGB:    "This bot is developed and maintained by MadLad Squad",
			discordgo.EnglishUS:    "This bot is developed and maintained by MadLad Squad",
			discordgo.Bulgarian:    "Този бот бива разработен и поддържан от MadLad Squad",
			discordgo.ChineseCN:    "This bot is developed and maintained by MadLad Squad",
			discordgo.ChineseTW:    "This bot is developed and maintained by MadLad Squad",
			discordgo.Croatian:     "This bot is developed and maintained by MadLad Squad",
			discordgo.Czech:        "This bot is developed and maintained by MadLad Squad",
			discordgo.Danish:       "This bot is developed and maintained by MadLad Squad",
			discordgo.Dutch:        "This bot is developed and maintained by MadLad Squad",
			discordgo.Finnish:      "This bot is developed and maintained by MadLad Squad",
			discordgo.French:       "This bot is developed and maintained by MadLad Squad",
			discordgo.German:       "This bot is developed and maintained by MadLad Squad",
			discordgo.Greek:        "This bot is developed and maintained by MadLad Squad",
			discordgo.Hungarian:    "This bot is developed and maintained by MadLad Squad",
			discordgo.Italian:      "This bot is developed and maintained by MadLad Squad",
			discordgo.Japanese:     "This bot is developed and maintained by MadLad Squad",
			discordgo.Korean:       "This bot is developed and maintained by MadLad Squad",
			discordgo.Lithuanian:   "This bot is developed and maintained by MadLad Squad",
			discordgo.Norwegian:    "This bot is developed and maintained by MadLad Squad",
			discordgo.Polish:       "This bot is developed and maintained by MadLad Squad",
			discordgo.PortugueseBR: "This bot is developed and maintained by MadLad Squad",
			discordgo.Romanian:     "This bot is developed and maintained by MadLad Squad",
			discordgo.Russian:      "This bot is developed and maintained by MadLad Squad",
			discordgo.SpanishES:    "This bot is developed and maintained by MadLad Squad",
			discordgo.Swedish:      "This bot is developed and maintained by MadLad Squad",
			discordgo.Ukrainian:    "This bot is developed and maintained by MadLad Squad",
			discordgo.Hindi:        "This bot is developed and maintained by MadLad Squad",
			discordgo.Thai:         "This bot is developed and maintained by MadLad Squad",
			discordgo.Turkish:      "This bot is developed and maintained by MadLad Squad",
			discordgo.Vietnamese:   "This bot is developed and maintained by MadLad Squad",
			discordgo.Unknown:      "This bot is developed and maintained by MadLad Squad",
		},
		{
			discordgo.EnglishGB:    "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.EnglishUS:    "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Bulgarian:    "Ако има проблеми с бота, искате да предложите нови команди или просто искате да говорите с нас, влезте в нашият discord сървър",
			discordgo.ChineseCN:    "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.ChineseTW:    "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Croatian:     "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Czech:        "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Danish:       "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Dutch:        "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Finnish:      "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.French:       "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.German:       "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Greek:        "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Hungarian:    "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Italian:      "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Japanese:     "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Korean:       "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Lithuanian:   "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Norwegian:    "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Polish:       "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.PortugueseBR: "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Romanian:     "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Russian:      "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.SpanishES:    "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Swedish:      "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Ukrainian:    "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Hindi:        "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Thai:         "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Turkish:      "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Vietnamese:   "If you have any problems, want to suggest a feature or just chat join our discord server",
			discordgo.Unknown:      "If you have any problems, want to suggest a feature or just chat join our discord server",
		},
		{
			discordgo.EnglishGB:    "If you want to submit bug reports report them here",
			discordgo.EnglishUS:    "If you want to submit bug reports report them here",
			discordgo.Bulgarian:    "Ако намирате бъгове и желаете да допринесете към тяхното оправяме, направете bug report тук",
			discordgo.ChineseCN:    "If you want to submit bug reports report them here",
			discordgo.ChineseTW:    "If you want to submit bug reports report them here",
			discordgo.Croatian:     "If you want to submit bug reports report them here",
			discordgo.Czech:        "If you want to submit bug reports report them here",
			discordgo.Danish:       "If you want to submit bug reports report them here",
			discordgo.Dutch:        "If you want to submit bug reports report them here",
			discordgo.Finnish:      "If you want to submit bug reports report them here",
			discordgo.French:       "If you want to submit bug reports report them here",
			discordgo.German:       "If you want to submit bug reports report them here",
			discordgo.Greek:        "If you want to submit bug reports report them here",
			discordgo.Hungarian:    "If you want to submit bug reports report them here",
			discordgo.Italian:      "If you want to submit bug reports report them here",
			discordgo.Japanese:     "If you want to submit bug reports report them here",
			discordgo.Korean:       "If you want to submit bug reports report them here",
			discordgo.Lithuanian:   "If you want to submit bug reports report them here",
			discordgo.Norwegian:    "If you want to submit bug reports report them here",
			discordgo.Polish:       "If you want to submit bug reports report them here",
			discordgo.PortugueseBR: "If you want to submit bug reports report them here",
			discordgo.Romanian:     "If you want to submit bug reports report them here",
			discordgo.Russian:      "If you want to submit bug reports report them here",
			discordgo.SpanishES:    "If you want to submit bug reports report them here",
			discordgo.Swedish:      "If you want to submit bug reports report them here",
			discordgo.Ukrainian:    "If you want to submit bug reports report them here",
			discordgo.Hindi:        "If you want to submit bug reports report them here",
			discordgo.Thai:         "If you want to submit bug reports report them here",
			discordgo.Turkish:      "If you want to submit bug reports report them here",
			discordgo.Vietnamese:   "If you want to submit bug reports report them here",
			discordgo.Unknown:      "If you want to submit bug reports report them here",
		},
	}
	embed := NewEmbed().
		SetTitle(embedStrings[0][m.Locale]).
		SetThumbnail("https://cdn.discordapp.com/avatars/697420452712284202/924342db89aa1f0acd5239646a835bec.png").
		AddField(embedStrings[1][m.Locale], "https://github.com/MadLadSquad/ , https://madladsquad.com/").
		AddField(embedStrings[2][m.Locale], "https://discord.com/invite/heA4FpTaTj").
		AddField(embedStrings[3][m.Locale], "https://github.com/MadLadSquad/MadLadSquadBot/issues").
		SetFooter(footerTranslations[m.Locale], "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
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
			SetFooter(footerTranslations[m.Locale], "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
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
			SetFooter(footerTranslations[m.Locale], "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
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
