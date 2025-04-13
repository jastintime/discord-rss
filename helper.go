package main

import (
	"database/sql"
	"fmt"
	"github.com/bwmarrin/discordgo"
	_ "modernc.org/sqlite"
	"strconv"
)

func addPost(post *discordgo.Message, db *sql.DB) {
	if db == nil {
		fmt.Println("uninitialized database")
		return
	}
	// convert https://github.com/bwmarrin/discordgo/issues/380
	// we are given IDs but they are unique uint64 snowflakes, cast them for the database.
	ID, err := strconv.ParseUint(post.ID, 10, 64)
	if err != nil {
		fmt.Println("Failed to parse ID while trying to add post: ", err)
		return
	}
	guildID, err := strconv.ParseUint(post.GuildID, 10, 64)
	if err != nil {
		fmt.Println("Failed to parse guildID while trying to add post: ", err)
		return
	}
	author := post.Author.GlobalName
	content := post.Content
	//TODO: handle embeds
	_, err = db.Exec("INSERT INTO post (id,guild_id,author,content) VALUES(?,?,?,?)", ID, guildID, author, content)
	if err != nil {
		fmt.Println("Failed to insert post into database ", err)
		return
	}
}
