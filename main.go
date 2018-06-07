package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

var (
	yellow = color.New(color.BgHiYellow).Add(color.FgBlack)
)

func main() {
	keywords := os.Args[1:]
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l := scanner.Text()
		for _, keyword := range keywords {
			l = strings.Replace(l, keyword, yellow.Sprint(keyword), -1)
		}
		fmt.Println(l)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
