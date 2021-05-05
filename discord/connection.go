package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	Token	string
	sess  	*discordgo.Session
)

func OpenConnection() {
	fmt.Printf("connecting w/ token '%s'\n", Token)
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		panic(err)
	}
	err = dg.Open()
	if err != nil {
		panic(err)
	}

	sess = dg
}