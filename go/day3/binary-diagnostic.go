package main

import (
	"aoc2021/utilities"
	"fmt"
	"strconv"
	"strings"
)


func main() {
    file, input := utilities.ScanStringsFromFile("./day3/input.txt")
    defer file.Close()

    orderedBits := make([][]string, len(input[0]))
    
    for _, row := range input {
        for i, bitStr := range row {
            bit := string(bitStr)
            orderedBits[i] = append(orderedBits[i], bit)
        }
    }

    var gammaRateBitStr []string
    var epsilonRateBitStr []string
    for _, bits := range orderedBits {
        gammaRateBitStr = append(gammaRateBitStr, findCommonBit(bits, false))
        epsilonRateBitStr = append(epsilonRateBitStr, findCommonBit(bits, true))
    }

    gammaRate, _ := strconv.ParseInt(strings.Join(gammaRateBitStr, ""), 2, 64)
    epsilonRate, _ := strconv.ParseInt(strings.Join(epsilonRateBitStr, ""), 2, 64)

    fmt.Printf("gamma rate: %d", gammaRate)
    fmt.Printf("epsilon rate: %d", epsilonRate)
    fmt.Printf("power consumption: %d", gammaRate * epsilonRate)
}

func findCommonBit(bits []string, invert bool) string {
    bitCount := 0

    for _, b := range bits {
        if b == "1" {
            bitCount++
        }    
    }

    if (bitCount > len(bits) / 2) && invert == false || (bitCount < len(bits) / 2) && invert == true {
        return "1"
    }
    
    return "0"
}

