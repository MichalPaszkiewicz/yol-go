package main

import (
	"bufio"
	"fmt"
	"gogame"
	"os"
	"strconv"
	"strings"
)

func getPieceFromArg(arg string, colour string) gogame.Piece {
	trimmed := strings.TrimRight(arg, "\r\n")
	split := strings.Split(trimmed, ",")
	x, e := strconv.Atoi(strings.Trim(split[0], " "))
	if e != nil {
		return gogame.Piece{}
	}
	y, e := strconv.Atoi(strings.Trim(split[1], " "))
	if e != nil {
		return gogame.Piece{}
	}
	return gogame.Piece{Colour: colour, I: x, J: y}
}

func main() {
	var list = []gogame.Piece{}
	reader := bufio.NewReader(os.Stdin)

	gameBoard := gogame.Board{Width: 9, Height: 9, Pieces: list}

	colour := "w"

	gameIsGoing := true
	for gameIsGoing == true {
		gogame.PrintBoard(gameBoard)

		fmt.Print("(" + colour + ") - take your move in format x,y: \r\n")
		text, _ := reader.ReadString('\n')

		if strings.TrimRight(text, "\r\n") == "exit" {
			fmt.Println("exiting...")
			gameIsGoing = false
		}

		newPiece := getPieceFromArg(text, colour)
		gameBoard = gogame.AddPiece(gameBoard, newPiece)

		colour = gogame.GetOppositeColour(colour)
	}
}
