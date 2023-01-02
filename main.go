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

	if !isValidMove(fx, fy, tx, ty) {
		return errors.New("invalid move")
	}

	// Update the board with the new position of the piece
	b[tx][ty] = b[fx][fy]
	b[fx][fy] = Empty

	return nil
}

func isValidMove(fx, fy, tx, ty int) bool {
	// Check if the move is legal according to the rules of chess
	// (e.g. pawns can only move forward, knights can jump over pieces, etc.)
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
		// Parse the input and make a move on the board
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
		// Print the updated board
		printBoard(b)
	}
}
