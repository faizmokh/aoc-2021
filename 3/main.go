package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// read over binaries
// store binaries in multi dimensional
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

	var binaries [][]string
	for scanner.Scan() {
		text := scanner.Text()
		binary := strings.Split(text, "")
		binaries = append(binaries, binary)
	}

	length := len(binaries[0])
	var gammaRateArray []string
	var epsilonRateArray []string

	for i := 0; i < length; i++ {
		occurence := make(map[string]int)
		for _, binary := range binaries {
			occurence[binary[i]] += 1
		}
		gammaRateArray = append(gammaRateArray, getMostCommon(occurence))
		epsilonRateArray = append(epsilonRateArray, getLeastCommon(occurence))
	}

	gammaRateBinary := strings.Join(gammaRateArray, "")
	gammaRate, _ := strconv.ParseInt(gammaRateBinary, 2, 64)

	epsilonRateBinary := strings.Join(epsilonRateArray, "")
	epsilonRate, _ := strconv.ParseInt(epsilonRateBinary, 2, 64)

	powerConsumption := gammaRate * epsilonRate
	fmt.Println(powerConsumption)
}

func getLeastCommon(occurence map[string]int) string {
	if occurence["1"] > occurence["0"] {
		return "0"
	} else {
		return "1"
	}
}

func getMostCommon(occurence map[string]int) string {
	if occurence["1"] > occurence["0"] {
		return "1"
	} else {
		return "0"
	}
}
