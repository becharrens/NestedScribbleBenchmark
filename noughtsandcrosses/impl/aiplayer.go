package impl

import "strconv"

type Tile int

const (
	Empty Tile = iota
	P1Move
	P2Move
)

func (t Tile) String() string {
	switch t {
	case Empty:
		return " "
	case P1Move:
		return "X"
	case P2Move:
		return "O"
	default:
		panic("Invalid Tile")
	}
}

type Player int

const (
	P1 Player = iota
	P2
)

func (p Player) String() string {
	if p == P1 {
		return "X"
	}
	return "O"
}

func (p Player) Opponent() Player {
	if p == P1 {
		return P2
	}
	return P1
}

type Outcome int

const (
	Continue = iota
	Win
	Draw
)

type Board []Tile

func (b Board) String() string {
	str := ""
	idx := 0
	for i := 0; i < BoardSide; i++ {
		row := ""
		for j := 0; j < BoardSide; j++ {
			switch b[idx] {
			case Empty:
				row += strconv.Itoa(idx)
			default:
				row += b[idx].String()
			}
			row += " "
			idx++
		}
		row += "\n"
		str += row
	}
	return str
}

type MoveStack []int

func (s *MoveStack) Push(v int) {
	*s = append(*s, v)
}

func (s *MoveStack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

var waysToWin = [][]int{{0, 1, 2},
	{0, 4, 8},
	{0, 3, 6},
	{2, 4, 6},
	{2, 5, 8},
	{3, 4, 5},
	{1, 4, 7},
	{6, 7, 8}}

const (
	BoardSide = 3
	BoardSize = BoardSide * BoardSide
)

func playerMove(player Player) Tile {
	switch player {
	case P1:
		return P1Move
	case P2:
		return P2Move
	}
	panic("Invalid player value")
}

func MakeMove(board Board, player Player, idx int) Board {
	// Assumes idx is empty
	newBoard := make(Board, BoardSize)
	copy(newBoard, board)
	newBoard[idx] = playerMove(player)
	return newBoard
}

func NextMoveIdx(board Board, currIdx int) int {
	// currIdx first non-checked Idx
	for i := currIdx; i < BoardSize; i++ {
		if board[i] == Empty {
			return i
		}
	}
	return -1
}

func BoardOutcome(board Board, player Player) Outcome {
	playerTile := playerMove(player)
	for _, wayToWin := range waysToWin {
		win := true
		for _, idx := range wayToWin {
			if board[idx] != playerTile {
				win = false
				break
			}
		}
		if win {
			return Win
		}
	}
	for i := range board {
		if board[i] == Empty {
			return Continue
		}
	}
	return Draw
}

func ScoreOutcome(outcome Outcome, playerToMove Player) int {
	switch outcome {
	case Draw:
		return 0
	case Win:
		if playerToMove == P1 {
			return 1
		}
		return -1
	}
	panic("Invalid game outcome - should only score final outcomes")
}

func InitMoveScore(playerToMove Player) int {
	if playerToMove == P1 {
		return -2
	}
	return 2
}
