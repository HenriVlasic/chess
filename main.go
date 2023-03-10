package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Piece int

const (
	Empty Piece = iota
	wP
	wN
	wB
	wR
	wQ
	wK
	bP
	bN
	bB
	bR
	bQ
	bK
)

type Board [8][8]Piece

func (b *Board) MovePiece(from, to string) error {
	fx, fy := parseCoordinates(from)
	tx, ty := parseCoordinates(to)

	if !b.isValidMove(fx, fy, tx, ty) {
		return errors.New("invalid move")
	}

	b[tx][ty] = b[fx][fy]
	b[fx][fy] = Empty

	return nil
}

func (b *Board) isValidMove(fx, fy, tx, ty int) bool {
	// Pawns Logic
	// Black Pawn
	if b[fx][fy] == bP {
		// Check if the pawn is trying to capture a piece
		if b[tx][ty] != Empty && (b[fx][fy] == bP || b[fx][fy] == bR || b[fx][fy] == bN || b[fx][fy] == bB || b[fx][fy] == bK || b[fx][fy] == bQ) && (b[tx][ty] == wP || b[tx][ty] == wR || b[tx][ty] == wN || b[tx][ty] == wB || b[tx][ty] == wK || b[tx][ty] == wQ) {
			if abs(fx-tx) != 1 || abs(fy-ty) != 1 {
				// Black pawns can only capture pieces that are 1 space forward and diagonal
				return false
			}
		} else {
			// Check if the pawn is moving forward
			if ty != fy {
				// Pawns can only move straight forward
				return false
			}
			if tx < fx {
				// Pawns can only move forward, not backward
				return false
			}
			if tx-fx > 2 {
				// Pawns can only move 1 or 2 spaces from their starting position
				return false
			}
			if tx-fx == 2 && fx != 1 {
				// Pawns can only move 2 spaces from their starting position
				return false
			}
			if b[tx][ty] != Empty {
				// Pawns can only move to an empty space
				return false
			}
		}
		return true
	}
	// White Pawn
	if b[fx][fy] == wP {
		// Check if the pawn is trying to capture a piece
		if b[tx][ty] != Empty && (b[fx][fy] == wP || b[fx][fy] == wR || b[fx][fy] == wN || b[fx][fy] == wB || b[fx][fy] == wK || b[fx][fy] == wQ) && (b[tx][ty] == bP || b[tx][ty] == bR || b[tx][ty] == bN || b[tx][ty] == bB || b[tx][ty] == bK || b[tx][ty] == bQ) {
			if abs(fx-tx) != 1 || abs(fy-ty) != 1 {
				// Black pawns can only capture pieces that are 1 space forward and diagonal
				return false
			}
		} else {
			// Check if the pawn is moving forward
			if ty != fy {
				// Pawns can only move straight forward
				return false
			}
			if tx > fx {
				// Pawns can only move forward, not backward
				return false
			}
			if fx-tx > 2 {
				// Pawns can only move 1 or 2 spaces from their starting position
				return false
			}
			if fx-tx == 2 && fx != 6 {
				// Pawns can only move 2 spaces from their starting position
				return false
			}
			if b[tx][ty] != Empty {
				// Pawns can only move to an empty space or capture an opponent's piece
				return false
			}
		}
		return true
	}
	// Rooks Logic
	// Black Rook
	if b[fx][fy] == bR {
		if b[tx][ty] != Empty && (b[tx][ty] == bP || b[tx][ty] == bR || b[tx][ty] == bN || b[tx][ty] == bB || b[tx][ty] == bK || b[tx][ty] == bQ) {

			return false

		} else {
			if ty != fy && tx != fx {
				// Rooks can only move in straight lines
				return false
			}
		}

		return true

	}

	// White Rook
	if b[fx][fy] == wR {
		if b[tx][ty] != Empty && (b[tx][ty] == wP || b[tx][ty] == wR || b[tx][ty] == wN || b[tx][ty] == wB || b[tx][ty] == wK || b[tx][ty] == wQ) {

			return false

		} else {
			if ty != fy && tx != fx {
				// Rooks can only move in straight lines
				return false
			}
		}

		return true

	}

	// Other checks for other piece types go here

	// Empty Squares
	// Checking if it is empty
	if b[fx][fy] == Empty {
		return false
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseCoordinates(s string) (int, int) {
	col := s[0] - 'a'
	row := 8 - (s[1] - '0')
	return int(row), int(col)
}

func printBoard(b Board) {
	fmt.Println("  ?????????????????????????????????????????????????????????")
	for i := 0; i < 8; i++ {
		fmt.Print(8 - i)
		fmt.Print(" ???")
		for j := 0; j < 8; j++ {
			fmt.Print(" ")
			if isWhiteSquare(b, i, j) {
				fmt.Print("")
			} else {
				fmt.Print("")
			}
			switch b[i][j] {
			case bP:
				fmt.Print("???")
			case bN:
				fmt.Print("???")
			case bB:
				fmt.Print("???")
			case bR:
				fmt.Print("???")
			case bQ:
				fmt.Print("???")
			case bK:
				fmt.Print("???")
			case wP:
				fmt.Print("???")
			case wN:
				fmt.Print("???")
			case wB:
				fmt.Print("???")
			case wR:
				fmt.Print("???")
			case wQ:
				fmt.Print("???")
			case wK:
				fmt.Print("???")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println(" ???")
	}
	fmt.Println("  ?????????????????????????????????????????????????????????")
	fmt.Println("    a b c d e f g h")

}

func isWhiteSquare(b Board, row, col int) bool {
	return b[row][col] == Empty && (row+col)%2 == 0
}

func parseInput(input string) (string, string, error) {
	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		return "", "", errors.New("invalid input")
	}
	return parts[0], parts[1], nil
}

func main() {
	b := Board{
		{bR, bN, bB, bQ, bK, bB, bN, bR},
		{bP, bP, bP, bP, bP, bP, bP, bP},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{wP, wP, wP, wP, wP, wP, wP, wP},
		{wR, wN, wB, wQ, wK, wB, wN, wR},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "quit" {
			break
		}
		from, to, err := parseInput(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = b.MovePiece(from, to)
		if err != nil {
			fmt.Println(err)
			continue
		}
		printBoard(b)
	}
}
