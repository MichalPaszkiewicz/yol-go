package gogame

import "math"

func selectPieces(pieces []Piece, test func(Piece) bool) (ret []Piece) {
	for _, p := range pieces {
		if test(p) {
			ret = append(ret, p)
		}
	}
	return
}

func selectPiecesOfColour(pieces []Piece, colour string) []Piece {
	colourTest := func(p Piece) bool { return colour == p.Colour }
	return selectPieces(pieces, colourTest)
}

func absDiff(x int, y int) int {
	i, j := float64(x), float64(y)
	k := math.Abs(i - j)
	return int(k)
}

func piecesAreAdjacent(p1 Piece, p2 Piece) bool {
	if absDiff(p1.I, p2.I) == 1 && p1.J-p2.J == 0 || absDiff(p1.J, p2.J) == 1 && p1.I-p2.I == 0 {
		return true
	}
	return false
}

func selectAdjacentPieces(pieces []Piece, x int, y int) []Piece {
	adjacentTest := func(p Piece) bool { return piecesAreAdjacent(p, Piece{I: x, J: y}) }
	return selectPieces(pieces, adjacentTest)
}

func pieceIsSurrounded(piece Piece, board Board) bool {
	surroundersNeeded := 4
	if piece.I == 1 || piece.I == board.Width {
		surroundersNeeded--
	}
	if piece.J == 1 || piece.J == board.Height {
		surroundersNeeded--
	}
	return len(selectAdjacentPieces(board.Pieces, piece.I, piece.J)) >= surroundersNeeded
}

func groupIsSurrounded(pieces []Piece, board Board) bool {
	for i := 0; i < len(pieces); i++ {
		if !pieceIsSurrounded(pieces[i], board) {
			return false
		}
	}
	return true
}

func containsPieceAt(pieces []Piece, x int, y int) bool {
	for i := 0; i < len(pieces); i++ {
		if pieces[i].I == x && pieces[i].J == y {
			return true
		}
	}
	return false
}

//GetGroup gets group of pieces connected to area x,y
func GetGroup(board Board, x int, y int, colour string) []Piece {
	group := make([]Piece, 0)

	friendlyPieces := selectPiecesOfColour(board.Pieces, colour)

	searchPieces := []Piece{Piece{I: x, J: y, Colour: colour}}
	search := true

	for search {
		basePiece := searchPieces[0]
		adjacentPieces := selectAdjacentPieces(friendlyPieces, basePiece.I, basePiece.J)
		for i := 0; i < len(adjacentPieces); i++ {
			ap := adjacentPieces[i]
			if !containsPieceAt(searchPieces, ap.I, ap.J) && !containsPieceAt(group, ap.I, ap.J) {
				searchPieces = append(searchPieces, ap)
			}
		}
		group = append(group, basePiece)
		if len(searchPieces) > 1 {
			searchPieces = searchPieces[1:]
		} else {
			search = false
		}
	}

	return group
}
