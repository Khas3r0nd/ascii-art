package function

import (
	"log"
	"os"
)

var (
	err1 string = "Expected 1 argument, got "
	err2 string = "PLZ use Standard character set (32 - 126)"
	err3 string = "Banner should not be changed!"
)

func Start(args []string) {
	if len(args) != 2 {
		log.Fatal(err1, len(args[1:]))
	}

	arg := args[1]

	// check cyrillic characters

	for _, i := range arg {
		if i > 126 || i < 32 {
			log.Fatal(err2)
		}
	}
	fileName := "banners/standard.txt"
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	hash := MD5(string(content))
	if !CheckHash(hash) {
		log.Fatal(err3)
	}

	standFont, err := ReadFont(fileName)
	if err != nil {
		log.Fatal(err)
	}
	PrintFormat(arg, standFont)
}
