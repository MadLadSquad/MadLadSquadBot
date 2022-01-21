package main

import (
	"github.com/MadLadSquad/discordgo"
)

func sanitizePings(str string) string {
	var newStr string
	for i := 0; i < len(str); i++ {
		if !(str[i] == '<' || str[i] == '#' || str[i] == '>' || str[i] == '@' || str[i] == '!') {
			newStr += string(str[i])
		}
	}

	return newStr
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

func parseMessage(str string) [103]string {
	ret := [103]string{}

	index := 0

	for i := 0; i < len(str); i++ {
		if index < 103 {
			if str[i] == ' ' {
				index++
			} else {
				ret[index] += string(str[i])
			}
		} else {
			break
		}
	}

	return ret
}

func channelChangeMetadata(arg string, s *discordgo.Session, m *discordgo.MessageCreate, template1 string, template2 string) {
	if arg != "" {
		channel, _ := s.Channel(sanitizePings(arg))

		e := discordgo.ChannelEdit{
			Name:                 channel.Name,
			Topic:                channel.Topic + template1,
			NSFW:                 channel.NSFW,
			Position:             channel.Position,
			Bitrate:              channel.Bitrate,
			UserLimit:            channel.UserLimit,
			PermissionOverwrites: channel.PermissionOverwrites,
			ParentID:             channel.ParentID,
			RateLimitPerUser:     channel.RateLimitPerUser,
		}

		s.ChannelEditComplex(channel.ID, &e)

		embed := NewEmbed().
			SetTitle("Set a channel as a "+template2+" channel!").
			AddField("Channel", channel.Mention()).
			SetFooter("Message delivered using Untitled Technology", "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
			SetColor(0xf1c40f).MessageEmbed

		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}
