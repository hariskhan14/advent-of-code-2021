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

func getGammaRateFromCounter(reportLength int, highBitCounter []int) (gammaRate string) {
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

func CalculateOxygenAndC02GeneratorRating(report []string) int {
	oxygen := CalculateOxygenGeneratorRating(report)
	c02 := CalculateC02GeneratorRating(report)
	return oxygen * c02
}

func CalculateOxygenGeneratorRating(report []string) int {
	return CalculateGeneratorRating(report, "1", "0")
}

func CalculateC02GeneratorRating(report []string) int {
	return CalculateGeneratorRating(report, "0", "1")
}

func CalculateGeneratorRating(report []string, useWhenGreater, useWhenSmaller string) int {
	if len(report) == 0 {
		return -1
	}

	lenOfCode := len(report[0])
	for i := 0 ; i < lenOfCode; i++ {
		highBits := 0

		if len(report) == 1 {
			break
		}

		for _, code := range report {
			bits := strings.Split(code, "")
			if bits[i] == "1" {
				highBits++
			}
		}

		var filterOn string
		if highBits >= len(report) - highBits {
			filterOn = useWhenGreater
		} else {
			filterOn = useWhenSmaller
		}

		report = filterReportOnBitAndIndex(report, filterOn, i)
	}

	return getDecimalFromBinary(report[0])
}

func filterReportOnBitAndIndex(report []string, filterOn string, idx int) []string {
	var result []string
	for _, code := range report {
		bits := strings.Split(code, "")
		if bits[idx] == filterOn {
			result = append(result, code)
		}
	}

	return result
}