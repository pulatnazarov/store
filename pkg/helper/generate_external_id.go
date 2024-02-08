package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func GenerateExternalID(extID string) string { // (I-0099) I-0100
	// logic
	parts := strings.Split(extID, "-")
	var newExtID string

	if len(parts) == 2 {
		prefix := parts[0]    // I
		numberStr := parts[1] // 0001

		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Println("error is while converting integer part", err.Error())
			return ""
		}
		number++

		newExtID = fmt.Sprintf("%s-%04d", prefix, number)
	}

	return newExtID
}
