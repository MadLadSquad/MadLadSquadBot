package main

import (
	"github.com/MadLadSquad/discordgo"
	"strconv"
	"strings"
)

func onReady(s *discordgo.Session, m *discordgo.Ready) {
	err := s.UpdateStreamingStatus(1, "Type / for commands | Click watch for music!", "https://open.spotify.com/playlist/0UA5R74EvBpUJdNWMokC1y")
	if err != nil {
		return
	}
	createApplicationCommands(s)
}

func createApplicationCommands(s *discordgo.Session) {
	// Information commands
	createUserInfo(s)
	createServerInfo(s)
	createPrivacy(s)
	createTos(s)
	createAbout(s)
	createAvatar(s)
	createInvite(s)

	// Verify commands
	createVerify(s)
	createMember(s)

	// Meme commands
	createSus(s)

	// Colour roles
	createGiveColour(s)
	createListColourRoles(s)

	// Meta<role/data>
	createMetaRole(s)
	createSetChannel(s)

	// Alias commands
	createAlias(s)

	// Removes unneeded commands
	removeLegacy(s)
}

func removeLegacy(s *discordgo.Session) {
	legacyCommands := []string{"alias-help", "list-aliases", "set-colour-role", "remove-meta-role", "pernik"}
	cmds, _ := s.ApplicationCommands(s.State.User.ID, "")

	for i := 0; i < len(cmds); i++ {
		for j := 0; j < len(legacyCommands); j++ {
			if strings.ToLower(cmds[i].Name) == legacyCommands[j] {
				_ = s.ApplicationCommandDelete(s.State.User.ID, "", cmds[i].ID)
			}
		}
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
	case 10:
		channelType = "News Thread"
		break
	case 11:
		channelType = "Public Thread"
		break
	case 13:
		channelType = "Stage"
		break
	case 14:
		channelType = "Directory"
		break
	case 15:
		channelType = "Forum"
		break
	default:
		channelType = "Unknown"
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
	case 10:
		channelType = "News Thread"
		break
	case 11:
		channelType = "Public Thread"
		break
	case 13:
		channelType = "Stage"
		break
	case 14:
		channelType = "Directory"
		break
	case 15:
		channelType = "Forum"
		break
	default:
		channelType = "Unknown"
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
	case 10:
		channelType = "News Thread"
		break
	case 11:
		channelType = "Public Thread"
		break
	case 13:
		channelType = "Stage"
		break
	case 14:
		channelType = "Directory"
		break
	case 15:
		channelType = "Forum"
		break
	default:
		channelType = "Unknown"
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

	nsfwLevel := "Default"
	switch m.Guild.NSFWLevel {
	case 0:
		nsfwLevel = "Default"
		break
	case 1:
		nsfwLevel = "Explicit"
		break
	case 2:
		nsfwLevel = "Safe"
		break
	case 3:
		nsfwLevel = "Age Restricted"
		break
	}

	explicitContentFilter := "Disabled"
	switch m.Guild.ExplicitContentFilter {
	case 0:
		explicitContentFilter = "None"
		break
	case 1:
		explicitContentFilter = "Members with no roles"
		break
	case 2:
		explicitContentFilter = "All members"
		break
	}

	defaultNotificationLevel := "All Messages"
	if m.Guild.DefaultMessageNotifications == 1 {
		defaultNotificationLevel = "Only Mentions"
	}

	guildDescription := "None"
	if len(m.Guild.Description) > 2 {
		guildDescription = m.Guild.Description
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
		AddField("NSFW Level", nsfwLevel).
		AddField("Explicit content filter level", explicitContentFilter).
		AddField("Sticker Count", strconv.Itoa(len(m.Guild.Stickers))).
		InlineAllFields().
		AddField("Boost Count", strconv.Itoa(m.Guild.PremiumSubscriptionCount)).
		AddField("Default Notification Level", defaultNotificationLevel).
		InlineAllFields().
		AddField("Description", guildDescription).
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

	userFlags := ""
	if (m.User.PublicFlags & discordgo.UserFlagDiscordEmployee) != 0 {
		userFlags += "Staff, "
	}
	if (m.User.PublicFlags & discordgo.UserFlagDiscordPartner) != 0 {
		userFlags += "Partner, "
	}
	if (m.User.PublicFlags & discordgo.UserFlagHypeSquadEvents) != 0 {
		userFlags += "Hype Squad, "
	}
	if (m.User.PublicFlags & discordgo.UserFlagBugHunterLevel1) != 0 {
		userFlags += "Bug Hunter 1, "
	}
	if (m.User.PublicFlags & discordgo.UserFlagBugHunterLevel2) != 0 {
		userFlags += "Bug Hunter 2, "
	}
	if (m.User.PublicFlags & discordgo.UserFlagHouseBravery) != 0 {
		userFlags += "Hype Squad House Bravery, "
	}
	if (m.User.PublicFlags & discordgo.UserFlagHouseBrilliance) != 0 {
		userFlags += "Hype Squad House Brilliance, "
	}
	if (m.User.PublicFlags & discordgo.UserFlagHouseBalance) != 0 {
		userFlags += "Hype Squad House Balance, "
	}
	if (m.User.PublicFlags & discordgo.UserFlagEarlySupporter) != 0 {
		userFlags += "Nitro Early Supporter, "
	}
	if (m.User.PublicFlags & discordgo.UserFlagVerifiedBot) != 0 {
		userFlags += "Verified Bot, "
	}
	if (m.User.PublicFlags & discordgo.UserFlagVerifiedBotDeveloper) != 0 {
		userFlags += "Verified Bot Developer, "
	}
	if (m.User.PublicFlags & discordgo.UserFlagDiscordCertifiedModerator) != 0 {
		userFlags += "Certified Moderator, "
	}

	if userFlags == "" {
		userFlags = "None"
	} else {
		// Truncate the last 2 symbols if not empty i.e. ',' and ' '
		userFlags = userFlags[:2]
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
		AddField("User flags", userFlags).
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

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.ToLower(m.Content)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Author.Bot {
		return
	}

	channel, _ := s.State.Channel(m.ChannelID)
	if strings.Contains(strings.ToLower(channel.Topic), "ubot-restrict-text-only") && (len(m.Attachments) > 0 || strings.Contains(content, "http://") || strings.Contains(content, "https://")) {
		_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Message deleted due to it containing links or attachments in a text only channel!")
	} else if strings.Contains(strings.ToLower(channel.Topic), "ubot-restrict-attachments-only") && len(m.Attachments) == 0 {
		_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Message deleted due to it containing no attachments in an attachment only channel!")
	} else if strings.Contains(strings.ToLower(channel.Topic), "ubot-restrict-links-only") && !(strings.Contains(content, "http://") || strings.Contains(content, "https://")) {
		_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Message deleted due to it containing no links in link only channel!")
	}
}
