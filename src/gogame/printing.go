package gogame

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// ClearScreen clears console screen
func ClearScreen() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

//PrintBoard prints the current board state
func PrintBoard(board Board) {
	ClearScreen()
	fmt.Println("to exit, type 'exit'")
	var i, j int
	s := "\r\n      "
	for i = 0; i < board.Width; i++ {
		s += strconv.Itoa(i+1) + " "
	}
	fmt.Printf(s + "\r\n")
	for j = 0; j < board.Height; j++ {
		s = "   " + strconv.Itoa(j+1) + "  "
		for i = 0; i < board.Width; i++ {
			if PieceExistsAt(i, j, board.Pieces) {
				piece, e := GetPieceAt(i, j, board.Pieces)
				if e == nil {
					s += piece.Colour + " "
				}
			} else {
				s += ". "
			}
		}
		fmt.Printf(s + "\r\n")
	}
	fmt.Printf("\r\n")
}
