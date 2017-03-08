package gogame

import (
	"errors"
)

//Piece a white or black piece
type Piece struct {
	Colour string
	I      int
	J      int
}

//Board containing all the pieces in the game
type Board struct {
	Width  int
	Height int
	Pieces []Piece
}

//AddPiece returns a new board with extra piece
func AddPiece(board Board, piece Piece) Board {
	newPieces := make([]Piece, len(board.Pieces)+1)
	newPieces = append(board.Pieces, piece)
	newBoard := Board{board.Width, board.Height, newPieces}
	newBoard = RemoveEnemies(newBoard, piece.I, piece.J, piece.Colour)
	return newBoard
}

// PieceExistsAt tells you if piece is in array
func PieceExistsAt(x int, y int, pieces []Piece) bool {
	for i := 0; i < len(pieces); i++ {
		if pieces[i].I == x+1 && pieces[i].J == y+1 {
			return true
		}
	}
	return false
}

// GetPieceAt gives you a piece from an array
func GetPieceAt(x int, y int, pieces []Piece) (Piece, error) {
	for i := 0; i < len(pieces); i++ {
		if pieces[i].I == x+1 && pieces[i].J == y+1 {
			return pieces[i], nil
		}
	}
	return Piece{}, errors.New("Piece does not exist in this piece array")
}
