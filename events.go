package main

import (
	"github.com/MadLadSquad/discordgo"
	"strconv"
	"strings"
)

func onReady(s *discordgo.Session, m *discordgo.Ready) {
	err := s.UpdateStreamingStatus(1, "&help | Click watch for music!", "https://www.youtube.com/watch?v=XoX6zS5-jOY&list=PL8yFU3veFghtteYYFdSnc-vaBZ3hHh4Wc&index=1")
	if err != nil {
		return
	}
}

func onChannelCreate(s *discordgo.Session, m *discordgo.ChannelCreate) {
	topic := "None"
	if m.Channel.Topic != "" {
		topic = m.Channel.Topic
	}

	channelType := "Text Channel"

	switch m.Channel.Type {
	case 0:
		channelType = "Text Channel"
		break
	case 2:
		channelType = "Voice Channel"
		break
	case 4:
		channelType = "Category"
		break
	case 5:
		channelType = "News"
		break
	case 6:
		channelType = "Store"
		break
	case 10:
		channelType = "News Thread"
		break
	case 11:
		channelType = "Public Thread"
		break
	}

	embed := NewEmbed().
		SetTitle("New channel created!").
		AddField("Channel", m.Channel.Name).
		InlineAllFields().
		AddField("ID", m.Channel.ID).
		AddField("NSFW", strconv.FormatBool(m.Channel.NSFW)).
		AddField("Topic", topic).
		AddField("Channel Type", channelType).
		InlineAllFields().
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	guild, _ := s.State.Guild(m.Channel.GuildID)

	for i := 0; i < len(guild.Channels); i++ {
		if strings.Contains(strings.ToLower(guild.Channels[i].Topic), "ubot-event-log") {
			_, _ = s.ChannelMessageSendEmbed(guild.Channels[i].ID, embed)
			return
		}
	}
}

func onChannelUpdate(s *discordgo.Session, m *discordgo.ChannelUpdate) {
	topic := "None"
	if m.Channel.Topic != "" {
		topic = m.Channel.Topic
	}

	channelType := "Text Channel"

	switch m.Channel.Type {
	case 0:
		channelType = "Text Channel"
		break
	case 2:
		channelType = "Voice Channel"
		break
	case 4:
		channelType = "Category"
		break
	case 5:
		channelType = "News"
		break
	case 6:
		channelType = "Store"
		break
	case 10:
		channelType = "News Thread"
		break
	case 11:
		channelType = "Public Thread"
		break
	}

	embed := NewEmbed().
		SetTitle("Channel was updated!").
		AddField("Channel", m.Channel.Name).
		AddField("ID", m.Channel.ID).
		AddField("NSFW", strconv.FormatBool(m.Channel.NSFW)).
		AddField("Channel Type", channelType).
		AddField("Topic", topic).
		InlineAllFields().
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	guild, _ := s.State.Guild(m.Channel.GuildID)

	for i := 0; i < len(guild.Channels); i++ {
		if strings.Contains(strings.ToLower(guild.Channels[i].Topic), "ubot-event-log") {
			_, _ = s.ChannelMessageSendEmbed(guild.Channels[i].ID, embed)
			return
		}
	}
}

func onChannelRemove(s *discordgo.Session, m *discordgo.ChannelDelete) {
	topic := "None"
	if m.Channel.Topic != "" {
		topic = m.Channel.Topic
	}

	channelType := "Text Channel"

	switch m.Channel.Type {
	case 0:
		channelType = "Text Channel"
		break
	case 2:
		channelType = "Voice Channel"
		break
	case 4:
		channelType = "Category"
		break
	case 5:
		channelType = "News"
		break
	case 6:
		channelType = "Store"
		break
	case 10:
		channelType = "News Thread"
		break
	case 11:
		channelType = "Public Thread"
		break
	}

	embed := NewEmbed().
		SetTitle("Channel was removed!").
		AddField("Channel", m.Channel.Name).
		InlineAllFields().
		AddField("ID", m.Channel.ID).
		AddField("NSFW", strconv.FormatBool(m.Channel.NSFW)).
		AddField("Channel Type", channelType).
		InlineAllFields().
		AddField("Topic", topic).
		InlineAllFields().
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	guild, _ := s.State.Guild(m.Channel.GuildID)

	for i := 0; i < len(guild.Channels); i++ {
		if strings.Contains(strings.ToLower(guild.Channels[i].Topic), "ubot-event-log") {
			_, _ = s.ChannelMessageSendEmbed(guild.Channels[i].ID, embed)
			return
		}
	}
}

func onPinUpdate(s *discordgo.Session, m *discordgo.ChannelPinsUpdate) {
	guild, _ := s.State.Guild(m.GuildID)
	channel, _ := s.Channel(m.ChannelID)

	embed := NewEmbed().
		SetTitle("Pins were updated!").
		AddField("Channel", channel.Mention()).
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	for i := 0; i < len(guild.Channels); i++ {
		if strings.Contains(strings.ToLower(guild.Channels[i].Topic), "ubot-event-log") {
			_, _ = s.ChannelMessageSendEmbed(guild.Channels[i].ID, embed)
			return
		}
	}
}

func onBanRemove(s *discordgo.Session, m *discordgo.GuildBanRemove) {
	guild, _ := s.State.Guild(m.GuildID)

	embed := NewEmbed().
		SetTitle("Ban was removed!").
		SetThumbnail(m.User.AvatarURL("")).
		AddField("User", m.User.Mention()).
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	for i := 0; i < len(guild.Channels); i++ {
		if strings.Contains(strings.ToLower(guild.Channels[i].Topic), "ubot-event-log") {
			_, _ = s.ChannelMessageSendEmbed(guild.Channels[i].ID, embed)
			return
		}
	}
}

func onGuildUpdate(s *discordgo.Session, m *discordgo.GuildUpdate) {
	owner, _ := s.User(m.Guild.OwnerID)

	mfaLevel := "None"
	premiumTier := "None"
	switch m.Guild.MfaLevel {
	case 0:
		mfaLevel = "None"
		break
	case 1:
		mfaLevel = "Elevated"
		break
	}

	switch m.Guild.PremiumTier {
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
	switch m.Guild.VerificationLevel {
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
	if len(m.Guild.Description) > 2 {
		//guildDescription = guild.Description
	}
	embed := NewEmbed().
		SetTitle(m.Guild.Name+" was updated!").
		SetThumbnail(m.Guild.IconURL()).
		AddField("Member Count", strconv.Itoa(len(m.Guild.Members))).
		AddField("Region", m.Guild.Region).
		AddField("Channel Count", strconv.Itoa(len(m.Guild.Channels))).
		InlineAllFields().
		AddField("Server ID", m.Guild.ID).
		AddField("Large", strconv.FormatBool(m.Guild.Large)).
		AddField("AFK Timeout", strconv.Itoa(m.Guild.AfkTimeout)+" seconds").
		InlineAllFields().
		AddField("Emoji Count", strconv.Itoa(len(m.Guild.Emojis))).
		AddField("Owner", owner.Mention()).
		AddField("Role Count", strconv.Itoa(len(m.Guild.Roles))).
		InlineAllFields().
		AddField("MFA Level", mfaLevel).
		AddField("Boost Level", premiumTier).
		AddField("Verification Level", verificationLevel).
		InlineAllFields().
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	for i := 0; i < len(m.Guild.Channels); i++ {
		if strings.Contains(strings.ToLower(m.Guild.Channels[i].Topic), "ubot-event-log") {
			_, _ = s.ChannelMessageSendEmbed(m.Guild.Channels[i].ID, embed)
			return
		}
	}
}

func onGuildMemberUpdate(s *discordgo.Session, m *discordgo.GuildMemberUpdate) {
	guild, _ := s.State.Guild(m.GuildID)

	premiumType := "None"
	switch m.User.PremiumType {
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
		SetTitle(m.User.Username+"'s profile was updated!").
		SetThumbnail(m.User.AvatarURL("")).
		AddField("Bot", strconv.FormatBool(m.User.Bot)).
		AddField("Mention", m.User.Mention()).
		AddField("Nitro type", premiumType).
		InlineAllFields().
		AddField("User ID", m.User.ID).
		AddField("System", strconv.FormatBool(m.User.System)).
		InlineAllFields().
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	for i := 0; i < len(guild.Channels); i++ {
		if strings.Contains(strings.ToLower(guild.Channels[i].Topic), "ubot-event-log") {
			_, _ = s.ChannelMessageSendEmbed(guild.Channels[i].ID, embed)
			return
		}
	}
}

func onRoleCreate(s *discordgo.Session, m *discordgo.GuildRoleCreate) {
	guild, _ := s.State.Guild(m.GuildID)

	embed := NewEmbed().
		SetTitle("A new role was created!").
		AddField("Role", m.GuildRole.Role.Mention()).
		AddField("Colour", "0x"+strconv.FormatInt(int64(m.GuildRole.Role.Color), 16)).
		InlineAllFields().
		AddField("ID", m.GuildRole.Role.ID).
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	for i := 0; i < len(guild.Channels); i++ {
		if strings.Contains(strings.ToLower(guild.Channels[i].Topic), "ubot-event-log") {
			_, _ = s.ChannelMessageSendEmbed(guild.Channels[i].ID, embed)
			return
		}
	}
}

func onRoleUpdate(s *discordgo.Session, m *discordgo.GuildRoleUpdate) {
	guild, _ := s.State.Guild(m.GuildID)

	embed := NewEmbed().
		SetTitle("A role was updated!").
		AddField("Role", m.GuildRole.Role.Mention()).
		AddField("Colour", "0x"+strconv.FormatInt(int64(m.GuildRole.Role.Color), 16)).
		InlineAllFields().
		AddField("ID", m.GuildRole.Role.ID).
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	for i := 0; i < len(guild.Channels); i++ {
		if strings.Contains(strings.ToLower(guild.Channels[i].Topic), "ubot-event-log") {
			_, _ = s.ChannelMessageSendEmbed(guild.Channels[i].ID, embed)
			return
		}
	}
}

func onRoleRemove(s *discordgo.Session, m *discordgo.GuildRoleDelete) {
	guild, _ := s.State.Guild(m.GuildID)

	embed := NewEmbed().
		SetTitle("A role was removed!").
		AddField("ID", m.RoleID).
		SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
		SetColor(0xf1c40f).MessageEmbed

	for i := 0; i < len(guild.Channels); i++ {
		if strings.Contains(strings.ToLower(guild.Channels[i].Topic), "ubot-event-log") {
			_, _ = s.ChannelMessageSendEmbed(guild.Channels[i].ID, embed)
			return
		}
	}
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate, message []string, content string) {
	channel, _ := s.Channel(m.ChannelID)
	redirect := ""
	f := 0
	bAlreadyChecked := false
	var aliasStrings [10]string
	var cmdStrings [10]string
	var currentAliasIndex = 0

checkAgain:
	if strings.Contains(content, prefix+"userinfo") || redirect == "userinfo" {
		showUserInfo(message[1], s, m)
	} else if strings.Contains(content, prefix+"serverinfo") || redirect == "serverinfo" {
		showServerInfo(s, m)
	} else if strings.Contains(content, prefix+"help") || redirect == "help" {
		help(s, m)
	} else if (strings.Contains(content, prefix+"kick") || redirect == "kick") && len(message) > 1 {
		kick(message[1], s, m)
	} else if (strings.Contains(content, prefix+"ban") || redirect == "ban") && len(message) > 1 {
		ban(message, s, m)
	} else if strings.Contains(content, prefix+"invite") || redirect == "invite" {
		invite(s, m)
	} else if strings.Contains(content, prefix+"privacy") || redirect == "privacy" {
		privacyPolicy(s, m)
	} else if strings.Contains(content, prefix+"tos") || redirect == "tos" {
		termsOfService(s, m)
	} else if strings.Contains(content, prefix+"about") || redirect == "about" {
		about(s, m)
	} else if strings.Contains(content, prefix+"verify") || redirect == "verify" {
		verify(s, m)
	} else if (strings.Contains(content, prefix+"avatar") || redirect == "avatar") && len(message) > 1 {
		avatar(message[1], s, m)
	} else if (strings.Contains(content, prefix+"mute") || redirect == "mute") && len(message) > 1 {
		mute(message[1], s, m)
	} else if strings.Contains(content, prefix+"sus") || redirect == "sus" {
		sus(s, m)
	} else if (strings.Contains(content, prefix+"set-welcome") || redirect == "set-welcome") && len(message) > 1 {
		channelChangeMetadata(message[1], s, m, " ubot-welcome", "welcome")
	} else if (strings.Contains(content, prefix+"set-event-tracking") || redirect == "set-event-tracking") && len(message) > 1 {
		channelChangeMetadata(message[1], s, m, " ubot-event-log", "event logging")
	} else if (strings.Contains(content, prefix+"set-text-only") || redirect == "set-text-only") && len(message) > 1 {
		channelChangeMetadata(message[1], s, m, " ubot-restrict-text-only", "text only")
	} else if (strings.Contains(content, prefix+"set-attachments-only") || redirect == "set-attachments-only") && len(message) > 1 {
		channelChangeMetadata(message[1], s, m, " ubot-restrict-attachments-only", "attachments only")
	} else if (strings.Contains(content, prefix+"set-links-only") || redirect == "set-links-only") && len(message) > 1 {
		channelChangeMetadata(message[1], s, m, " ubot-restrict-links-only", "links only")
	} else if strings.Contains(content, prefix+"generate-member-role") || redirect == "generate-member-role" {
		createMemberRole(s, m)
	} else if strings.Contains(content, prefix+"pernik") || redirect == "pernik" {
		pernik(s, m)
	} else if (strings.Contains(content, prefix+"set-colour-role-channel") || redirect == "set-colour-role-channel") && len(message) > 1 && message[1] != "" {
		channelChangeMetadata(message[1], s, m, " ubot-colour-pick", "colour role")
	} else if (strings.Contains(content, prefix+"set-colour-role") || redirect == "set-colour-role") && len(message) > 1 && message[1] != "" {
		giveColour(message[1], s, m)
	} else if (strings.Contains(content, prefix+"set-meta-role-channel") || redirect == "set-meta-role-channel") && len(message) > 1 {
		channelChangeMetadata(message[1], s, m, "ubot-meta-role-pick", "meta role")
	} else if (strings.Contains(content, prefix+"set-meta-role") || redirect == "set-meta-role") && len(message) > 1 && message[1] != "" {
		giveMetarole(message[1], s, m)
	} else if strings.Contains(content, prefix+"list-colour-roles") || redirect == "list-colour-roles" {
		listColours(s, m)
	} else if strings.Contains(content, prefix+"list-aliases") || redirect == "list-aliases" {
		listAliases(s, m)
	} else if (strings.Contains(content, prefix+"remove-meta-role") || redirect == "remove-meta-role") && len(message) > 1 {
		removeMetarole(message[1], s, m)
	} else if strings.Contains(content, prefix+"alias-help") || redirect == "alias-help" {
		aliasHelp(s, m)
	} else if strings.Contains(strings.ToLower(channel.Topic), "ubot-macro:") {
		if !bAlreadyChecked {
			topic := strings.ToLower(channel.Topic)
			i := strings.Index(topic, "ubot-macro:")

			if i != -1 {
				i += len("ubot-macro:")
			} else {
				return
			}

			bRecordingAlias := true
			for j := i; j < len(topic); j++ {
				if bRecordingAlias {
					if topic[j] == '>' {
						bRecordingAlias = false
					} else {
						aliasStrings[currentAliasIndex] += string(topic[j])
					}
				} else {
					if topic[j] == ';' && currentAliasIndex < 10 {
						bRecordingAlias = true
						currentAliasIndex++
					} else if topic[j] == ']' || topic[j] == '[' || currentAliasIndex == 10 {
						break
					} else {
						cmdStrings[currentAliasIndex] += string(topic[j])
					}
				}
			}
			bAlreadyChecked = true
		}

		for ; f <= currentAliasIndex; f++ {
			if strings.Contains(content, prefix+aliasStrings[f]) {
				redirect = strings.ToLower(cmdStrings[f])

				goto checkAgain
			}
		}
	}
}
