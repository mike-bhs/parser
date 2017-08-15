package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

	s "github.com/parser/services"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Website Url: ")

	input, _ := reader.ReadString('\n')
	_, err := url.ParseRequestURI(input)

	if err != nil {
		fmt.Println("Input isn't valid!")
		return
	}

	// s.GetParsersList(input)
	s.ParseSite(input)
}
