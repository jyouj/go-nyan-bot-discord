package main

import (
  "fmt"
  "os"

  "github.com/bwmarrin/discordgo"
)

func main() {
  d, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
  if err != nil {
    fmt.Println("failed to create discord session", err)
  }

  // bot, err := d.User("@" + os.Getenv("DISCORD_CLIENT_ID"))
  if err != nil {
    fmt.Println("failed to access account", err)
  }

  d.AddHandler(handleCmd)
  err = d.Open()
  if err != nil {
    fmt.Println("unable to establish connection", err)
  }

  defer d.Close()

  <-make(chan struct{})
}

func handleCmd(s *discordgo.Session, msg *discordgo.MessageCreate) {
  user := msg.Author
  if user.ID == s.State.User.ID || user.Bot {
    return
  }

  content := msg.Content

  if (content == "hello") {
    s.ChannelMessageSend(msg.ChannelID, content + ", World!")
  } else {
    s.ChannelMessageSend(msg.ChannelID, content + "にゃーん")
  }

  fmt.Printf("Message: %+v\n", msg.Message)
}
