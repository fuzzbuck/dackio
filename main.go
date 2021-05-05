package main

import (
	"dackio/comms"
	"dackio/discord"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func parseArgs() {
	flag.StringVar(&discord.Token, "token", "", "Bot token")
	flag.Parse()
}

func main() {
	parseArgs()

	discord.OpenConnection()
	fmt.Println("discord session online")

	comms.StartWSHub()
	fmt.Println("websocket hub online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
