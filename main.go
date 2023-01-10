package main

import (
	"fmt"
	"math"
)

// Piece represents a chess piece
type Piece struct {
	Name     string
	Position string
	Color    string
	Symbol   string
}

// Board represents a chess board
type Board [8][8]string

// NewBoard creates a new chess board
func NewBoard() *Board {
	board := Board{}
	return &board
}

// SetPiece sets a chess piece on the board
func (b *Board) SetPiece(piece Piece) {
	pos := piece.Position
	ind, ok := posToInd[pos]
	if !ok {
		return
	}
	b[ind[0]][ind[1]] = piece.Symbol
}

// PrintBoard prints the chess board
func (b *Board) PrintBoard() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%s ", b[i][j])
		}
		fmt.Println()
	}
}

// posToInd maps chess board positions to indices
var posToInd = map[string][]int{
	"a1": {7, 0}, "b1": {7, 1}, "c1": {7, 2}, "d1": {7, 3}, "e1": {7, 4}, "f1": {7, 5}, "g1": {7, 6}, "h1": {7, 7},
	"a2": {6, 0}, "b2": {6, 1}, "c2": {6, 2}, "d2": {6, 3}, "e2": {6, 4}, "f2": {6, 5}, "g2": {6, 6}, "h2": {6, 7},
	"a3": {5, 0}, "b3": {5, 1}, "c3": {5, 2}, "d3": {5, 3}, "e3": {5, 4}, "f3": {5, 5}, "g3": {5, 6}, "h3": {5, 7},
	"a4": {4, 0}, "b4": {4, 1}, "c4": {4, 2}, "d4": {4, 3}, "e4": {4, 4}, "f4": {4, 5}, "g4": {4, 6}, "h4": {4, 7},
	"a5": {3, 0}, "b5": {3, 1}, "c5": {3, 2}, "d5": {3, 3}, "e5": {3, 4}, "f5": {3, 5}, "g5": {3, 6}, "h5": {3, 7},
	"a6": {2, 0}, "b6": {2, 1}, "c6": {2, 2}, "d6": {2, 3}, "e6": {2, 4}, "f6": {2, 5}, "g6": {2, 6}, "h6": {2, 7},
	"a7": {1, 0}, "b7": {1, 1}, "c7": {1, 2}, "d7": {1, 3}, "e7": {1, 4}, "f7": {1, 5}, "g7": {1, 6}, "h7": {1, 7},
	"a8": {0, 0}, "b8": {0, 1}, "c8": {0, 2}, "d8": {0, 3}, "e8": {0, 4}, "f8": {0, 5}, "g8": {0, 6}, "h8": {0, 7},
}

func main() {
	// Creating White Pieces
	whiteKing := Piece{Name: "king", Position: "e1", Color: "white", Symbol: "♔"}
	whiteQueen := Piece{Name: "queen", Position: "d1", Color: "white", Symbol: "♕"}
	whiteRooks := [2]Piece{
		{Name: "rook", Position: "a1", Color: "white", Symbol: "♖"},
		{Name: "rook", Position: "h1", Color: "white", Symbol: "♖"},
	}
	whiteKnights := [2]Piece{
		{Name: "knight", Position: "b1", Color: "white", Symbol: "♘"},
		{Name: "knight", Position: "g1", Color: "white", Symbol: "♘"},
	}
	whiteBishops := [2]Piece{
		{Name: "bishop", Position: "c1", Color: "white", Symbol: "♗"},
		{Name: "bishop", Position: "f1", Color: "white", Symbol: "♗"},
	}
	whitePawns := [8]Piece{
		{Name: "pawn", Position: "a2", Color: "white", Symbol: "♙"},
		{Name: "pawn", Position: "b2", Color: "white", Symbol: "♙"},
		{Name: "pawn", Position: "c2", Color: "white", Symbol: "♙"},
		{Name: "pawn", Position: "d2", Color: "white", Symbol: "♙"},
		{Name: "pawn", Position: "e2", Color: "white", Symbol: "♙"},
		{Name: "pawn", Position: "f2", Color: "white", Symbol: "♙"},
		{Name: "pawn", Position: "g2", Color: "white", Symbol: "♙"},
		{Name: "pawn", Position: "h2", Color: "white", Symbol: "♙"},
	}

	// Creating Black Pieces
	blackKing := Piece{Name: "king", Position: "e8", Color: "black", Symbol: "♚"}
	blackQueen := Piece{Name: "queen", Position: "d8", Color: "black", Symbol: "♛"}
	blackRooks := [2]Piece{
		{Name: "rook", Position: "a8", Color: "black", Symbol: "♜"},
		{Name: "rook", Position: "h8", Color: "black", Symbol: "♜"},
	}
	blackKnights := [2]Piece{
		{Name: "knight", Position: "b8", Color: "black", Symbol: "♞"},
		{Name: "knight", Position: "g8", Color: "black", Symbol: "♞"},
	}
	blackBishops := [2]Piece{
		{Name: "bishop", Position: "c8", Color: "black", Symbol: "♝"},
		{Name: "bishop", Position: "f8", Color: "black", Symbol: "♝"},
	}
	blackPawns := [8]Piece{
		{Name: "pawn", Position: "a7", Color: "black", Symbol: "♟"},
		{Name: "pawn", Position: "b7", Color: "black", Symbol: "♟"},
		{Name: "pawn", Position: "c7", Color: "black", Symbol: "♟"},
		{Name: "pawn", Position: "d7", Color: "black", Symbol: "♟"},
		{Name: "pawn", Position: "e7", Color: "black", Symbol: "♟"},
		{Name: "pawn", Position: "f7", Color: "black", Symbol: "♟"},
		{Name: "pawn", Position: "g7", Color: "black", Symbol: "♟"},
		{Name: "pawn", Position: "h7", Color: "black", Symbol: "♟"},
	}

	board := NewBoard()

	// Place white pieces
	board.SetPiece(whiteKing)
	board.SetPiece(whiteQueen)
	for _, rook := range whiteRooks {
		board.SetPiece(rook)
	}
	for _, knight := range whiteKnights {
		board.SetPiece(knight)
	}
	for _, bishop := range whiteBishops {
		board.SetPiece(bishop)
	}
	for _, pawn := range whitePawns {
		board.SetPiece(pawn)
	}

	// Place black pieces
	board.SetPiece(blackKing)
	board.SetPiece(blackQueen)
	for _, rook := range blackRooks {
		board.SetPiece(rook)
	}
	for _, knight := range blackKnights {
		board.SetPiece(knight)
	}
	for _, bishop := range blackBishops {
		board.SetPiece(bishop)
	}
	for _, pawn := range blackPawns {
		board.SetPiece(pawn)
	}

	// Print board
	board.PrintBoard()
}

// IsLegalMove determines if a move is legal according to the rules of chess
func (b *Board) IsLegalMove(piece Piece, newPos string) bool {
	// Check that the new position is on the board
	ind, ok := posToInd[newPos]
	if !ok {
		return false
	}

	// Check that the piece is not already at the new position
	if piece.Position == newPos {
		return false
	}

	// Get the current position of the piece
	currInd, _ := posToInd[piece.Position]
	currRow, currCol := currInd[0], currInd[1]
	newRow, newCol := ind[0], ind[1]

	// Check that the move is legal for the specific piece type
	switch piece.Name {
	case "pawn":
		// Pawns can only move forward
		if piece.Color == "white" {
			if newRow >= currRow {
				return false
			}
		} else {
			if newRow <= currRow {
				return false
			}
		}
		// Pawns can move one square forward, or two squares on their first move
		if math.Abs(float64(newRow-currRow)) > 2 {
			return false
		}
		// Pawns can only capture diagonally
		if newCol != currCol {
			if math.Abs(float64(newCol-currCol)) != 1 {
				return false
			}
			if b[newRow][newCol] == "" {
				return false
			}
		} else {
			if b[newRow][newCol] != "" {
				return false
			}
		}
	case "knight":
		// Knights can move in an "L" shape
		if math.Abs(float64(newRow-currRow)) == 2 {
			if math.Abs(float64(newCol-currCol)) != 1 {
				return false
			}
		} else if math.Abs(float64(newCol-currCol)) == 2 {
			if math.Abs(float64(newRow-currRow)) != 1 {
				return false
			}
		} else {
			return false
		}
	case "bishop":
		// Bishops can only move diagonally
		if math.Abs(float64(newRow-currRow)) != math.Abs(float64(newCol-currCol)) {
			return false
		}
		// Check for pieces blocking the path
		if newRow > currRow {
			if newCol > currCol {
				for i := currRow + 1; i < newRow; i++ {
					if b[i][i] != "" {

						return false
					}
				}
			} else {
				for i, j := currRow+1, currCol-1; i < newRow; i, j = i+1, j-1 {
					if b[i][j] != "" {
						return false
					}
				}
			}
		} else {
			if newCol > currCol {
				for i, j := currRow-1, currCol+1; i > newRow; i, j = i-1, j+1 {
					if b[i][j] != "" {
						return false
					}
				}
			} else {
				for i, j := currRow-1, currCol-1; i > newRow; i, j = i-1, j-1 {
					if b[i][j] != "" {
						return false
					}
				}
			}
		}
	case "queen":
		// Queens can move diagonally, horizontally, or vertically
		if newRow != currRow && newCol != currCol {
			if math.Abs(float64(newRow-currRow)) != math.Abs(float64(newCol-currCol)) {
				return false
			}
			// Check for pieces blocking the path
			if newRow > currRow {
				if newCol > currCol {
					for i, j := currRow+1, currCol+1; i < newRow; i, j = i+1, j+1 {
						if b[i][j] != "" {
							return false
						}
					}
				} else {
					for i, j := currRow+1, currCol-1; i < newRow; i, j = i+1, j-1 {
						if b[i][j] != "" {
							return false
						}
					}
				}
			} else {
				if newCol > currCol {
					for i, j := currRow-1, currCol+1; i > newRow; i, j = i-1, j+1 {
						if b[i][j] != "" {
							return false
						}
					}
				} else {
					for i, j := currRow-1, currCol-1; i > newRow; i, j = i-1, j-1 {
						if b[i][j] != "" {
							return false
						}
					}
				}
			}
		} else {
			// Check for pieces blocking the path
			if newRow == currRow {
				if newCol > currCol {
					for i := currCol + 1; i < newCol; i++ {
						if b[newRow][i] != "" {
							return false
						}
					}
				} else {
					for i := currCol - 1; i > newCol; i-- {
						if b[newRow][i] != "" {
							return false
						}
					}
				}
			} else {
				if newRow > currRow {
					for i := currRow + 1; i < newRow; i++ {
						if b[i][newCol] != "" {
							return false
						}
					}
				} else {
					for i := currRow - 1; i > newRow; i-- {
						if b[i][newCol] != "" {

							return false
						}
					}
				}
			}
		}
	case "king":
		// Kings can only move one square at a time
		if math.Abs(float64(newRow-currRow)) > 1 || math.Abs(float64(newCol-currCol)) > 1 {
			return false
		}
	default:
		return false
	}

	return true
}
