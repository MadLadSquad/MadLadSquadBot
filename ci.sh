#!/bin/bash
vertmp=$(grep "github.com/MadLadSquad/discordgo" go.mod)
ver=${vertmp: -8}

go get github.com/MadLadSquad/discordgo@"${ver}"
go build .
