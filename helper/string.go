package helper

import (
	"fmt"
	"strings"
)

func CutString(text string, length int) string {
	str := strings.Split(text, "");

	var result string
	for index := 0; index < len(str); index++ {
		if index > length {
			str[index] = ""
			// str = ;
			result = strings.Join(str, "") + "..."
		} 
	}
	fmt.Println(result)
	if result != "" {
		return result
	}
	return text
}