package main

import "github.com/MadLadSquad/discordgo"

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
		member, err = s.GuildMember(m.GuildID, m.Author.ID);
		if err != nil {
			return false
		}
	}

	for _, roleID := range member.Roles {
		role, err := s.State.Role(m.GuildID, roleID)
		if err != nil {
			return false
		}
		if role.Permissions & perm != 0 {
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