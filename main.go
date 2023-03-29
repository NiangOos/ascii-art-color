package main

import (
	src "color/src"
	"fmt"
	"os"
	"strings"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			fmt.Println()
			fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		}
	}()

	//setColor(inttouint8(HSLtoRGB(setHSL("120, 50, 50"))))
	//setColorANSI("blue")
	//setColorHex("0E91FF")
	//fmt.Println(StartLineStr("hel"))
	//setColor(51, 255, 212)
	if len(os.Args) == 2 {
		if !src.ChaineValide(os.Args[1]) {

			fmt.Println("Usage: go run . [OPTION] [STRING]")
			fmt.Println()
			fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
			os.Exit(0)
		} else {
			s := os.Args[1]
			s = src.Antislash(s)
			src.All(s)

		}
	} else if len(os.Args) == 3 {
		col := os.Args[1]
		if len(col) > 8 && col[:8] == "--color=" {
			col = col[8:]
			s := os.Args[2]
			s = src.Antislash(s)
			if len(col) == 7 && col[0] == '#' {
				// hexadecimal
				col = strings.ReplaceAll(col, "#", "")
				src.SetColorHex(col)
				src.All(s)
				src.SetColorANSI("white")
			} else if len(col) > 11 && col[0:4] == "hsl(" && col[len(col)-1] == ')' {
				// HSL
				col = strings.ReplaceAll(col, "hsl(", "")
				col = strings.ReplaceAll(col, ")", "")
				//fmt.Println(col)
				src.SetColorHSL(col)
				src.All(s)
				src.SetColorANSI("white")

			} else if len(col) > 11 && col[:4] == "rgb(" && col[len(col)-1] == ')' {
				col = strings.ReplaceAll(col, "rgb(", "")
				col = strings.ReplaceAll(col, ")", "")
				src.SetColorRGB(col)
				src.All(s)
				src.SetColorANSI("white")
			} else {
				//test
				src.SetColorANSI(col)
				src.All(s)
				src.SetColorANSI("white")
			}

		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			fmt.Println()
			fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
			os.Exit(0)
		}
	} else if len(os.Args) == 4 {
		col := os.Args[1]
		if len(col) > 8 && col[:8] == "--color=" {
			col = col[8:]
			s := os.Args[3]
			s = src.Antislash(s)
			src.AllBis(s, col)

		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			fmt.Println()
			fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
			os.Exit(0)
		}
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}
	//setColorANSI("white")
}
