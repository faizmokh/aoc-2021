package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input/input.txt")
	if err != nil {
		log.Fatalf("file error: %v", err)
	}

	depths := getDepthsFrom(file)
	largeMeasurementCount := getMeasurementCount(depths)

	fmt.Println(largeMeasurementCount)
}

func getDepthsFrom(file []byte) []int {
	splits := strings.Split(string(file), "\n")
	var depths []int
	for i := 0; i < len(splits); i++ {
		num, _ := strconv.Atoi(splits[i])
		depths = append(depths, num)
	}
	return depths
}

func getMeasurementCount(depths []int) int {
	count := 0
	for i := 0; i < len(depths)-1; i++ {
		isLarger := isLargerThanPrevious(depths[i], depths[i+1])
		if isLarger {
			count++
		}
	}
	return count
}

func isLargerThanPrevious(prev int, next int) bool {
	if next > prev {
		return true
	}
	return false
}
