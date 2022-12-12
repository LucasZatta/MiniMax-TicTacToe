package main

import (
	"maxmintictac/game"
)

func main() {

	players := make([]game.Player, 2)
	players[0] = *game.NewPlayer(1, game.Cross, true)
	players[1] = *game.NewPlayer(2, game.Circle, false)
	tictacGame := game.NewGameState(3, players)

	tictacGame.Start()
}
