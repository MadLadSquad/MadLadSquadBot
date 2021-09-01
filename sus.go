package main

import (
	"github.com/MadLadSquad/discordgo"
	"math/rand"
	"time"
)

func sus(s* discordgo.Session, m* discordgo.MessageCreate) {
	sussyMessages := []string{"Say \"sus\" but backwards", "Say \"bus\" but replace the B with S",
		"Say \"pegasus\" but replace the 'p' with 'm'", "Say \"sus\" but replace \"sus\" with \"sus\"",
		"Say \"Amongus\" but replace \"Among\" with s", "Say \"I invented\" but remove the \"in\"",
		"Are you sus? A: Yes; B: A", "Say \"I'm poster\" without the space", "Say \"sups\" without your lips touching",
		"Say \"sus\" without touching your lips", "Say \"I'm gay\" but replace \"I'm gay\" with \"sus\" <@676763652367450112>",
		"Say \"sussy\" but remove the 'y'", "https://www.youtube.com/watch?v=B-5nXUyD6l0",
		"https://www.youtube.com/watch?v=gBNJxlhanoI", "https://cdn.discordapp.com/attachments/785618029647364116/882576989314228244/unknown.png",
		"https://www.youtube.com/watch?v=LTmpLr0BPdM", "https://www.youtube.com/watch?v=riP4xifJq4w",
		"Say \"sushi\" without the \"hi\"", "https://www.youtube.com/watch?v=OX0ceSQqED4", "https://cdn.discordapp.com/attachments/785618029647364116/882581905894178856/44769205c8fbd59cecea48f545b336d9.mp4",
		"https://cdn.discordapp.com/attachments/785618029647364116/882579678412566558/VID_103171106_022930_740.mp4", "https://cdn.discordapp.com/attachments/785618029647364116/882579820364566528/b40ae52fc0707052be45af94367a9f92.mp4", "https://cdn.discordapp.com/attachments/785618029647364116/882580411883094016/d45249f55518039e4f2a6ccf1a1d2026.mp4", "https://cdn.discordapp.com/attachments/785618029647364116/882580496884826132/VID_273871207_103024_481.mp4",
		"https://cdn.discordapp.com/attachments/785618029647364116/882580647464542258/aa69e48e55deda875f921fdb36af4305.mp4", "https://cdn.discordapp.com/attachments/785618029647364116/882581234163806219/e0152898d75fcc8f7704a19ec277da35.mp4", "https://cdn.discordapp.com/attachments/785618029647364116/882581379920056370/a39a1a617c63f04ed4cdd54338902159.mp4", "https://cdn.discordapp.com/attachments/785618029647364116/882581591040352256/02ee9b5b06a45fba5f6913a290bcaffc.mp4",
		"https://cdn.discordapp.com/attachments/785618029647364116/882581784351637514/fadeeba67fc873cbc6de697e4e846645.mp4", "https://cdn.discordapp.com/attachments/785618029647364116/882582117861699644/543d1fca063a106431e845549c0b4a3e.mp4", "https://cdn.discordapp.com/attachments/785618029647364116/882582755576279100/375fdf4b26ae608941cc17de6ed281ee.mp4", "https://cdn.discordapp.com/attachments/785618029647364116/882583053602537492/fd08a12fa0a3cdc3cb7e1befeadfeab1.mp4",
		"https://cdn.discordapp.com/attachments/785618029647364116/882583230010757160/8e46383b9ff1df0fbf3a789c9be4fea7.mp4",
	}
	rand.Seed(time.Now().UnixNano())
	s.ChannelMessageSend(m.ChannelID, sussyMessages[rand.Intn(len(sussyMessages))])
}