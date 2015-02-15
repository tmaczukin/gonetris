package main

import ()

type Player struct {
	Board    *Board
	Nickname string
	Host     string
}

var (
	Players     []*Player
	MyPlayer    *Player
	PlayersList = make(chan *Player)
	PlayerEvent = make(chan *Player)
)

func init() {
	Players = make([]*Player, 0, opts.Players)
}

func NewPlayer() *Player {

	var player Player
	player.Board = NewBoard(5, 5)
	PlayersList <- &player

	return &player
}

func HandlePlayers() {

	MyPlayer = NewPlayer()
	Players = append(Players, MyPlayer)
	defer Wg.Done()

	for player := range PlayerEvent {

		player.Board.Draw()
	}

}
