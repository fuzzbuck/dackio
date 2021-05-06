package game

var Current	World

type Vector2 struct {
	x			int
	y			int
}

type Block struct {
	Position	Vector2
	BlockId		byte
}

type Unit struct {
	Id			byte
	Username	string
	Position	Vector2
}

type Info struct {
	Name		string
}

type World struct {
	Info		Info
	Blocks		[]Vector2
	Units		[]Unit
}