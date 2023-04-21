package function

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFont(banner string) (map[rune][]string, error) {
	input, err := os.ReadFile(banner)
	if err != nil {
		log.Fatal(err)
	}

	standardList := strings.Split(string(input), "\n")
	s := make(map[rune][]string)
	var startPoint rune = ' '
	var temp []string
	for i, v := range standardList {
		i++
		if i%9 == 0 {
			temp = append(temp, v)
			s[startPoint] = temp
			startPoint++
			temp = nil
		} else if v != "" {
			temp = append(temp, v)
		}
	}

	return s, nil
}

func PrintFormat(text string, format map[rune][]string) {
	if len(text) < 1 {
		return
	}
	if text == "\\n" {
		fmt.Println("")
		return
	}

	line := strings.Split(text, "\\n")
	for _, word := range line {
		if len(word) < 1 {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {

			for _, char := range word {
				fmt.Print(format[char][i])
			}
			fmt.Println()
		}
	}
}
