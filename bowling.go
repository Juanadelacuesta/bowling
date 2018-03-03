package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	split  = 0
	strike = 1
)

func main() {
	in := "| 1 4 | 4 5 | 6 / | 5 / | X | 0 1 | 7 / | 6 / | X | 2 / 6 |"
	in = strings.Replace(in, " ", "", -1)
	in = in[1:len(in)-1] + "|00"

	score := 0
	frames := strings.Split(in, "|")

	for i := 0; i < 10; i++ {
		score += getFrameScore(frames[i])

		if isStrike(frames[i]) {
			score += getStrikeBonus(frames[i+1])
		}

		if isSplit(frames[i]) {
			score += getSplitBonus(frames[i+1])
		}
	}
	fmt.Println(score)
}

func isStrike(frame string) bool {
	return strings.Contains(frame, "X")
}
func isSplit(frame string) bool {
	return strings.Contains(frame, "/")
}

func getFrameScore(f string) (score int) {
	frame := strings.Split(f, "")
	for i, v := range frame {
		switch v {
		case "/":
			score += (10 - aToI(frame[i-1]))
		case "X":
			score += 10
		default:
			score += aToI(v)
		}
	}
	return
}

func getStrikeBonus(f string) (bonus int) {
	return getFrameScore(f[:2])
}

func getSplitBonus(f string) (bonus int) {
	frame := strings.Split(f, "")
	if frame[0] == "X" {
		bonus = 10
	} else {
		bonus = aToI(frame[0])
	}
	return
}

func aToI(val string) int {
	i, _ := strconv.Atoi(val)
	return i
}
