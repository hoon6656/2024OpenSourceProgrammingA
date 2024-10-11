package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("점수 입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	score, err := strconv.parseInt(i, 16, 32)
	if score >= 60 {
		fmt.Println("A")
		fmt.Printf("%d\n", score)
	} else {
		fmt.Println("BCDEF")
		fmt.Printf("%d\n", score)
	}
}
