package submarine

import (
	"testing"

	"github.com/WTFox/aoc2021/util"
)

func TestIntegration(t *testing.T) {
	testInputs := util.ReadStringsFromFile("./inputs/day04.txt")
	bingoBoards := NewBingoMultiple(testInputs)

	if len(bingoBoards) != 100 {
		t.Errorf("Expected 100 boards, got %d", len(bingoBoards))
	}
}

func TestMultiBingo(t *testing.T) {
	testInput := []string{
		"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
		"",
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
		"",
		" 3 15  0  2 22",
		" 9 18 13 17  5",
		"19  8  7 25 23",
		"20 11 10 24  4",
		"14 21 16 12  6",
		"",
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		" 2  0 12  3  7",
	}

	bingos := NewBingoMultiple(testInput)
	if len(bingos) != 3 {
		t.Errorf("Expected %d got %d", 3, len(bingos))
	}
}
func TestBingo(t *testing.T) {
	testInput := []string{
		"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
		"",
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
	}

	t.Run("Test Mark", func(t *testing.T) {
		b := NewBingo(testInput)
		b.Mark(22)
		if !b.Board[0][0].isMarked {
			t.Error("No mark")
		}
	})

	t.Run("Test HasBingoVertically", func(t *testing.T) {
		b := NewBingo(testInput)

		b.Mark(22)
		b.Mark(8)
		b.Mark(21)
		b.Mark(6)
		b.Mark(1)

		if !b.HasBingoVertically() {
			t.Error("No bingo")
		}
	})

	t.Run("Test HasBingoHorizontally", func(t *testing.T) {
		b := NewBingo(testInput)

		b.Mark(22)
		b.Mark(13)
		b.Mark(17)
		b.Mark(11)
		b.Mark(0)

		if !b.HasBingoHorizontally() {
			t.Error("No bingo")
		}
	})

	t.Run("Test Score - wins horizontally", func(t *testing.T) {
		testInput := []string{
			"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
			"",
			"14 21 17 24  4",
			"10 16 15  9 19",
			"18  8 23 26 20",
			"22 11 13  6  5",
			" 2  0 12  3  7",
		}

		b := NewBingo(testInput)
		b.Play()

		expectedScore := 4512
		got := b.Score()
		if b.Score() != expectedScore {
			t.Errorf("Expected %d got %d", expectedScore, got)
		}
	})
	t.Run("Test Score - wins vertically", func(t *testing.T) {
		testInput := []string{
			"14,10,18,22,2",
			"",
			"14 21 17 24  4",
			"10 16 15  9 19",
			"18  8 23 26 20",
			"22 11 13  6  5",
			" 2  0 12  3  7",
		}

		b := NewBingo(testInput)
		b.Play()

		expectedScore := 518
		got := b.Score()
		if b.Score() != expectedScore {
			t.Errorf("Expected %d got %d", expectedScore, got)
		}
	})
}
