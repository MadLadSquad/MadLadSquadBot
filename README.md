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
- Event logs ✅
- Welcome message ✅
- Easy set commands ✅
- Welcome messages with custom images
- Automatic moderation
- Verification ✅
- Verification with CAPTCHA
- Restricted channels ✅
- Topic locked channels ✅
- Twitch, Twitter and Youtube updates
## How to use
To use the bot you can invite the UntitledDiscordBot using this link or you can build it yourself. By default the bot uses `&` as it's prefix. 
### Verification
Create a base role called `Member`. Next create a channel called `verify` for anyone except people with roles. Make them type `&verify` to enter the server. This is a base level verify command so it is better to use `Verification with CAPTCHA` when it comes out

You can also use the following set command to make a member role
```
&generate-member-role
```
### Welcome/Goodbye messages
To use the welcome and goodbye messages add `ubot-welcome` to the channel's topic

You can also use the following set function
```
&set-welcome <ping to channel>
```
### Ban
To ban a user type in `&ban <a ping to them> <reason up to 100 words>`. All messages in the past 30 days will be deleted from the server
### Event tracking
To track events make a channel that is available to the bot and add to the topic the following word `ubot-event-log`. Now every time something happens in the server it will be logged!

You can also use the following set command to make this channel
```
&set-event-tracking <ping to channel>
```
### Restricted/Topic locked channels
Some channels can be restricted by adding one of the following keywords in the topic of a channel
1. `ubot-restrict-text-only` text messages only
2. `ubot-restrict-attachments-only` attachments and links only
3. `ubot-restrict-links-only` links only

Instead of doing this manually you can also use the set commands
1. `&set-text-only <ping to channnel>`
2. `&set-attachments-only <ping to channnel>`
3. `&set-links-only <ping to channnel>`
## How to build
1. Install golang version atleast 1.16
2. Clone this repo with `git clone https://github.com/MadLadSquad/MadLadSquadBot.git`
3. Enter the folder where the source is stored
4. Run `go build . -o Bot` to create the executable
5. Run `Bot -t <your bot's token here>`
6. The bot should start to print some information to the screen
7. You are ready to use your own instance of the bot  
