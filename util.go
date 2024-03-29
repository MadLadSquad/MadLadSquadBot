package main

import (
	"github.com/MadLadSquad/discordgo"
	"strings"
)

var (
	pingReplacer        = strings.NewReplacer("<", "", "#", "", ">", "", "@", "", "!", "")
	channelTranslations = map[discordgo.Locale]string{
		discordgo.EnglishGB:    "Channel",
		discordgo.EnglishUS:    "Channel",
		discordgo.Bulgarian:    "Канал",
		discordgo.ChineseCN:    "Channel",
		discordgo.ChineseTW:    "Channel",
		discordgo.Croatian:     "Channel",
		discordgo.Czech:        "Channel",
		discordgo.Danish:       "Channel",
		discordgo.Dutch:        "Channel",
		discordgo.Finnish:      "Channel",
		discordgo.French:       "Channel",
		discordgo.German:       "Channel",
		discordgo.Greek:        "Channel",
		discordgo.Hungarian:    "Channel",
		discordgo.Italian:      "Channel",
		discordgo.Japanese:     "Channel",
		discordgo.Korean:       "Channel",
		discordgo.Lithuanian:   "Channel",
		discordgo.Norwegian:    "Channel",
		discordgo.Polish:       "Channel",
		discordgo.PortugueseBR: "Channel",
		discordgo.Romanian:     "Channel",
		discordgo.Russian:      "Channel",
		discordgo.SpanishES:    "Channel",
		discordgo.Swedish:      "Channel",
		discordgo.Ukrainian:    "Channel",
		discordgo.Hindi:        "Channel",
		discordgo.Thai:         "Channel",
		discordgo.Turkish:      "Channel",
		discordgo.Vietnamese:   "Channel",
		discordgo.Unknown:      "Channel",
	}
)

func sanitizePings(str string) string {
	return pingReplacer.Replace(str)
}

func checkPerm(s *discordgo.Session, m *discordgo.MessageCreate, perm int64) bool {
	member, err := s.State.Member(m.GuildID, m.Author.ID)
	if err != nil {
		member, err = s.GuildMember(m.GuildID, m.Author.ID)
		if err != nil {
			return false
		}
	}

	for _, roleID := range member.Roles {
		role, err := s.State.Role(m.GuildID, roleID)
		if err != nil {
			return false
		}
		if role.Permissions&perm != 0 {
			return true
		}
	}

	return false
}

func channelChangeMetadata(arg string, s *discordgo.Session, m *discordgo.InteractionCreate, template1 string, template2 string) {
	if arg != "" {
		channel, _ := s.Channel(sanitizePings(arg))

		e := discordgo.ChannelEdit{
			Name:                 channel.Name,
			Topic:                channel.Topic + template1,
			NSFW:                 &channel.NSFW,
			Position:             &channel.Position,
			Bitrate:              channel.Bitrate,
			UserLimit:            channel.UserLimit,
			PermissionOverwrites: channel.PermissionOverwrites,
			ParentID:             channel.ParentID,
			RateLimitPerUser:     &channel.RateLimitPerUser,
		}

		_, _ = s.ChannelEditComplex(channel.ID, &e)

		embed := NewEmbed().
			SetTitle("Set a channel as a "+template2+" channel!").
			AddField(channelTranslations[m.Locale], channel.Mention()).
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
}

func truncateString(str string, ln int) string {
	runeSlice := []rune(str)

	if len(runeSlice) > ln {
		return string(runeSlice[:ln])
	}
	return str
}
