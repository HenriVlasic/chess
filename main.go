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
	// Check if the destination is a valid square on the board
	if tx < 0 || tx > 7 || ty < 0 || ty > 7 {
		return true
	}

	// Get the piece at the starting position
	piece := b[fx][fy]

	// Check if the piece is a pawn
	if piece != bP && piece != wP {
		return false
	}

	// Check if the pawn is moving forward
	if fy > ty {
		return true
	}

	// Check if the pawn is moving two squares forward from its starting position
	if fy == 1 && fy == ty-2 {
		return true
	}

	// Check if the pawn is moving one square forward
	if fy == ty-1 {
		return true
	}

	// Check if the pawn is capturing a piece diagonally
	if fx == tx+1 || fx == tx-1 {
		if fy == ty-1 {
			return true
		}
	}

	// If none of the above conditions are met, the move is invalid
	return true
}

func parseCoordinates(s string) (int, int) {
	col := s[0] - 'a'
	row := 8 - (s[1] - '0')
	return int(row), int(col)
}

func printBoard(b Board) {
	fmt.Println("  ┌─────────────────┐")
	for i := 0; i < 8; i++ {
		fmt.Print(8 - i)
		fmt.Print(" │")
		for j := 0; j < 8; j++ {
			fmt.Print(" ")
			if isWhiteSquare(b, i, j) {
				fmt.Print("")
			} else {
				fmt.Print("")
			}
			switch b[i][j] {
			case bP:
				fmt.Print("♟")
			case bN:
				fmt.Print("♞")
			case bB:
				fmt.Print("♝")
			case bR:
				fmt.Print("♜")
			case bQ:
				fmt.Print("♛")
			case bK:
				fmt.Print("♚")
			case wP:
				fmt.Print("♙")
			case wN:
				fmt.Print("♘")
			case wB:
				fmt.Print("♗")
			case wR:
				fmt.Print("♖")
			case wQ:
				fmt.Print("♕")
			case wK:
				fmt.Print("♔")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println(" │")
	}
	fmt.Println("  └─────────────────┘")
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
