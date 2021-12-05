package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = `
199
200
208
210
200
`

func main() {
	count := 0
	prev := 0
	input := strings.Split(strings.TrimSpace(input), "\n")

	for i, line := range input {
		current, _ := strconv.Atoi(line)

		if i == 0 {
			prev = current
			continue
		}

		if current > prev {
			count++
		}
		prev = current
	}

	fmt.Println(count)
}
