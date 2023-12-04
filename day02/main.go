package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var input io.Reader

	if len(os.Args) > 1 {
		var err error
		file := os.Args[1]
		input, err = os.Open(file)
		if err != nil {
			return err
		}
	} else {
		input = os.Stdin
	}

	s := bufio.NewScanner(input)
	games := []game{}
	for s.Scan() {
		g, err := parseGame(s.Text())
		if err != nil {
			return err
		}
		games = append(games, *g)
	}
	if s.Err() != nil {
		return s.Err()
	}

	restrictions := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	total := 0

OUTER:
	for _, g := range games {
		for _, round := range g.rounds {
			for color, count := range round {
				maxCount, ok := restrictions[color]
				if !ok {
					continue OUTER
				}
				if count > maxCount {
					continue OUTER
				}
			}
		}

		total += g.id
	}
	fmt.Println(total)

	total2 := 0
	for _, g := range games {
		maxValues := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}
		for _, round := range g.rounds {
			for color, count := range round {
				maxValues[color] = max(maxValues[color], count)
			}
		}
		total2 += mul(maxValues)
	}
	fmt.Println(total2)
	return nil
}

func mul(data map[string]int) int {
	i := 1
	for _, n := range data {
		i *= n
	}
	return i
}

type game struct {
	id     int
	rounds []map[string]int
}

func parseGame(input string) (*game, error) {
	gameIDPrefix := "Game "

	gameIDRaw, gameRaw, ok := strings.Cut(input, ":")
	if !ok {
		return nil, fmt.Errorf("invalid input '%s'", input)
	}
	if len(gameIDRaw) < len(gameIDPrefix) {
		return nil, fmt.Errorf("invalid input '%s'", input)
	}
	gameID, err := strconv.Atoi(gameIDRaw[len(gameIDPrefix):])
	if err != nil {
		return nil, err
	}

	rounds := []map[string]int{}
	for _, colorSetsRaw := range strings.Split(gameRaw, ";") {
		colorSetsRaw = strings.TrimSpace(colorSetsRaw)
		round := map[string]int{}
		for _, colorSetRaw := range strings.Split(colorSetsRaw, ", ") {
			colorSetRaw = strings.TrimSpace(colorSetRaw)
			countRaw, color, ok := strings.Cut(colorSetRaw, " ")
			if !ok {
				return nil, fmt.Errorf("invalid input '%s'", input)
			}
			count, err := strconv.Atoi(countRaw)
			if err != nil {
				return nil, fmt.Errorf("invalid input '%s': %w", input, err)
			}
			round[color] += count
		}
		rounds = append(rounds, round)

	}

	return &game{
		id:     gameID,
		rounds: rounds,
	}, nil
}
