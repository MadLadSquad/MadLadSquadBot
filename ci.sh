#!/bin/bash
vertmp=$(grep "github.com/MadLadSquad/discordgo" go.mod)
ver=${vertmp: -8}

go mod download github.com/MadLadSquad/discordgo
go get github.com/MadLadSquad/discordgo@"${ver}"
go build .
