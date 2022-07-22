package controFunc

import (
	"fmt"
	"regexp"
)

func LetOnlyNumbers(str string) string {

	regx := regexp.MustCompile(`[^ 0-9]`)

	return fmt.Sprint(regx.ReplaceAllString(str, ""))

}
