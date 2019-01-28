package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type bot struct {
	token  string `json:"token"`
	botid  string
	prefix string
	owner  string
}

var (
	commandPrefix string
	botID         string
)

func main() {
	discord, err := discordgo.New("Bot TOKEN")
	errCheck("error creating discord session", err)
	user, err := discord.User("@me")
	errCheck("error retrieving account", err)

	botID = user.ID
	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err = discord.UpdateStatus(0, "Yeah")
		if err != nil {
			fmt.Println("Error attempting to set my status")
		}
		servers := discord.State.Guilds
		fmt.Printf("OhDaddy has started on %d servers", len(servers))
	})

	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	commandPrefix = "!"

	<-make(chan struct{})

}

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		//Do nothing because the bot is talking
		return
	}
	if message.Content == "!ping" {
		discord.ChannelMessageSend(message.ChannelID, "Pong")
		return
	}
	if message.Content == "!github" {
		discord.ChannelMessageSend(message.ChannelID, "https://github.com/Bmbus/Oh_Daddy")
		return
	}
	//content := message.Content

	// fmt.Printf("Message: %+v || From: %s\n", message.Message, message.Author)
}

// func github(discord *discordgo.Session, message discordgo.Message) {
// 	if message.Content == "!github" {
// 		discord.ChannelMessageSend(message.ChannelID, "https://github.com/Bmbus/Oh_Daddy")
// 		return
// 	}
// }
