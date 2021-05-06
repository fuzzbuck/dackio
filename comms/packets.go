package comms

import "dackio/game"

const (
	// CLIENT -> SERVER
	HI = iota
	GIVE_ME_WORLD_DATA
	GIVE_ME_PLAYER_DATA

	// SERVER -> CLIENT
	WORLD_DATA_CHUNK 	/* Chunk containing WorldData struct */
	PLAYER_DATA_CHUNK	/* Chunk containing PlayerData struct */

	// PLAYER ACTIONS
	PLAYER_CHAT			/* Player sent a message on a channel */
	PLAYER_HOP			/* Hop to a specified Position */
	PLAYER_DISAPPEAR	/* Make the player jump off into the void and disappear */
	PLAYER_APPEAR		/* New player joined the server, make him fall from the sky */
)

var PacketHandler = map[byte]func(ws *Client){
	HI: func(c *Client) {
		GlobalHub.SendTo(c, []byte("hello back"))
	},
}

type Player struct {
	Username	string
	Position	game.Vector2
}

type PlayerDataChunk struct {
	ChunkSize	int
	Players		[]Player
}

type WorldInfo struct {
	Size		game.Vector2
}

type WorldChunk struct {
	ChunkSize	int
	Blocks		[]game.Block
}