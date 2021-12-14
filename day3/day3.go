package day3

import (
	"math"
	"strconv"
	"strings"
)

func CalculatePowerConsumption(report []string) int {
	if len(report) == 0 {
		return 0
	}

	lenOfCode := len(report[0])
	highBitCounter := make([]int, lenOfCode)
	for _, code := range report {
		if lenOfCode != len(code) {
			return -1
		}

		bits := strings.Split(code, "")
		for bitIndex, bit := range bits {
			if bit == "1" {
				highBitCounter[bitIndex]++
			}
		}
	}

	gammaRate := getGammaRateFromCounter(len(report), highBitCounter)
	epsilonRate := getEpsilonRateFromCounter(len(report), highBitCounter)

	return getDecimalFromBinary(gammaRate) * getDecimalFromBinary(epsilonRate)
}

func getGammaRateFromCounter(reportLength int, highBitCounter []int) (gammaRate string){
	for _, highBits := range highBitCounter {
		if highBits > reportLength-highBits {
			gammaRate += "1"
		} else {
			gammaRate += "0"
		}
	}
	return
}

func getEpsilonRateFromCounter(reportLength int, highBitCounter []int) (epsilonRate string){
	for _, highBits := range highBitCounter {
		if highBits > reportLength-highBits {
			epsilonRate += "0"
		} else {
			epsilonRate += "1"
		}
	}
	return
}

func getDecimalFromBinary(bits string) (decimal int) {
	bitLength := len(bits)

	for i, bitStr := range strings.Split(bits, "") {
		bit, _ := strconv.Atoi(bitStr)
		binaryValue := math.Pow(2, float64(bitLength-i-1))
		decimal = decimal + (bit * int(binaryValue))
	}

	return
}