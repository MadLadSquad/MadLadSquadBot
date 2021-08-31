package main

import "github.com/MadLadSquad/discordgo"

func onReady(s *discordgo.Session, m *discordgo.Ready) {
	err := s.UpdateStreamingStatus(1, "&help | Click watch for music!", "https://www.youtube.com/watch?v=hvfe_JPovW8&list=PL8yFU3veFghtteYYFdSnc-vaBZ3hHh4Wc&index=1")
	if err != nil {
		return
	}
}
