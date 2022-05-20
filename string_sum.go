package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")
	r := []rune(input)
	sum := 0
	operandsCount := 0
	numbersRegexp := regexp.MustCompile(`\d`)
	if len(r) == 0 {
		return "", fmt.Errorf("Error while calculating sum: %w", errorEmptyInput)
	}
	for i := 0; i < len(r); i++ {
		num, _ := strconv.Atoi(string(r[i]))
		if operandsCount > 2 {
			return "", fmt.Errorf("Error while calculating sum: %w", errorNotTwoOperands)
		} else if string(r[i]) == "-" {
			i++
			num, _ := strconv.Atoi(string(r[i]))
			sum -= num
			operandsCount += 1
		} else if string(r[i]) == "+" {
			i++
			num, _ := strconv.Atoi(string(r[i]))
			sum += num
			operandsCount += 1
		} else if numbersRegexp.MatchString(string(r[i])) {
			sum += num
			operandsCount += 1
		}
	}
	output = strconv.Itoa(sum)
	return output, nil
}
