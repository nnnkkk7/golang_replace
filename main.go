package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "you me he you me he we she you you me she we"
	fmt.Println(str)

	replace1 := strings.Replace(str, "you", "re", 1)
	fmt.Println(replace1)

	replace2 := strings.Replace(str, "me", "re", -1)
	fmt.Println(replace2)

	replace3 := strings.Replace(str, "he", "re", -1)
	fmt.Println(replace3)
}
