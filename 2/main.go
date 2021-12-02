package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input/input.txt")
	if err != nil {
		log.Fatalf("file error: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	commands := getCommandsFrom(scanner)
	horizontalPosition, depth := getTotalHorizontalPositionAndAim(commands)
	multipled := horizontalPosition * depth
	fmt.Println(multipled)
}

type Command struct {
	direction string
	unit      int
}

func getCommandsFrom(scanner *bufio.Scanner) []Command {
	var commands []Command
	for scanner.Scan() {
		s := scanner.Text()
		splitted := strings.Split(s, " ")
		unit, _ := strconv.Atoi(splitted[1])
		command := Command{
			direction: strings.ToLower(splitted[0]),
			unit:      unit,
		}
		commands = append(commands, command)
	}
	return commands
}

func getTotalHorizontalPositionAndAim(commands []Command) (int, int) {
	position := 0
	depth := 0
	aim := 0
	for _, command := range commands {
		switch command.direction {
		case "forward":
			position += command.unit
			depth += aim * command.unit
		case "down":
			aim += command.unit
		case "up":
			aim -= command.unit
		}
	}
	return position, depth
}
