package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/bwmarrin/discordgo"
	_ "modernc.org/sqlite"
	"os"
	"os/signal"
	"syscall"
)

//go:embed create_table.sql
var createTable string

type Bot struct {
	db *sql.DB
}

func main() {
	db, err := sql.Open("sqlite", "file:discord.sqlite3?_pragma=foreign_keys(1)")
	if err != nil {
		fmt.Println("Error creating DB Object:", err)
		return
	}
	defer db.Close()
	bot := Bot{db: db}

	token, ok := os.LookupEnv("DISCORD_TOKEN")
	if !ok {
		fmt.Println("DISCORD_TOKEN not set, this should be set as your discord bot API key")
		return
	}
	_, err = db.Exec(createTable)
	if err != nil {
		fmt.Println("Failed to create database:", err)
		return
	}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("failed to start discord session, ", err)
		return
	}
	//	session.AddHandler(eventPrinter)
	session.AddHandler(bot.messageCreate)
	session.AddHandler(bot.joinServer)
	session.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuilds

	err = session.Open()
	if err != nil {
		fmt.Println("error opening a discord connection,", err)
		return
	}

	defer session.Close()
	fmt.Println("Bot running, Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// print out every event as it occurs, useful for debugging.
func eventPrinter(s *discordgo.Session, event interface{}) {
	switch e := event.(type) {
	default:
		fmt.Printf("event: %T\n", e)
	}
}
