package minimax_test

import (
	"github.com/kkrull/gosandbox/minimax"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ComputerPlayer", func() {
	Describe("#PickMove", func() {
		It("picks a move in an unfinished game", func() {
			opponent := minimax.Player("Opponent")
			computer := minimax.NewComputerPlayer(opponent)

			game := &GameWithKnownStates{}
			leftTree := &GameWithKnownStates{}
			game.AddKnownState(minimax.Move("Left"), leftTree)
			leftTree.AddKnownState(minimax.Move("Draw"), &GameWithKnownStates{isOver: true})
			leftTree.AddKnownState(minimax.Move("Computer Wins"), &GameWithKnownStates{isOver: true, winner: computer})

			rightTree := &GameWithKnownStates{}
			game.AddKnownState(minimax.Move("Right"), rightTree)
			rightTree.AddKnownState(minimax.Move("Draw"), &GameWithKnownStates{isOver: true})
			rightTree.AddKnownState(minimax.Move("Computer Loses"), &GameWithKnownStates{isOver: true, winner: opponent})

			Expect(computer.PickMove(game)).To(Equal(minimax.Move("Left")))
		})
	})
})
