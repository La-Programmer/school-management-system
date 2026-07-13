package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ExtractId(requestUrl string, entity string) (int, error) {
	extractedId := strings.TrimLeft(requestUrl, fmt.Sprintf("/%s/", entity))

	if extractedId == "" {
		return 0, nil
	}

	convertedId, err := strconv.Atoi(extractedId)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return convertedId, nil
}
