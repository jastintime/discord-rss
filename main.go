package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
	"database/sql"
	_ "modernc.org/sqlite"
)
type Guild struct {
	GuildID  string `json:"guild_id"`
	Name string 
}

type Post struct {
	Id string  `json:"id"`
	GuildID string `json:"guild_id"` 
	Author User `json:"author"`
}
type User struct {
	Username string `json:"username"`
}



var db *sql.DB

func main() {
	db, err := sql.Open("sqlite", "discord.db");
	if err != nil {
		fmt.Println("failed to open discord.db");
		return 
	}
	defer db.Close()
	token := os.Getenv("DISCORD_TOKEN")
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("failed to start discord session, ", err)
		return
	}
	session.AddHandler(messageCreate)
	session.Identify.Intents = discordgo.IntentsGuildMessages
	err = session.Open()
	if err != nil {
		fmt.Println("error opening a discord connection,", err)
		return
	}
	fmt.Println("Bot running, Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	session.Close()
}



func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == session.State.User.ID {
		return
	}
	//TODO: add a mechanicism to allow people to say which channel is 
	// the announcement.
	if message.ChannelID != "1360458157939494953" {
		return 
	}
	if message.Content == "Pong" {
		fmt.Println("received");
	}



}
