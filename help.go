package main

import (
	"fmt"
	"github.com/MadLadSquad/discordgo"
)

var (
	footerTranslations = map[discordgo.Locale]string{
		discordgo.EnglishGB:    "Message delivered using UntitledTechnology",
		discordgo.EnglishUS:    "Message delivered using UntitledTechnology",
		discordgo.Bulgarian:    "Това съобщение стига до вас, благодарение на Untitled технологии",
		discordgo.ChineseCN:    "Message delivered using UntitledTechnology",
		discordgo.ChineseTW:    "Message delivered using UntitledTechnology",
		discordgo.Croatian:     "Message delivered using UntitledTechnology",
		discordgo.Czech:        "Message delivered using UntitledTechnology",
		discordgo.Danish:       "Message delivered using UntitledTechnology",
		discordgo.Dutch:        "Message delivered using UntitledTechnology",
		discordgo.Finnish:      "Message delivered using UntitledTechnology",
		discordgo.French:       "Message delivered using UntitledTechnology",
		discordgo.German:       "Message delivered using UntitledTechnology",
		discordgo.Greek:        "Message delivered using UntitledTechnology",
		discordgo.Hungarian:    "Message delivered using UntitledTechnology",
		discordgo.Italian:      "Message delivered using UntitledTechnology",
		discordgo.Japanese:     "Message delivered using UntitledTechnology",
		discordgo.Korean:       "Message delivered using UntitledTechnology",
		discordgo.Lithuanian:   "Message delivered using UntitledTechnology",
		discordgo.Norwegian:    "Message delivered using UntitledTechnology",
		discordgo.Polish:       "Message delivered using UntitledTechnology",
		discordgo.PortugueseBR: "Message delivered using UntitledTechnology",
		discordgo.Romanian:     "Message delivered using UntitledTechnology",
		discordgo.Russian:      "Message delivered using UntitledTechnology",
		discordgo.SpanishES:    "Message delivered using UntitledTechnology",
		discordgo.Swedish:      "Message delivered using UntitledTechnology",
		discordgo.Ukrainian:    "Message delivered using UntitledTechnology",
		discordgo.Hindi:        "Message delivered using UntitledTechnology",
		discordgo.Thai:         "Message delivered using UntitledTechnology",
		discordgo.Turkish:      "Message delivered using UntitledTechnology",
		discordgo.Vietnamese:   "Message delivered using UntitledTechnology",
		discordgo.Unknown:      "Message delivered using UntitledTechnology",
	}
)

func createInvite(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "invite",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Returns a link for inviting the bot",
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB:    "invite",
			discordgo.EnglishUS:    "invite",
			discordgo.Bulgarian:    "покани",
			discordgo.ChineseCN:    "邀请",
			discordgo.ChineseTW:    "邀請",
			discordgo.Croatian:     "pozovi",
			discordgo.Czech:        "pozvat",
			discordgo.Danish:       "inviter",
			discordgo.Dutch:        "uitnodigen",
			discordgo.Finnish:      "kutsu",
			discordgo.French:       "inviter",
			discordgo.German:       "einladen",
			discordgo.Greek:        "προσκάλεσε",
			discordgo.Hungarian:    "meghívása",
			discordgo.Italian:      "invita",
			discordgo.Japanese:     "招待",
			discordgo.Korean:       "초대하기",
			discordgo.Lithuanian:   "kviesti",
			discordgo.Norwegian:    "inviter",
			discordgo.Polish:       "zaproś",
			discordgo.PortugueseBR: "convidar",
			discordgo.Romanian:     "invită",
			discordgo.Russian:      "пригласить",
			discordgo.SpanishES:    "invitar",
			discordgo.Swedish:      "bjud",
			discordgo.Ukrainian:    "запросити",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "Returns a link for inviting the bot",
			discordgo.EnglishUS: "Returns a link for inviting the bot",
			discordgo.Bulgarian: "Дава линк за покана на бота в друг сървър",
		},
	}
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
	if err != nil {
		fmt.Println(err)
	}
}

func invite(s *discordgo.Session, m *discordgo.InteractionCreate) {
	//embedStrings := []string{"Invite", "Use this link", "Message delivered using UntitledTechnology"}
	embedStrings := []map[discordgo.Locale]string{
		{
			discordgo.EnglishGB:    "Invite",
			discordgo.EnglishUS:    "Invite",
			discordgo.Bulgarian:    "Покани",
			discordgo.ChineseCN:    "邀请",
			discordgo.ChineseTW:    "邀請",
			discordgo.Croatian:     "Pozovi",
			discordgo.Czech:        "Pozvat",
			discordgo.Danish:       "Inviter",
			discordgo.Dutch:        "Uitnodigen",
			discordgo.Finnish:      "Kutsu",
			discordgo.French:       "Inviter",
			discordgo.German:       "Einladen",
			discordgo.Greek:        "Пροσκάλεσε",
			discordgo.Hungarian:    "Meghívása",
			discordgo.Italian:      "Invita",
			discordgo.Japanese:     "招待",
			discordgo.Korean:       "초대하기",
			discordgo.Lithuanian:   "Kviesti",
			discordgo.Norwegian:    "Inviter",
			discordgo.Polish:       "Zaproś",
			discordgo.PortugueseBR: "Convidar",
			discordgo.Romanian:     "Invită",
			discordgo.Russian:      "Пригласить",
			discordgo.SpanishES:    "Invitar",
			discordgo.Swedish:      "Bjud",
			discordgo.Ukrainian:    "Запросити",
			discordgo.Hindi:        "Invite",
			discordgo.Thai:         "Invite",
			discordgo.Turkish:      "Invite",
			discordgo.Vietnamese:   "Invite",
			discordgo.Unknown:      "Invite",
		},
		{
			discordgo.EnglishGB:    "Use this link",
			discordgo.EnglishUS:    "Use this link",
			discordgo.Bulgarian:    "Ползвайте този линк",
			discordgo.ChineseCN:    "Use this link",
			discordgo.ChineseTW:    "Use this link",
			discordgo.Croatian:     "Use this link",
			discordgo.Czech:        "Use this link",
			discordgo.Danish:       "Use this link",
			discordgo.Dutch:        "Use this link",
			discordgo.Finnish:      "Use this link",
			discordgo.French:       "Use this link",
			discordgo.German:       "Use this link",
			discordgo.Greek:        "Use this link",
			discordgo.Hungarian:    "Use this link",
			discordgo.Italian:      "Use this link",
			discordgo.Japanese:     "Use this link",
			discordgo.Korean:       "Use this link",
			discordgo.Lithuanian:   "Use this link",
			discordgo.Norwegian:    "Use this link",
			discordgo.Polish:       "Use this link",
			discordgo.PortugueseBR: "Use this link",
			discordgo.Romanian:     "Use this link",
			discordgo.Russian:      "Use this link",
			discordgo.SpanishES:    "Use this link",
			discordgo.Swedish:      "Use this link",
			discordgo.Ukrainian:    "Use this link",
			discordgo.Hindi:        "Use this link",
			discordgo.Thai:         "Use this link",
			discordgo.Turkish:      "Use this link",
			discordgo.Vietnamese:   "Use this link",
			discordgo.Unknown:      "Use this link",
		},
	}

	embed := NewEmbed().
		SetTitle(embedStrings[0][m.Locale]).
		AddField(embedStrings[1][m.Locale], "https://discord.com/oauth2/authorize?client_id=697420452712284202&scope=bot&permissions=2134207679").
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
