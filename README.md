# MadLadSquadBot
[![CI](https://github.com/MadLadSquad/MadLadSquadBot/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/MadLadSquad/MadLadSquadBot/actions/workflows/ci.yml)

The official discord bot for the MadLadSquad discord server
## Features
- Serverinfo ✅
- Userinfo ✅
- Help ✅
- Playing music in vc
- Ban ✅
- Kick ✅
- Event logs
- Welcome message
- Automatic moderation
- Verification
- Verification with CAPTCHA
- Restricted channels
- Topic locked channels
- Twitch, Twitter and Youtube updates
## How to use
To use the bot you can invite the UntitledDiscordBot using this link or you can build it yourself. By default the bot uses `&` as it's prefix. 
### Verification
Create a base role called `Member`. Next create a channel called `verify` for anyone except people with roles. Make them type `&verify` to enter the server. This is a base level verify command so it is better to use `Verification with CAPTCHA` when it comes out
### Ban
To ban a user type in `&ban <a ping to them> <reason up to 100 words>`. All messages in the past 30 days will be deleted from the server
## How to build
1. Install golang version atleast 1.16
2. Clone this repo with `git clone https://github.com/MadLadSquad/MadLadSquadBot.git`
3. Enter the folder where the source is stored
4. Run `go build . -o Bot` to create the executable
5. Run `Bot -t <your bot's token here>`
6. The bot should start to print some information to the screen
7. You are ready to use your own instance of the bot  
