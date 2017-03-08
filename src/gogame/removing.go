package gogame

func removePieceFromList(pieces []Piece, index int) []Piece {
	return append(pieces[:index], pieces[index+1:]...)
}

//RemovePiece removes piece at index
func RemovePiece(board Board, index int) Board {
	newPieces := removePieceFromList(board.Pieces, index)
	newBoard := Board{Width: board.Width, Height: board.Height, Pieces: newPieces}
	return newBoard
}

//RemovePieceAt removes piece from coordinates
func RemovePieceAt(board Board, x int, y int) Board {
	notThisCoordinateFunc := func(p Piece) bool { return p.I != x || p.J != y }
	newPieces := selectPieces(board.Pieces, notThisCoordinateFunc)
	return Board{Width: board.Width, Height: board.Height, Pieces: newPieces}
}

//RemoveGroup removes group from the board
func RemoveGroup(board Board, x int, y int, colour string) Board {
	group := GetGroup(board, x, y, colour)

	if !groupIsSurrounded(group, board) {
		return board
	}

	newPieces := []Piece{}

	for i := 0; i < len(board.Pieces); i++ {
		p := board.Pieces[i]
		if !containsPieceAt(group, p.I, p.J) {
			newPieces = append(newPieces, p)
		}
	}

	return Board{Height: board.Height, Width: board.Width, Pieces: newPieces}
}

//RemoveEnemies remove all enemies from board after move at x, y with colour
func RemoveEnemies(board Board, x int, y int, colour string) Board {

	enemyPieces := selectPiecesOfColour(board.Pieces, GetOppositeColour(colour))
	adjacentEnemies := selectAdjacentPieces(enemyPieces, x, y)

	for i := 0; i < len(adjacentEnemies); i++ {
		enemy := adjacentEnemies[i]
		board = RemoveGroup(board, enemy.I, enemy.J, enemy.Colour)
	}

	return board
}
