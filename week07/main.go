package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	broken := "G# r#cks"
	replacer := strings.NewReplacer("#", "o")
	fmt.Println(broken)
	fmt.Println(replacer.Replace(broken))

	var now time.Time = time.Now()
	var year int = int(now.Year())
	var month int = int(now.Month())
	var day int = int(now.Day())

	fmt.Printf("오늘은 %d년 %d월 %d일\n", year, month, day)
	fmt.Printf("지금은 %d시 %d분 %d초", now.Hour(), now.Minute(), now.Second())

}
