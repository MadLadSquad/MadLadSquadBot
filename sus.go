package main

import (
	"github.com/MadLadSquad/discordgo"
	"math/rand"
	"time"
)

func sus(s *discordgo.Session, m *discordgo.MessageCreate) {
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
		"https://cdn.discordapp.com/attachments/785618029647364116/882583230010757160/8e46383b9ff1df0fbf3a789c9be4fea7.mp4", "https://youtube.com/shorts/zs8cC6ujbJc", "https://youtube.com/shorts/ZDZPr9_WM8Q", "https://youtube.com/shorts/q0YZS3Y182g", "https://youtube.com/shorts/qZC4mwq1MAE", "https://youtube.com/shorts/qZC4mwq1MAE", "https://cdn.discordapp.com/attachments/887988701907533856/892098783133962301/38f1aa17bb1b04b99ec2e112455d5930.mp4",
		"https://www.youtube.com/watch?v=IbwxM_vRDIg", "https://cdn.discordapp.com/attachments/887988701907533856/892098783981207582/bb44cd5e518d0959fadc1bc50fe85a50.mp4", "https://cdn.discordapp.com/attachments/887988701907533856/892098784350334986/4f8c6742ce759f719af0256a88c6d5d9.mp4", "https://cdn.discordapp.com/attachments/887988701907533856/892098784350334986/4f8c6742ce759f719af0256a88c6d5d9.mp4",
		"https://cdn.discordapp.com/attachments/887988701907533856/891947435688087582/d-1.mp4", "https://cdn.discordapp.com/attachments/813982446869151744/891923535721807882/Impostor.mp4", "https://youtu.be/5Z3Zo1nY6W0", "https://cdn.discordapp.com/attachments/887988701907533856/892099152006242374/4b64fd621bce64af243205f51d2e6e51.mp4", "https://cdn.discordapp.com/attachments/887988701907533856/892099286492401724/67061f471a02c43128f357a20a8336d2.mp4",
		"https://cdn.discordapp.com/attachments/887988701907533856/892100345700966400/eba36cd940792a84ff43018a5d71711c.mp4", "https://cdn.discordapp.com/attachments/887988701907533856/892100346095210516/526e176017b85c6d0e38c438e27c8d1b.mp4", "https://cdn.discordapp.com/attachments/887988701907533856/892101096598798346/unknown.png",
	}
	rand.Seed(time.Now().UnixNano())
	_, _ = s.ChannelMessageSend(m.ChannelID, sussyMessages[rand.Intn(len(sussyMessages))])
}

func pernik(s *discordgo.Session, m *discordgo.MessageCreate) {
	sussyMessages := []string{
		"https://cdn.discordapp.com/attachments/827543014959480842/917454171001737236/9k.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917454638276546570/himichesko.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917454740118466620/media.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917455069694263336/1b0dbcc91b13c8142d4d57ac9f7bcaa4.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917455933079166976/IMG_4247.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917456284255649832/PERNIK12.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917456785244315678/IMG_4248.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917456925019500604/PERNIK14.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917457674285764608/IMG_4249.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917457995468767252/PERNIK6.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917458754465845338/IMG_4252.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917459589916655647/PERNIK10.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917460380576542730/IMG_4253.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917460456636035092/PERNIK7.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917460779811360798/IMG_4251.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917460884354375740/PERNIK8.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917461192023367750/IMG_4255.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917461247526596608/PERNIK9.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917461451768205342/IMG_4250.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917461538061840384/PERNIK1.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917462208819130448/IMG_4256.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917462315358617610/PERNIK3.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917462580392493076/IMG_4257.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917462638932422706/PERNIK4.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917462842905612339/IMG_4262.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917463078835212298/346_img_gallery.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917463773416144936/IMG_4264.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917463810577690665/pernishko-asfaltirane.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917463905318625340/IMG_4268.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917464031592329256/pernishko-taxi.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917464139306262548/IMG_4266.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917464182918615040/pernishko-rodeo.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917464281052766218/43D39E0D-5E2D-4570-B16E-0263F49D3685.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917464375835631686/tp9Kp6_UxL6fhmhtAjvhWf-iV6VVqpWBytjrI7_zcg-Vdi9O5R7cVwQp7xxp3gb9ZnYKMTsg05YiE-B6-4lfninFZUrElSsz4l9NM-tBjw.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917464731210625074/IMG_4269.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917464894431969290/pernishka-baba.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917465266210869298/IMG_4270.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917465403083599873/74-09-01-2.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917465571635904532/IMG_4261.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917465641882091581/D09FD0B5D180D0BDD0B8D0BA-D181D0BCD18FD185.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917466066857365554/IMG_4274.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917466236630204466/D0BFD0B5D180D0BDD0B8D188D0BAD0B0-D0BAD0BED181D182D0B5D0BDD183D180D0BAD0B0.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917466568802320484/pernishka-torta.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917466779507388456/IMG_4272.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917467611695050822/IMG_4278.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917468013123481620/attachment.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917468452380364831/IMG_4279.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917469365291614218/1689188_585716838177617_241186335_n.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917469740857950238/IMG_4276.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917469855786090526/48390786_10151611401024945_2044066530638954496_n.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917470217750343690/IMG_4273.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917470280824258622/pernik2B15825950_347802638935604_874724176177034871_n.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917470924557664286/IMG_4277.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917471154762051654/409549_280439372010517_1466644182_n.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917479186002772039/D95CC441-EDBB-4F22-9791-006A69FBA75B.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917479447278522479/d0b2d0b5d0bbd0bed0b5d180d0b3d0bed0bcd0b5d182d18ad180.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917491257780101160/258_1331560155.png",
		"https://tenor.com/view/%D0%BF%D0%B5%D1%80%D0%BD%D0%B8%D0%BA-pernik-bulgaria-%D0%B1%D1%8A%D0%BB%D0%B3%D0%B0%D1%80%D0%B8%D1%8F-gif-15400647",
		"https://cdn.discordapp.com/attachments/827543014959480842/917496558633746464/pernishki-pozhelaniya.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917505498373296168/IMG_4260.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917506120665423882/78772639_437321033611638_3805015684778295296_n.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917506610522366063/IMG_4254.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917507454760267847/242169522_3899293473510701_2010850538164865819_n.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917508186209128448/IMG_4267.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917508287489011763/207408429_3665040243602693_863138013704856222_n.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917512277853356062/7cfb16b3b6083c7a.jpeg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917521202870878218/IMG_4285.jpg",
		"https://media.discordapp.net/attachments/830385500887711786/917522698102534154/motivate.png?width=533&height=676",
		"https://cdn.discordapp.com/attachments/827543014959480842/917523482168938526/IMG_4283.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917525377948201000/197162266_3617596768347041_4392908811757837678_n.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917525791376572426/IMG_4284.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917525960763514940/53500045_1807302519376484_3113171958917758976_n.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917526034134499379/IMG_4285.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917526330319446076/53229618_1805754319531304_7231227687979712512_n.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917527022132158514/IMG_4281.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917527061835444244/53602129_1797586717014731_2244352366163787776_n.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917527179502448700/465D32B1-C16F-4C46-B33E-E8F5068D10D3.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917527833545412648/48374423_1691750167598387_5305639878900318208_n.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917529395848151080/IMG_4287.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917529460759220285/48395743_1689154764524594_5979990013985685504_n.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917530184373108766/39042060_1532842003489205_6760760694112518144_n.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917530547570507777/IMG_4289.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917530686188027964/37147124.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917531268386799646/IMG_4295.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917531485370716230/oCDsStC9diS4WcqSz6E9szIFUEIKo6xChTDmouNrND5Jm4RKH7pmeLw9HsqPJCduyyAErUP0jKQOkMaRX28QF4GGdKNWUEsuetMhScclP1sVyRYthitHr-j2qTlVSpby5NXcWJ2PvwazdXlgbGTWZkAB_KNdyLic968.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917535092581212190/IMG_4299.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917535222554312704/pernishka-yahta.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917539104604966922/IMG_4296.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917540271338709002/D0BFD0B5D180D0BDD0B8D0BA-D188D0B5D0B3D0B8.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917541230433435678/IMG_4294.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917541305842798612/72830829.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917541518917664768/IMG_4293.jpg",
		"https://www.flagman.bg/files/images/4%281169%29.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917568732392480798/istinski_pernichanin.png",
		"https://cdn.discordapp.com/attachments/827543014959480842/917571247842426930/IMG_4291.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917738838531649556/Screenshot_20211207-132554_Google.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917739558525878282/IMG_4292.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917739910218256384/Screenshot_20211207-133035_Google.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917742572036821012/IMG_4303.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917751142409207838/Screenshot_20211207-141507_Google.jpg",
		"https://cdn.discordapp.com/attachments/827543014959480842/917769793745076254/IMG_4306.jpg",
	}
	rand.Seed(time.Now().UnixNano())
	_, _ = s.ChannelMessageSend(m.ChannelID, sussyMessages[rand.Intn(len(sussyMessages))])
}
