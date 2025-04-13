package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
	"strconv"
)

func (bot *Bot) messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	//TODO: SLASH commands for which channel is the announcement channel.
	if message.ChannelID != "1360458157939494953" {
		return
	}
	addPost(message.Message, bot.db)

}

func (bot *Bot) joinServer(session *discordgo.Session, guild *discordgo.GuildCreate) {
	ID, err := strconv.ParseUint(guild.ID, 10, 64)
	if err != nil {
		fmt.Println("Failed to parse server ID, ", err)
		return
	}
	Name := guild.Guild.Name
	_, err = bot.db.Query("INSERT INTO guild (id,name) VALUES(?,?)", ID, Name)
	if err != nil {
		// https://dev.to/bitecode/catch-error-when-using-sqlite-in-golang-58nn
		// to get our error code we need to type assert to the sqlite error type
		// which has Code() defined that gives us our sqlite3 error codes.
		if sqlErr, ok := err.(*sqlite.Error); ok {
			//We already have an entry for this guild in our guild table
			if sqlErr.Code() == sqlite3.SQLITE_CONSTRAINT_PRIMARYKEY {
				return
			}
		}
		fmt.Printf("Failed to insert new guild %s with id %d ,%s\n", Name, ID, err)
		return
	}
	fmt.Printf("Joined a new guild! Name: %s , ID: %d\n", Name, ID)
}
