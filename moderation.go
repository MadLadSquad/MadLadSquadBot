package main

import (
	"github.com/MadLadSquad/discordgo"
	"strings"
)

func createVerify(s *discordgo.Session) {
	dmPerm := false

	command := &discordgo.ApplicationCommand{
		Name:         "verify",
		Type:         discordgo.ChatApplicationCommand,
		Description:  "Allows you to verify yourself",
		DMPermission: &dmPerm,
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "verify",
			discordgo.EnglishUS: "verify",
			discordgo.Bulgarian: "верификация",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "Allows you to verify yourself",
			discordgo.EnglishUS: "Allows you to verify yourself",
			discordgo.Bulgarian: "Позволява ви да се верифицирате в даден сървър",
		},
	}
	_, _ = s.ApplicationCommandCreate(s.State.User.ID, "", command)
}

func createMember(s *discordgo.Session) {
	dmPerm := false
	command := &discordgo.ApplicationCommand{
		Name:         "generate-member-role",
		Type:         discordgo.ChatApplicationCommand,
		Description:  "Creates a Member role for the server",
		DMPermission: &dmPerm,
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "generate-member-role",
			discordgo.EnglishUS: "generate-member-role",
			discordgo.Bulgarian: "генерирай-роля-за-член",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "Creates a Member role for the server",
			discordgo.EnglishUS: "Creates a Member role for the server",
			discordgo.Bulgarian: "Създава базова роля за нови членове, за която се получава след верификация",
		},
	}
	_, _ = s.ApplicationCommandCreate(s.State.User.ID, "", command)
}

func verify(s *discordgo.Session, m *discordgo.InteractionCreate) {
	guild, _ := s.State.Guild(m.GuildID)

	for j := 0; j < len(guild.Roles); j++ {
		if strings.ToLower(guild.Roles[j].Name) == "member" || strings.ToLower(guild.Roles[j].Name) == "members" {
			_ = s.GuildMemberRoleAdd(m.GuildID, m.Member.User.ID, guild.Roles[j].ID)
			_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
			break
		}
	}

	_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: "You're verified!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func createMemberRole(s *discordgo.Session, m *discordgo.InteractionCreate) {
	roleinfo := discordgo.RoleParams{}
	roleinfo.Name = "Members"

	_, _ = s.GuildRoleCreate(m.GuildID, &roleinfo)
	_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: "Successfully created member role!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
