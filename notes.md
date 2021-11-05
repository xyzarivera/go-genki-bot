# Personal Notes

## How I started my project

- created a github repository, cloned to local
- created gitignore file
- start go project
  - `go mod init https://github.com/xyzarivera/go-genki-bot`
- install discordgo package
  - `go get github.com/bwmarrin/discordgo`
- create `main.go` file

### Testing app locally

- `go run main.go -t $BOT_TOKEN`

## Deploying to Heroku

- prefer to use the console to setup CI/CD

- create Procfile
  - use `worker` process type to accommodate websockets
  - `worker: ./bin/go-genki-bot -t $TOKEN`
  - set `TOKEN` in Heroku app's config vars 

- add `go.mod` file to include this line above go version declaration
  ``` 
  // +heroku goVersion go1.17
  go 1.17
  ```

## Reference

- [Learning Go by examples tutorial by Aurélie Vache](https://dev.to/aurelievache/learning-go-by-examples-part-4-create-a-bot-for-discord-in-go-43cf)
- [Getting Started on Heroku with Go](https://devcenter.heroku.com/articles/getting-started-with-go)