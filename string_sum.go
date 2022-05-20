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
	if len(input) == 0 {
		return "", fmt.Errorf("error while calculating sum: %w", errorEmptyInput)
	}
	sum := 0
	r := regexp.MustCompile(`[-+](\w+)|(\w+)`)
	s := r.FindAllString(input, -1)
	if len(s) > 2 || len(s) <= 1 {
		return "", fmt.Errorf("error while calculating sum: %w", errorNotTwoOperands)
	}
	for i := 0; i < len(s); i++ {
		num, err := strconv.Atoi(s[i])
		if err != nil {
			e := err.(*strconv.NumError)
			return "", fmt.Errorf("error while calculating sum: %w", e)
		}
		sum += num
	}
	output = strconv.Itoa(sum)
	return output, nil
}
