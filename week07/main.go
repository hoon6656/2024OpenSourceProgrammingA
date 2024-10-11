package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks"
	replacer := strings.NewReplacer("#", "o")
	fmt.Println(broken)
	fmt.Println(replacer.Replace(broken))

}
