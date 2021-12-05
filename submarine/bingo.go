package submarine

import (
	"strconv"
	"strings"
)

type cell struct {
	isMarked bool
	value    int
}

type Bingo struct {
	NumberDraws []int
	Board       [][]cell
	width       int
	lastNumber  int
	NumTurns    int
}

func NewBingo(input []string) Bingo {
	b := Bingo{}

	drawLine := strings.Split(string(input[0]), ",")
	for _, v := range drawLine {
		value, _ := strconv.Atoi(string(v))
		b.NumberDraws = append(b.NumberDraws, value)
	}

	for _, lineString := range input[1:] {
		if len(string(lineString)) == 0 {
			continue
		}

		lineString := strings.Split(lineString, " ")
		boardLine := []cell{}
		for _, v := range lineString {
			if v == "" {
				continue
			}
			value, _ := strconv.Atoi(string(v))
			boardLine = append(boardLine, cell{isMarked: false, value: value})
		}
		b.Board = append(b.Board, boardLine)
	}
	b.width = len(b.Board[0])
	return b
}

func NewBingoMultiple(input []string) []Bingo {
	var bingos []Bingo

	numberDraws := input[0]
	for index := 1; index <= len(input)-1; index++ {
		bingoInput := []string{numberDraws}
		for subIndex := 0; subIndex < 5; subIndex++ {
			bingoInput = append(bingoInput, input[index+subIndex])
		}
		index += 4
		bingos = append(bingos, NewBingo(bingoInput))
	}

	return bingos
}

func (b *Bingo) Mark(value int) {
	for rowIndex := 0; rowIndex < b.width; rowIndex++ {
		for columnIndex := 0; columnIndex < b.width; columnIndex++ {
			cell := &b.Board[rowIndex][columnIndex]
			if cell.value == value {
				cell.isMarked = true
			}
		}
	}
}

func (b *Bingo) HasBingo() bool {
	return b.HasBingoHorizontally() || b.HasBingoVertically()
}

func (b *Bingo) HasBingoVertically() bool {
	for columnIndex := 0; columnIndex < b.width; columnIndex++ {
		var count int
		for rowIndex := 0; rowIndex < b.width; rowIndex++ {
			cell := b.Board[rowIndex][columnIndex]
			if cell.isMarked {
				count++
			}
		}
		if count == b.width {
			return true
		}
		count = 0
	}
	return false
}

func (b *Bingo) HasBingoHorizontally() bool {
	for _, row := range b.Board {
		var count int
		for _, column := range row {
			if column.isMarked {
				count++
			}
		}
		if count == b.width {
			return true
		}
		count = 0
	}
	return false
}

func (b *Bingo) Score() int {
	sum := 0
	for rowIndex := 0; rowIndex < b.width; rowIndex++ {
		for columnIndex := 0; columnIndex < b.width; columnIndex++ {
			cell := b.Board[rowIndex][columnIndex]
			if !cell.isMarked {
				sum += cell.value
			}
		}
	}
	return sum * b.lastNumber
}

func (b *Bingo) Play() {
	for _, number := range b.NumberDraws {
		b.NumTurns++
		b.Mark(number)
		b.lastNumber = number
		if b.HasBingo() {
			break
		}
	}
}
