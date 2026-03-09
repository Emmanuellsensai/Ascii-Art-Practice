package main

import (
	"fmt"
	"os"

	"ascii-art/ascii"
)

func main() {

	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]

	bannerLines, err := ascii.ReadBanner("thinkertoy.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	asciiMap := ascii.BuildAsciiMap(bannerLines)

	ascii.PrintAscii(input, asciiMap)
}