package callbacks

import (
	"NestedScribbleBenchmark/noughtsandcrosses/impl"
	"NestedScribbleBenchmark/noughtsandcrosses/results/standardstrategy"
	"fmt"
)

type StandardStrategy_P_Env interface {
	Done() standardstrategy.P_Result
}

type StandardStrategyPState struct {
	Move   int
	Board  impl.Board
	Player impl.Player
}

func (s *StandardStrategyPState) Done() standardstrategy.P_Result {
	var move int
	for {
		fmt.Println("Board state:")
		fmt.Println(s.Board)
		fmt.Printf("Enter move for player '%s': ", s.Player)
		_, err := fmt.Scanln(&move)
		if err == nil {
			if move < impl.BoardSize && s.Board[move] == impl.Empty {
				break
			}
			fmt.Println("Invalid input: Please enter the idx of a free space on the board")
		}
	}
	return standardstrategy.P_Result{
		Move: move,
	}
}
