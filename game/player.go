package game

import (
	"fmt"
)

type Player struct {
	Id      int
	Symbol  int
	IsHuman bool
}

type Board struct {
	BoardLayout   [][]int
	MovesCount    int
	GameCondition int
	Size          int
}

type Actions []Board

func NewBoard(size int) *Board {
	NewBoard := new(Board)
	NewBoard.BoardLayout = make([][]int, size)
	for i := range NewBoard.BoardLayout {
		NewBoard.BoardLayout[i] = make([]int, size)
		for j := range NewBoard.BoardLayout[i] {
			NewBoard.BoardLayout[i][j] = 0
		}
	}

	NewBoard.Size = size
	NewBoard.GameCondition = OnGoing
	NewBoard.MovesCount = 0

	return NewBoard
}

func NewPlayer(id, symbol int, IsHuman bool) *Player {
	p := new(Player)
	p.Id = id
	p.Symbol = symbol
	p.IsHuman = IsHuman

	return p
}

func (player *Player) MakePlay(originalBoard [][]int) (int, int) {
	fmt.Printf("next player to place a mark is: %v\n", player.Id)

	var row, column int

	fmt.Printf("where to place a %v? (input row then column, separated by space)\n> ", player.Id)

	fmt.Scanf("%d %d", &row, &column)

	return row, column
}

func (b Board) MaxMin(symbol int) Board {
	board := deepCopy(b)

	if b.GameCondition != OnGoing {
		return b
	}

	var actions Actions
	for i := range board.BoardLayout {
		for j := range board.BoardLayout[i] {
			if board.BoardLayout[i][j] == Empty {
				boardcopy := deepCopy(board)
				boardcopy.BoardLayout[i][j] = symbol
				boardcopy.CheckVictory()
				boardcopy.MovesCount++
				actions = append(actions, boardcopy)
			}
		}
	}

	if symbol == Circle {
		for i, action := range actions {
			actions[i].GameCondition = action.MaxMin(Cross).GameCondition
		}
		var min Board
		for i, action := range actions {
			if i == 0 {
				min = deepCopy(action)
			}
			if action.GameCondition < min.GameCondition {
				min = deepCopy(action)
			}
		}
		return min

	} else if symbol == Cross {
		for i, action := range actions {
			actions[i].GameCondition = action.MaxMin(Circle).GameCondition
		}
		var max Board
		for i, action := range actions {
			if i == 0 {
				max = deepCopy(action)
			}
			if action.GameCondition > max.GameCondition && action.GameCondition < OnGoing {
				max = deepCopy(action)
			}
		}
		return max
	}

	return Board{}
}

func (board *Board) CheckVictory() {
	var symbol int

	//rows
	for rowIndex := range board.BoardLayout {
		for columnIndex := range board.BoardLayout[rowIndex] {
			if columnIndex == 0 {
				symbol = board.BoardLayout[rowIndex][columnIndex]
			}

			if board.BoardLayout[rowIndex][columnIndex] != symbol {
				break
			}

			if columnIndex == board.Size-1 {
				switch symbol {
				case Circle:
					board.GameCondition = CircleWin
					return
				case Cross:
					board.GameCondition = CrossWin
					return
				}
			}
		}

	}

	//columns
	for rowIndex := range board.BoardLayout {
		for columnIndex := range board.BoardLayout[rowIndex] {
			if columnIndex == 0 {
				symbol = board.BoardLayout[columnIndex][rowIndex]
			}

			if board.BoardLayout[columnIndex][rowIndex] != symbol {
				break
			}
			if columnIndex == board.Size-1 {
				switch symbol {
				case Circle:
					board.GameCondition = CircleWin
					return
				case Cross:
					board.GameCondition = CrossWin
					return
				}
			}
		}

	}

	//diagonal
	for rowIndex := range board.BoardLayout {
		if rowIndex == 0 {
			symbol = board.BoardLayout[rowIndex][rowIndex]
		}
		if board.BoardLayout[rowIndex][rowIndex] != symbol {
			break
		}
		if rowIndex == board.Size-1 {
			switch symbol {
			case Circle:
				board.GameCondition = CircleWin
				return
			case Cross:
				board.GameCondition = CrossWin
				return
			}
		}
	}

	for rowIndex := range board.BoardLayout {
		if rowIndex == 0 {
			symbol = board.BoardLayout[rowIndex][(board.Size-1)-rowIndex]
		}
		if board.BoardLayout[rowIndex][(board.Size-1)-rowIndex] != symbol {
			break
		}
		if rowIndex == board.Size-1 {
			switch symbol {
			case Circle:
				board.GameCondition = CircleWin
				return
			case Cross:
				board.GameCondition = CrossWin
				return
			}
		}
	}

	for rowIndex := range board.BoardLayout {
		for columnIndex := range board.BoardLayout[rowIndex] {
			if board.BoardLayout[rowIndex][columnIndex] == Empty {
				board.GameCondition = OnGoing
				return
			}
		}
	}

	board.GameCondition = Draw
}

func deepCopy(board Board) Board {
	deepCopied := board

	deepCopied.BoardLayout = make([][]int, len(board.BoardLayout))
	for i := range board.BoardLayout {
		deepCopied.BoardLayout[i] = make([]int, len(board.BoardLayout[i]))
		copy(deepCopied.BoardLayout[i], board.BoardLayout[i])
	}

	return deepCopied
}
