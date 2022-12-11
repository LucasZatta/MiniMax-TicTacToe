package game

import (
	"errors"
	"fmt"
)

const (
	Empty = iota
	Cross
	Circle
)

const (
	CircleWin = iota
	Draw
	CrossWin
	OnGoing
)

var ErrPosNotVacant error = errors.New("posicao ocupada")
var ErrPosNotValid error = errors.New("posicao invalida")

type GameState struct {
	Board            Board
	Players          []Player
	CurrentPlayer    Player
	CurrentPlayerNum int
}

func NewGameState(size int, players []Player) *GameState {
	NewGame := new(GameState)
	NewGame.Players = players
	NewGame.Board = *NewBoard(size)

	NewGame.CurrentPlayerNum = 0
	NewGame.CurrentPlayer = players[0]

	return NewGame
}

func (state *GameState) DrawBoard() {
	if state.Board.GameCondition != OnGoing {
		switch state.Board.GameCondition {
		case CircleWin:
			fmt.Println("Circle Wins!")
		case CrossWin:
			fmt.Println("Cross Wins!")
		case Draw:
			fmt.Println("Oh no! Its a Draw!!!")
		}
	}
	for i, row := range state.Board.BoardLayout {
		for j, square := range row {
			fmt.Print(" ")
			switch square {
			case Empty:
				fmt.Print(" ")
			case Cross:
				fmt.Print("X")
			case Circle:
				fmt.Print("O")
			}
			if j != len(row)-1 {
				fmt.Print(" |")
			}
		}
		if i != len(state.Board.BoardLayout)-1 {
			fmt.Print("\n------------")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n \n")
}

func (state *GameState) Start() {
	for state.Board.GameCondition == OnGoing {
		state.DrawBoard()
		state.Action()
		state.Board.CheckVictory()
	}

	state.DrawBoard()
}

func (state *GameState) Action() {
	if state.CurrentPlayer.IsHuman {
		row, column := state.CurrentPlayer.MakePlay(state.Board.BoardLayout)
		err := state.placeSymbol(row, column)
		if err != nil {
			switch err {
			case ErrPosNotVacant:
				fmt.Println("A posicao selecionada nao esta disponivel")
				return
			case ErrPosNotValid:
				fmt.Println("A posicao selecionado nao eh valida.")
				return
			}
		}

	} else {
		// newBoard := state.Board.MaxMin(state.CurrentPlayer.Symbol)
		// fmt.Println(newBoard)
		state.Board.BoardLayout = state.Board.MaxMin(state.CurrentPlayer.Symbol).BoardLayout
		// state.Board.GameCondition = newBoard.GameCondition
	}

	state.CurrentPlayerNum = 1 - state.CurrentPlayerNum
	state.CurrentPlayer = state.Players[state.CurrentPlayerNum]
}

func (state *GameState) placeSymbol(row, column int) error {
	if row <= 0 || row > len(state.Board.BoardLayout) || column <= 0 || column > len(state.Board.BoardLayout) {
		return ErrPosNotValid
	}

	if state.Board.BoardLayout[row-1][column-1] != Empty {
		return ErrPosNotVacant
	}

	state.Board.BoardLayout[row-1][column-1] = state.CurrentPlayer.Symbol
	return nil
}
