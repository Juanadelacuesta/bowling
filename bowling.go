package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	split  = 0
	strike = 1
)

func main() {
	file, err := os.Open(os.Args[1])
	defer file.Close()
	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		fmt.Println(getScore(scanner.Text()))
	}
	os.Exit(0)
}

func getScore(input string) (score int) {
	frames := strings.Split(cleanInput(input), "|")
	for i := 0; i <= 9; i++ {
		score += getFrameScore(frames[i])

		if isStrike(frames[i]) {
			if i < 8 {
				score += getStrikeBonus(frames[i+1], frames[i+2])
			}
			if i == 8 {
				score += getStrikeBonus(frames[i+1], "")
			}
		}

		if isSplit(frames[i]) && i < 9 {
			score += getSplitBonus(frames[i+1])
		}
	}
	return
}

func cleanInput(input string) string {
	input = strings.Replace(input, " ", "", -1)
	input = input[1 : len(input)-1]
	return input
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

func getStrikeBonus(f1, f2 string) (bonus int) {
	return getFrameScore((f1 + f2)[:2])
}

func getSplitBonus(f string) (bonus int) {
	return getFrameScore(f[:1])
}

func aToI(val string) int {
	i, _ := strconv.Atoi(val)
	return i
}
