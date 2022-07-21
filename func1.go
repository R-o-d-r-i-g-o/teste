package controFunc

import (
	"fmt"
	"regexp"
)

func RevomeSpecialChars(str string) string {

	regx := regexp.MustCompile(`[^ A-Za-z0-9]`)

	return fmt.Sprint(regx.ReplaceAllString(str, ""))
}

func CleanTerminal() {

	fmt.Print("\033[H\033[2J")
}
