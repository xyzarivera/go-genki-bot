# 🌟  Genki Bot

An experimental Discord bot created while learning Golang as a language and deploying it as a project.

This was created during my time as a student at [Code Chrysalis](https://github.com/codechrysalis)

## 🌙 Goals

### Personal

- Learn how to create a Discord bot
- Connect this Discord bot to my live [cc-solo-mvp](https://github.com/xyzarivera/cc-solo-mvp) project
- more of my learning notes [here](notes.md)

### Basic

- Explore a completely new programming language
- Build an application
- Deploy to the cloud!

### Advanced

- Create tests

## ☀️ Features

- Go as language
- [DiscordGo](https://github.com/bwmarrin/discordgo) as main component to create the Discord bot
- Heroku as cloud platform

## ⚙️ Setup

### Pre-requisites

Install Go version 1.17

### Clone project

```
git clone https://github.com/xyzarivera/go-genki-bot.git

```

### Testing the app without building

```
export BOT_TOKEN=<your bot token>
go run main.go -t $BOT_TOKEN
```

## 📖 Discord Bot Commands

### !helloWorld

- test connection command
- expected answer
  ```
  "Hi " + username + " ! :smiley:"
  ```

### !standup

- standup session start command
- expected answer
  ```
  Start of standup 
        What did you do?
        What will you do?
        Do you have any blockers?
  ```

### Replying to the bot

- test response when a user replies to **ANY** of the Bot's message
- expected answer 
  ```
  "Got this! See you later " + username + " :grin:"
  ```