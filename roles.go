package main

import (
	"fmt"
	"github.com/MadLadSquad/discordgo"
	"strconv"
	"strings"
)

var (
	coloursHex = []uint64{
		0xf0f8ff,
		0xfaebd7,
		0x00ffff,
		0x7fffd4,
		0xf0ffff,
		0xf5f5dc,
		0xffe4c4,
		0x000000,
		0xffebcd,
		0x0000ff,
		0x8a2be2,
		0xa52a2a,
		0xdeb887,
		0x5f9ea0,
		0x7fff00,
		0xd2691e,
		0xff7f50,
		0x6495ed,
		0xfff8dc,
		0xdc143c,
		0x00ffff,
		0x00008b,
		0x008b8b,
		0xb8860b,
		0xa9a9a9,
		0xbdb76b,
		0x8b008b,
		0x556b2f,
		0xff8c00,
		0x9932cc,
		0x8b0000,
		0xe9967a,
		0x8fbc8f,
		0x483d8b,
		0x2f4f4f,
		0x00ced1,
		0x9400d3,
		0xff1493,
		0x00bfff,
		0x696969,
		0x1e90ff,
		0xb22222,
		0xfffaf0,
		0x228b22,
		0xff00ff,
		0xdcdcdc,
		0xf8f8ff,
		0xffd700,
		0xdaa520,
		0x808080,
		0x008000,
		0xadff2f,
		0xf0fff0,
		0xff69b4,
		0xcd5c5c,
		0x4b0082,
		0xfffff0,
		0xf0e68c,
		0xe6e6fa,
		0xfff0f5,
		0x7cfc00,
		0xfffacd,
		0xadd8e6,
		0xf08080,
		0xe0ffff,
		0xfafad2,
		0xd3d3d3,
		0x90ee90,
		0xffb6c1,
		0xffa07a,
		0x20b2aa,
		0x87cefa,
		0x778899,
		0xb0c4de,
		0xffffe0,
		0x00ff00,
		0x32cd32,
		0xfaf0e6,
		0xff00ff,
		0x7f0000,
		0x66cdaa,
		0x0000cd,
		0xba55d3,
		0x9370db,
		0x3cb371,
		0x7b68ee,
		0x00fa9a,
		0x48d1cc,
		0xc71585,
		0x191970,
		0xf5fffa,
		0xffe4e1,
		0xffe4b5,
		0xffdead,
		0x000080,
		0xfdf5e6,
		0x808000,
		0x6b8e23,
		0xffa500,
		0xff4500,
		0xda70d6,
		0xeee8aa,
		0x98fb98,
		0xafeeee,
		0xdb7093,
		0xffefd5,
		0xffdab9,
		0xcd853f,
		0xffc0cb,
		0xdda0dd,
		0xb0e0e6,
		0x7f007f,
		0xff0000,
		0xbc8f8f,
		0x4169e1,
		0x8b4513,
		0xfa8072,
		0xf4a460,
		0x2e8b57,
		0xfff5ee,
		0xa0522d,
		0xc0c0c0,
		0x87ceeb,
		0x6a5acd,
		0x708090,
		0xfffafa,
		0x00ff7f,
		0x4682b4,
		0xd2b48c,
		0x008080,
		0xd8bfd8,
		0xff6347,
		0x40e0d0,
		0xee82ee,
		0xf5deb3,
		0xffffff,
		0xf5f5f5,
		0xffff00,
		0x9acd32,
	}

	colours = []string{
		"Alice-Blue",
		"Antique-White",
		"Aqua",
		"Aquamarine",
		"Azure-Mist",
		"Beige",
		"Bisque",
		"Black",
		"Blanched-Almond",
		"Blue",
		"Blue-violet",
		"Auburn",
		"Burlywood",
		"Cadet-Blue",
		"Chartreuse",
		"Cinnamon",
		"Coral",
		"Cornflower-Blue",
		"Cornsilk",
		"Crimson",
		"Cyan",
		"Dark-Blue",
		"Dark-Cyan",
		"Dark-Goldenrod",
		"Dark-Gray",
		"Dark-Khaki",
		"Dark-Magenta",
		"Dark-Olive-Green",
		"Dark-Orange",
		"Dark-Orchid",
		"Dark-Red",
		"Dark-Salmon",
		"Dark-Sea-Green",
		"Dark-Slate-Blue",
		"Dark-Slate-Gray",
		"Dark-Turquoise",
		"Dark-Violet",
		"Deep-Pink",
		"Capri",
		"Dim-Gray",
		"Dodger-Blue",
		"Firebrick",
		"Floral-White",
		"Forest-Green",
		"Magenta",
		"Gainsboro",
		"Ghost-White",
		"Gold",
		"Goldenrod",
		"Gray",
		"Office-Green",
		"Green-yellow",
		"Honeydew",
		"Hot-Pink",
		"Chestnut",
		"Indigo",
		"Ivory",
		"Khaki",
		"Lavender",
		"Lavender-Blush",
		"Lawn-Green",
		"Lemon-Chiffon",
		"Light-Blue",
		"Light-Coral",
		"Light-Cyan",
		"Light-Goldenrod-Yellow",
		"Light-Gray",
		"Light-Green",
		"Light-Pink",
		"Light-Salmon",
		"Light-Sea-Green",
		"Light-Sky-Blue",
		"Light-Slate-Gray",
		"Light-Steel-Blue",
		"Light-Yellow",
		"Lime",
		"Lime-Green",
		"Linen",
		"Fuchsia",
		"Wood-Red",
		"Tropical-Green",
		"Medium-Blue",
		"Medium-Orchid",
		"Medium-Purple",
		"Medium-Sea-Green",
		"Medium-Slate-Blue",
		"Medium-Spring-Green",
		"Medium-Turquoise",
		"Red-violet",
		"Midnight-Blue",
		"Mint-Cream",
		"Misty-Rose",
		"SandySahara",
		"Navajo-White",
		"Navy-Blue",
		"Old-Lace",
		"Olive",
		"Olive-Drab",
		"Orange",
		"Orange-red",
		"Orchid",
		"Pale-Goldenrod",
		"Pale-Green",
		"Pale-Blue",
		"Pale-Violet-red",
		"Papaya-Whip",
		"Peach-Puff",
		"Peru",
		"Pink",
		"Plum",
		"Powder-Blue",
		"Sienna",
		"Red",
		"Rosy-Brown",
		"Royal-Blue",
		"Saddle-Brown",
		"Salmon",
		"Sandy-Brown",
		"Sea-Green",
		"Seashell",
		"Dark-Sand",
		"Silver",
		"Sky-Blue",
		"Slate-Blue",
		"Slate-Gray",
		"Snow",
		"Spring Green",
		"Steel-Blue",
		"Tan",
		"Teal",
		"Thistle",
		"Tomato",
		"Light-Cyan-2",
		"Violet",
		"Wheat",
		"White",
		"White-Smoke",
		"Electric-Yellow",
		"Yellow-green",
	}
	roleName = map[discordgo.Locale]string{
		discordgo.EnglishGB:    "Role",
		discordgo.EnglishUS:    "Role",
		discordgo.Bulgarian:    "Роля",
		discordgo.ChineseCN:    "Role",
		discordgo.ChineseTW:    "Role",
		discordgo.Croatian:     "Role",
		discordgo.Czech:        "Role",
		discordgo.Danish:       "Role",
		discordgo.Dutch:        "Role",
		discordgo.Finnish:      "Role",
		discordgo.French:       "Role",
		discordgo.German:       "Role",
		discordgo.Greek:        "Role",
		discordgo.Hungarian:    "Role",
		discordgo.Italian:      "Role",
		discordgo.Japanese:     "Role",
		discordgo.Korean:       "Role",
		discordgo.Lithuanian:   "Role",
		discordgo.Norwegian:    "Role",
		discordgo.Polish:       "Role",
		discordgo.PortugueseBR: "Role",
		discordgo.Romanian:     "Role",
		discordgo.Russian:      "Role",
		discordgo.SpanishES:    "Role",
		discordgo.Swedish:      "Role",
		discordgo.Ukrainian:    "Role",
		discordgo.Hindi:        "Role",
		discordgo.Thai:         "Role",
		discordgo.Turkish:      "Role",
		discordgo.Vietnamese:   "Role",
		discordgo.Unknown:      "Role",
	}
)

func createGiveColour(s *discordgo.Session) {
	var choices []*discordgo.ApplicationCommandOptionChoice

	for i := 0; i < len(colours); i++ {
		_ = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  strings.ToLower(colours[i]),
			Value: strconv.Itoa(i),
		})
	}

	command := &discordgo.ApplicationCommand{
		Name:        "set-colour-role",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Gives the user a colour role",
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "set-colour-role",
			discordgo.EnglishUS: "set-color-role",
			discordgo.Bulgarian: "сложи-роля-за-цвят",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "Gives the user a colour role",
			discordgo.EnglishUS: "Gives the user a color role",
			discordgo.Bulgarian: "Дава на потребителя роля с даден цвят",
		},
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "colour",
				Description: "The name of your colour, run the /list-colour-roles command for more info",
				Type:        discordgo.ApplicationCommandOptionString,
				Choices:     choices,
				Options:     nil,
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "colour",
					discordgo.EnglishUS: "color",
					discordgo.Bulgarian: "цвят",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "The name of your colour, run the /list-colour-roles command for more info",
					discordgo.EnglishUS: "The name of your color, run the /list-color-roles command for more info",
					discordgo.Bulgarian: "Името на цвета, пуснете /покажи-цветовите-роли за повече информация",
				},
			},
		},
	}
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
	if err != nil {
		fmt.Println(err)
	}
}

func createListColourRoles(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "list-colour-roles",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Gives info on the colour roles",
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "list-color-roles",
			discordgo.EnglishUS: "list-colour-roles",
			discordgo.Bulgarian: "покажи-цветовите-роли",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "Gives info on the colour roles",
			discordgo.EnglishUS: "Gives info on the colour roles",
			discordgo.Bulgarian: "Дава информация за цветовите роли",
		},
	}
	_, _ = s.ApplicationCommandCreate(s.State.User.ID, "", command)
}

func createMetaRole(s *discordgo.Session) {
	command := &discordgo.ApplicationCommand{
		Name:        "meta-role",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Manages meta roles",
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "meta-role",
			discordgo.EnglishUS: "meta-role",
			discordgo.Bulgarian: "мета-роля",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishGB: "Manages meta roles",
			discordgo.EnglishUS: "Manages meta roles",
			discordgo.Bulgarian: "Контролира мета ролите",
		},
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "give",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Creates and/or gives the user a certain meta role",
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "give",
					discordgo.EnglishUS: "give",
					discordgo.Bulgarian: "дай",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "Creates and/or gives the user a certain meta role",
					discordgo.EnglishUS: "Creates and/or gives the user a certain meta role",
					discordgo.Bulgarian: "Създава и/или дава на потребителя дадена мета роля",
				},
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "role-name",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The name of the role(case-sensitive)",
						Required:    true,
						NameLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "role-name",
							discordgo.EnglishUS: "role-name",
							discordgo.Bulgarian: "име-на-ролята",
						},
						DescriptionLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "The name of the role(case-sensitive)",
							discordgo.EnglishUS: "The name of the role(case-sensitive)",
							discordgo.Bulgarian: "Името на ролята(шрифта създава различни роли)",
						},
					},
				},
			},
			{
				Name:        "remove",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Removes a meta role",
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "remove",
					discordgo.EnglishUS: "remove",
					discordgo.Bulgarian: "премахни",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishGB: "Removes a meta role",
					discordgo.EnglishUS: "Removes a meta role",
					discordgo.Bulgarian: "Премахва мета роля",
				},
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "role-name",
						Type:        discordgo.ApplicationCommandOptionString,
						Description: "The name of the role(case-sensitive)",
						Required:    true,
						NameLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "role-name",
							discordgo.EnglishUS: "role-name",
							discordgo.Bulgarian: "име-на-ролята",
						},
						DescriptionLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishGB: "The name of the role(case-sensitive)",
							discordgo.EnglishUS: "The name of the role(case-sensitive)",
							discordgo.Bulgarian: "Името на ролята(шрифта премахва различни роли)",
						},
					},
				},
			},
		},
	}
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
	if err != nil {
		fmt.Println(err)
	}
}

func listColours(s *discordgo.Session, m *discordgo.InteractionCreate) {
	embedStrings := []map[discordgo.Locale]string{
		{
			discordgo.EnglishGB:    "Available colour names",
			discordgo.EnglishUS:    "Available color names",
			discordgo.Bulgarian:    "Възможни имена за цветове",
			discordgo.ChineseCN:    "Available colour names",
			discordgo.ChineseTW:    "Available colour names",
			discordgo.Croatian:     "Available colour names",
			discordgo.Czech:        "Available colour names",
			discordgo.Danish:       "Available colour names",
			discordgo.Dutch:        "Available colour names",
			discordgo.Finnish:      "Available colour names",
			discordgo.French:       "Available colour names",
			discordgo.German:       "Available colour names",
			discordgo.Greek:        "Available colour names",
			discordgo.Hungarian:    "Available colour names",
			discordgo.Italian:      "Available colour names",
			discordgo.Japanese:     "Available colour names",
			discordgo.Korean:       "Available colour names",
			discordgo.Lithuanian:   "Available colour names",
			discordgo.Norwegian:    "Available colour names",
			discordgo.Polish:       "Available colour names",
			discordgo.PortugueseBR: "Available colour names",
			discordgo.Romanian:     "Available colour names",
			discordgo.Russian:      "Available colour names",
			discordgo.SpanishES:    "Available colour names",
			discordgo.Swedish:      "Available colour names",
			discordgo.Ukrainian:    "Available colour names",
			discordgo.Hindi:        "Available colour names",
			discordgo.Thai:         "Available colour names",
			discordgo.Turkish:      "Available colour names",
			discordgo.Vietnamese:   "Available colour names",
			discordgo.Unknown:      "Available colour names",
		},
		{
			discordgo.EnglishGB:    "A color/name table can be found here",
			discordgo.EnglishUS:    "A colour/name table can be found here",
			discordgo.Bulgarian:    "Таблица с имена и цветове може да намерите тук",
			discordgo.ChineseCN:    "A colour/name table can be found here",
			discordgo.ChineseTW:    "A colour/name table can be found here",
			discordgo.Croatian:     "A colour/name table can be found here",
			discordgo.Czech:        "A colour/name table can be found here",
			discordgo.Danish:       "A colour/name table can be found here",
			discordgo.Dutch:        "A colour/name table can be found here",
			discordgo.Finnish:      "A colour/name table can be found here",
			discordgo.French:       "A colour/name table can be found here",
			discordgo.German:       "A colour/name table can be found here",
			discordgo.Greek:        "A colour/name table can be found here",
			discordgo.Hungarian:    "A colour/name table can be found here",
			discordgo.Italian:      "A colour/name table can be found here",
			discordgo.Japanese:     "A colour/name table can be found here",
			discordgo.Korean:       "A colour/name table can be found here",
			discordgo.Lithuanian:   "A colour/name table can be found here",
			discordgo.Norwegian:    "A colour/name table can be found here",
			discordgo.Polish:       "A colour/name table can be found here",
			discordgo.PortugueseBR: "A colour/name table can be found here",
			discordgo.Romanian:     "A colour/name table can be found here",
			discordgo.Russian:      "A colour/name table can be found here",
			discordgo.SpanishES:    "A colour/name table can be found here",
			discordgo.Swedish:      "A colour/name table can be found here",
			discordgo.Ukrainian:    "A colour/name table can be found here",
			discordgo.Hindi:        "A colour/name table can be found here",
			discordgo.Thai:         "A colour/name table can be found here",
			discordgo.Turkish:      "A colour/name table can be found here",
			discordgo.Vietnamese:   "A colour/name table can be found here",
			discordgo.Unknown:      "A colour/name table can be found here",
		},
		{
			discordgo.EnglishGB:    "Usage",
			discordgo.EnglishUS:    "Usage",
			discordgo.Bulgarian:    "Употреба",
			discordgo.ChineseCN:    "Usage",
			discordgo.ChineseTW:    "Usage",
			discordgo.Croatian:     "Usage",
			discordgo.Czech:        "Usage",
			discordgo.Danish:       "Usage",
			discordgo.Dutch:        "Usage",
			discordgo.Finnish:      "Usage",
			discordgo.French:       "Usage",
			discordgo.German:       "Usage",
			discordgo.Greek:        "Usage",
			discordgo.Hungarian:    "Usage",
			discordgo.Italian:      "Usage",
			discordgo.Japanese:     "Usage",
			discordgo.Korean:       "Usage",
			discordgo.Lithuanian:   "Usage",
			discordgo.Norwegian:    "Usage",
			discordgo.Polish:       "Usage",
			discordgo.PortugueseBR: "Usage",
			discordgo.Romanian:     "Usage",
			discordgo.Russian:      "Usage",
			discordgo.SpanishES:    "Usage",
			discordgo.Swedish:      "Usage",
			discordgo.Ukrainian:    "Usage",
			discordgo.Hindi:        "Usage",
			discordgo.Thai:         "Usage",
			discordgo.Turkish:      "Usage",
			discordgo.Vietnamese:   "Usage",
			discordgo.Unknown:      "Usage",
		},
		{
			discordgo.EnglishGB:    "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.EnglishUS:    "Find a color you like, get its name, fill any spaces with '-' and run the \"set-color-role\" command with it",
			discordgo.Bulgarian:    "Намерете цвят, който ви харесва, запишете му името, заместете всички пространства с '-' и пуснете командата \"сложи-роля-за-цвят\" command with it",
			discordgo.ChineseCN:    "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.ChineseTW:    "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Croatian:     "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Czech:        "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Danish:       "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Dutch:        "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Finnish:      "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.French:       "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.German:       "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Greek:        "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Hungarian:    "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Italian:      "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Japanese:     "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Korean:       "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Lithuanian:   "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Norwegian:    "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Polish:       "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.PortugueseBR: "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Romanian:     "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Russian:      "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.SpanishES:    "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Swedish:      "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Ukrainian:    "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Hindi:        "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Thai:         "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Turkish:      "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Vietnamese:   "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
			discordgo.Unknown:      "Find a colour you like, get its name, fill any spaces with '-' and run the \"set-colour-role\" command with it",
		},
		{
			discordgo.EnglishGB:    "The following colours don't have a name so we made up a name for them",
			discordgo.EnglishUS:    "The following colors don't have a name so we made up a name for them",
			discordgo.Bulgarian:    "Следните цветове нямат име, за това измислихме следните имена за тях",
			discordgo.ChineseCN:    "The following colours don't have a name so we made up a name for them",
			discordgo.ChineseTW:    "The following colours don't have a name so we made up a name for them",
			discordgo.Croatian:     "The following colours don't have a name so we made up a name for them",
			discordgo.Czech:        "The following colours don't have a name so we made up a name for them",
			discordgo.Danish:       "The following colours don't have a name so we made up a name for them",
			discordgo.Dutch:        "The following colours don't have a name so we made up a name for them",
			discordgo.Finnish:      "The following colours don't have a name so we made up a name for them",
			discordgo.French:       "The following colours don't have a name so we made up a name for them",
			discordgo.German:       "The following colours don't have a name so we made up a name for them",
			discordgo.Greek:        "The following colours don't have a name so we made up a name for them",
			discordgo.Hungarian:    "The following colours don't have a name so we made up a name for them",
			discordgo.Italian:      "The following colours don't have a name so we made up a name for them",
			discordgo.Japanese:     "The following colours don't have a name so we made up a name for them",
			discordgo.Korean:       "The following colours don't have a name so we made up a name for them",
			discordgo.Lithuanian:   "The following colours don't have a name so we made up a name for them",
			discordgo.Norwegian:    "The following colours don't have a name so we made up a name for them",
			discordgo.Polish:       "The following colours don't have a name so we made up a name for them",
			discordgo.PortugueseBR: "The following colours don't have a name so we made up a name for them",
			discordgo.Romanian:     "The following colours don't have a name so we made up a name for them",
			discordgo.Russian:      "The following colours don't have a name so we made up a name for them",
			discordgo.SpanishES:    "The following colours don't have a name so we made up a name for them",
			discordgo.Swedish:      "The following colours don't have a name so we made up a name for them",
			discordgo.Ukrainian:    "The following colours don't have a name so we made up a name for them",
			discordgo.Hindi:        "The following colours don't have a name so we made up a name for them",
			discordgo.Thai:         "The following colours don't have a name so we made up a name for them",
			discordgo.Turkish:      "The following colours don't have a name so we made up a name for them",
			discordgo.Vietnamese:   "The following colours don't have a name so we made up a name for them",
			discordgo.Unknown:      "The following colours don't have a name so we made up a name for them",
		},
	}

	embed := NewEmbed().
		SetTitle(embedStrings[0][m.Locale]).
		AddField(embedStrings[1][m.Locale], "https://www.spycolor.com/w3c-colors").
		AddField(embedStrings[2][m.Locale], embedStrings[3][m.Locale]).
		InlineAllFields().
		AddField(embedStrings[4][m.Locale], "#40e0d0, #a0522d, #fa8072, #7f007f, #ffe4b5, #66cdaa").
		AddField("#40e0d0", "Light-Cyan-2").
		AddField("#a0522d", "Dark-Sand").
		AddField("#fa8072", "Salmon").
		AddField("#7f007f", "Sienna").
		AddField("#ffe4b5", "SandySahara").
		AddField("#66cdaa", "Tropical-Green").
		InlineAllFields().
		SetFooter(footerTranslations[m.Locale], "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").SetColor(0xf1c40f).MessageEmbed

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

func giveColour(arg string, s *discordgo.Session, m *discordgo.InteractionCreate) {
	embedStrings := []map[discordgo.Locale]string{
		{},
		{
			discordgo.EnglishGB:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.EnglishUS:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Bulgarian:    "Не можахме да добавим мета роля, понеже канала в който сте, не е маркиран с тип за поставяне на мета роли, моля намерете канал с правилните права",
			discordgo.ChineseCN:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.ChineseTW:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Croatian:     "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Czech:        "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Danish:       "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Dutch:        "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Finnish:      "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.French:       "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.German:       "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Greek:        "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Hungarian:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Italian:      "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Japanese:     "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Korean:       "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Lithuanian:   "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Norwegian:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Polish:       "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.PortugueseBR: "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Romanian:     "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Russian:      "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.SpanishES:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Swedish:      "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Ukrainian:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Hindi:        "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Thai:         "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Turkish:      "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Vietnamese:   "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Unknown:      "Couldn't add meta role, as the current channel is not marked for meta roles",
		},
		{
			discordgo.EnglishGB:    "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.EnglishUS:    "Invalid colour! Run \"list-color-roles\" to get a list of all available colors!",
			discordgo.Bulgarian:    "Грешен цвят! Пуснете командата \"покажи-цветовите-роли\" за да получите лист от всички възможни цветоте!",
			discordgo.ChineseCN:    "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.ChineseTW:    "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Croatian:     "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Czech:        "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Danish:       "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Dutch:        "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Finnish:      "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.French:       "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.German:       "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Greek:        "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Hungarian:    "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Italian:      "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Japanese:     "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Korean:       "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Lithuanian:   "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Norwegian:    "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Polish:       "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.PortugueseBR: "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Romanian:     "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Russian:      "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.SpanishES:    "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Swedish:      "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Ukrainian:    "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Hindi:        "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Thai:         "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Turkish:      "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Vietnamese:   "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
			discordgo.Unknown:      "Invalid colour! Run \"list-colour-roles\" to get a list of all available colours!",
		},
		{
			discordgo.EnglishGB:    "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.EnglishUS:    "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel color-role\" command in order to set up colour roles!",
			discordgo.Bulgarian:    "Каналът, в който се намирате не е маркиран като \"ubot-colour-pick\", намерете друг канал или се свържете с модератора на сървъра и го накарайте да пусне командата \"променяне-на-канал цветови-роли\"!",
			discordgo.ChineseCN:    "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.ChineseTW:    "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Croatian:     "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Czech:        "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Danish:       "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Dutch:        "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Finnish:      "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.French:       "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.German:       "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Greek:        "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Hungarian:    "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Italian:      "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Japanese:     "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Korean:       "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Lithuanian:   "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Norwegian:    "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Polish:       "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.PortugueseBR: "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Romanian:     "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Russian:      "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.SpanishES:    "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Swedish:      "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Ukrainian:    "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Hindi:        "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Thai:         "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Turkish:      "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Vietnamese:   "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
			discordgo.Unknown:      "Channel not marked as \"ubot-colour-pick\", contact your server's moderator to run the \"set-channel colour-role\" command in order to set up colour roles!",
		},
	}
	channel, _ := s.Channel(m.ChannelID)
	if strings.Contains(strings.ToLower(channel.Topic), "ubot-colour-pick") {
		bFound := false
		colour := 0
		var name = ""

		for i := 0; i < len(colours); i++ {
			if strings.ToLower(colours[i]) == strings.ToLower(arg) {
				bFound = true
				colour = int(coloursHex[i])
				name = colours[i]
				break
			}
		}

		if bFound {
			roles, _ := s.GuildRoles(m.GuildID)
			for i := 0; i < len(roles); i++ {
				if name == roles[i].Name {
					_ = s.GuildMemberRoleAdd(m.GuildID, m.Member.User.ID, roles[i].ID)
					embed := NewEmbed().
						SetTitle("Added you to the "+name+" role!").
						AddField(roleName[m.Locale], roles[i].Mention()).
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
					return
				}
			}

			perms := int64(0)

			roledata := discordgo.RoleParams{}
			roledata.Permissions = &perms
			roledata.Color = &colour
			roledata.Name = name

			role, _ := s.GuildRoleCreate(m.GuildID, &roledata)
			_ = s.GuildMemberRoleAdd(m.GuildID, m.Member.User.ID, role.ID)
			embed := NewEmbed().
				SetTitle("Added you to the "+name+" role!").
				AddField(roleName[m.Locale], role.Mention()).
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
			_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					TTS:     false,
					Content: embedStrings[2][m.Locale],
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
		}
	} else {
		_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				TTS:     false,
				Content: embedStrings[3][m.Locale],
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
}

func giveMetarole(arg string, s *discordgo.Session, m *discordgo.InteractionCreate) {
	// Add translations for the header embed text
	embedStrings := []map[discordgo.Locale]string{
		{},
		{
			discordgo.EnglishGB:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.EnglishUS:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Bulgarian:    "Не можахме да добавим мета роля, понеже канала в който сте, не е маркиран с тип за поставяне на мета роли, моля намерете канал с правилните права",
			discordgo.ChineseCN:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.ChineseTW:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Croatian:     "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Czech:        "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Danish:       "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Dutch:        "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Finnish:      "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.French:       "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.German:       "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Greek:        "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Hungarian:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Italian:      "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Japanese:     "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Korean:       "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Lithuanian:   "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Norwegian:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Polish:       "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.PortugueseBR: "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Romanian:     "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Russian:      "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.SpanishES:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Swedish:      "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Ukrainian:    "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Hindi:        "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Thai:         "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Turkish:      "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Vietnamese:   "Couldn't add meta role, as the current channel is not marked for meta roles",
			discordgo.Unknown:      "Couldn't add meta role, as the current channel is not marked for meta roles",
		},
	}
	channel, _ := s.Channel(m.ChannelID)

	if strings.Contains(strings.ToLower(channel.Topic), "ubot-meta-role-pick") {
		roles, _ := s.GuildRoles(m.GuildID)
		for i := 0; i < len(roles); i++ {
			// That magic number that you see here is the default permissions integer when a new role is created
			if strings.ToLower(roles[i].Name) == strings.ToLower(arg) && roles[i].Permissions == 0 {
				_ = s.GuildMemberRoleAdd(m.GuildID, m.Member.User.ID, roles[i].ID)
				embed := NewEmbed().
					SetTitle("Added you to the "+roles[i].Name+" role!").
					AddField(roleName[m.Locale], roles[i].Mention()).
					SetFooter(footerTranslations[m.Locale], "https://avatars.githubusercontent.com/u/66491677?s=400&u=07d8dd94266f97e22ee5bd96aebb6a5f9190b4ec&v=4").
					SetColor(0xf1c40f).MessageEmbed
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
				return
			}
		}
		perms := int64(0)

		roleinfo := discordgo.RoleParams{}
		roleinfo.Permissions = &perms
		roleinfo.Name = arg

		role, _ := s.GuildRoleCreate(m.GuildID, &roleinfo)
		_ = s.GuildMemberRoleAdd(m.GuildID, m.Member.User.ID, role.ID)
		embed := NewEmbed().
			SetTitle("Added you to the "+role.Name+" role!").
			AddField(roleName[m.Locale], role.Mention()).
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
		_ = s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				TTS:     false,
				Content: embedStrings[1][m.Locale],
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
}

func removeMetarole(arg string, s *discordgo.Session, m *discordgo.InteractionCreate) {
	embedStrings := []map[discordgo.Locale]string{
		{
			discordgo.EnglishGB:    "Removed metarole",
			discordgo.EnglishUS:    "Removed metarole",
			discordgo.Bulgarian:    "Изтрихме мета роля",
			discordgo.ChineseCN:    "Removed metarole",
			discordgo.ChineseTW:    "Removed metarole",
			discordgo.Croatian:     "Removed metarole",
			discordgo.Czech:        "Removed metarole",
			discordgo.Danish:       "Removed metarole",
			discordgo.Dutch:        "Removed metarole",
			discordgo.Finnish:      "Removed metarole",
			discordgo.French:       "Removed metarole",
			discordgo.German:       "Removed metarole",
			discordgo.Greek:        "Removed metarole",
			discordgo.Hungarian:    "Removed metarole",
			discordgo.Italian:      "Removed metarole",
			discordgo.Japanese:     "Removed metarole",
			discordgo.Korean:       "Removed metarole",
			discordgo.Lithuanian:   "Removed metarole",
			discordgo.Norwegian:    "Removed metarole",
			discordgo.Polish:       "Removed metarole",
			discordgo.PortugueseBR: "Removed metarole",
			discordgo.Romanian:     "Removed metarole",
			discordgo.Russian:      "Removed metarole",
			discordgo.SpanishES:    "Removed metarole",
			discordgo.Swedish:      "Removed metarole",
			discordgo.Ukrainian:    "Removed metarole",
			discordgo.Hindi:        "Removed metarole",
			discordgo.Thai:         "Removed metarole",
			discordgo.Turkish:      "Removed metarole",
			discordgo.Vietnamese:   "Removed metarole",
			discordgo.Unknown:      "Removed metarole",
		},
	}
	g, _ := s.Guild(m.GuildID)
	roles := g.Roles

	for i := 0; i < len(roles); i++ {
		if strings.ToLower(arg) == strings.ToLower(roles[i].Name) && roles[i].Permissions == 0 {
			_ = s.GuildMemberRoleRemove(m.GuildID, m.Member.User.ID, roles[i].ID)
			embed := NewEmbed().
				SetTitle(embedStrings[0][m.Locale]).
				AddField(roleName[m.Locale], roles[i].Mention()).
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
			return
		}
	}
}
