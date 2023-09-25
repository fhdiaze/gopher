package twentytwo

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func Solve() {
	scanner := scanInput()
	var stacks [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		token := parseLine(line)
		stacks = run(stacks, token)
	}

	printTop(stacks)
}

type Token interface {
}

type Command struct {
	quantity int
	from     int
	to       int
}

func printTop(stacks [][]byte) {
	for _, stack := range stacks {
		if len(stack) > 0 {
			fmt.Printf("%c", stack[len(stack)-1])
		} else {
			fmt.Printf(" ")
		}
	}

	fmt.Println()
}

func run(stacks [][]byte, token Token) [][]byte {
	switch t := token.(type) {
	case Command:
		runMove(stacks, t)
	case string:
		stacks = runLoad(stacks, t)
	}

	return stacks
}

func runMove(stacks [][]byte, t Command) {
	fmt.Println(stacks)
	var marks []byte
	marks = append(marks, stacks[t.from-1][len(stacks[t.from-1])-t.quantity:]...)
	slices.Reverse(marks)
	stacks[t.to-1] = append(stacks[t.to-1], marks...)
	stacks[t.from-1] = stacks[t.from-1][:len(stacks[t.from-1])-t.quantity]
	fmt.Println(stacks)
}

func runLoad(stacks [][]byte, t string) [][]byte {
	if len(t) == 0 {
		return stacks
	}

	for i, c := range t {
		stacks = tryAddStack(stacks, i)
		// Check that c is a mark
		if unicode.IsLetter(c) {
			stacks[i] = append([]byte{byte(c)}, stacks[i]...)
		}
	}

	return stacks
}

func tryAddStack(stacks [][]byte, s int) [][]byte {
	for s >= len(stacks) {
		stacks = append(stacks, make([]byte, 0))
	}

	return stacks
}

func parseLine(line string) Token {
	switch {
	case len(line) == 0:
		return nil
	case line[0] == 'm':
		return parseMove(line)
	case line[0] == ' ', line[0] == '[':
		return parseMarks(line)
	default:
		return nil
	}
}

func parseMove(line string) Command {
	regex := regexp.MustCompile(`move (?P<q>\d*) from (?P<f>\d*) to (?P<t>\d*)`)
	matches := regex.FindStringSubmatch(line)
	quantity, _ := strconv.Atoi(matches[1])
	from, _ := strconv.Atoi(matches[2])
	to, _ := strconv.Atoi(matches[3])

	return Command{quantity: quantity, from: from, to: to}
}

func parseMarks(line string) string {
	ti := 1
	var tokens strings.Builder

	for i, c := range line {
		if i == ti {
			tokens.WriteByte(byte(c))

			ti += 4
		}
	}

	return tokens.String()
}

func scanInput() *bufio.Scanner {
	file, err := os.Open("internal/aoc/twentytwo/five.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	return scanner
}
