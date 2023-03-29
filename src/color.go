package src

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type colorsRGB struct {
	name  string
	red   int
	green int
	blue  int
}

func SetColorANSI(s string) {
	TabColors := []colorsRGB{
		{"green", 0, 255, 0},
		{"yellow", 255, 255, 0},
		{"red", 255, 0, 0},
		{"blue", 0, 0, 255},
		{"white", 255, 255, 255},
		{"black", 0, 0, 0},
		{"purple", 128, 0, 128},
		{"pink", 255, 192, 203},
		{"orange", 255, 165, 0},
		{"brown", 165, 42, 42},
		{"grey", 128, 128, 128},
		{"indigo", 75, 0, 130},
		{"beige", 245, 245, 220},
		{"cyan", 0, 255, 255},
		{"magenta", 255, 0, 255},
		{"turquoise", 64, 224, 208},
	}
	for _, v := range TabColors {
		if v.name == strings.ToLower(s) {
			SetColor(uint8(v.red), uint8(v.green), uint8(v.blue))
		}
	}
}

func SetColor(r, g, b uint8) {
	fmt.Fprintf(os.Stdout, "\033[38;2;%d;%d;%dm", r, g, b)
}

func SetColorRGB(s string) {
	t := strings.Split(s, ", ")
	r, err1 := strconv.ParseInt(t[0], 10, 64)
	g, err2 := strconv.ParseInt(t[1], 10, 64)
	b, err3 := strconv.ParseInt(t[2], 10, 64)
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}
	R := uint8(r)
	G := uint8(g)
	B := uint8(b)
	SetColor(R, G, B)
}

func SetColorHex(s string) {
	if len(s) != 6 {
		return
	}
	red, err1 := strconv.ParseInt(s[0:2], 16, 64)
	green, err2 := strconv.ParseInt(s[2:4], 16, 64)
	blue, err3 := strconv.ParseInt(s[4:], 16, 64)
	if err1 != nil || err2 != nil || err3 != nil {
		return
	}
	r := uint8(red)
	g := uint8(green)
	b := uint8(blue)
	SetColor(r, g, b)
}

func SetHSL(str string) (h, s, l float64) {
	str = strings.ReplaceAll(str, "%", "")
	t := strings.Split(str, ", ")
	if len(t) != 3 {
		return
	}
	H, err1 := strconv.ParseInt(t[0], 10, 64)
	S, err2 := strconv.ParseInt(t[1], 10, 64)
	L, err3 := strconv.ParseInt(t[2], 10, 64)
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}
	return float64(H), float64(S) / 100, float64(L) / 100
}

func HSLtoRGB(h, s, l float64) (r, g, b int) {
	if h > 320 || h < 0 || s > 100 || s < 0 || l > 100 || l < 0 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}
	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2
	var (
		r1, g1, b1 float64
	)
	if h < 60 {
		r1, g1, b1 = c, x, 0
	} else if h < 120 {
		r1, g1, b1 = x, c, 0
	} else if h < 180 {
		r1, g1, b1 = 0, c, x
	} else if h < 240 {
		r1, g1, b1 = 0, x, c
	} else if h < 300 {
		r1, g1, b1 = x, 0, c
	} else {
		r1, g1, b1 = c, 0, x
	}
	r = int((r1 + m) * 255)
	g = int((g1 + m) * 255)
	b = int((b1 + m) * 255)
	return
}

func Inttouint8(r, g, b int) (R, G, B uint8) {
	return uint8(r), uint8(g), uint8(b)
}

func SetColorHSL(s string) {
	SetColor(Inttouint8(HSLtoRGB(SetHSL(s))))
}

/* func main() {
	//setColor(inttouint8(HSLtoRGB(setHSL("120, 50, 50"))))
	//setColorANSI("blue")
	//setColorHex("0E91FF")
	//fmt.Println(StartLineStr("hel"))
	if len(os.Args) == 2 {
		if !ChaineValide(os.Args[1]) {
			fmt.Println("Invalid entry, please retry with a correct text")
		} else {
			s := os.Args[1]
			s = Antislash(s)
			All(s)
		}
	} else {
		fmt.Println("Wrong number of arguments")
	}
	SetColorANSI("white")
} */

func IndexToColor(s1, s2 string) []int {
	t := []int{}
	for _, v1 := range s1 {
		for i, v2 := range s2 {
			if v1 == v2 {
				t = append(t, i)
			}
		}
	}
	return t
}

func IsPresentInArray(num int, arr []int) bool {
	for _, value := range arr {
		if num == value {
			return true
		}
	}
	return false
}

func All(s string) {
	if len(s) == 0 {
		return
	}
	if !ChaineValide(s) {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
		return
	}
	tab := strings.Split(s, "\\n")
	for i, s := range tab {
		if s == "" && i == len(tab)-1 && i > 0 && tab[i-1] == "" {
			return
		}
		if s == "" {
			fmt.Println()
		} else {
			AffMultLine(s)
		}
	}
}

func AllBis(s, col string) {
	if len(s) == 0 {
		return
	}
	if !ChaineValide(s) {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}
	tab := strings.Split(s, "\\n")
	for i, s := range tab {
		if s == "" && i == len(tab)-1 && i > 0 && tab[i-1] == "" {
			return
		}
		if s == "" {
			fmt.Println()
		} else {
			AffMultLineBis(s, os.Args[2])
		}
	}
}

func AffMultLineBis(s, src string) {
	col := os.Args[1][8:]
	tab := IndexToColor(src, s)
	if len(col) == 7 && col[0] == '#' {
		// hexadecimal
		col = strings.ReplaceAll(col, "#", "")
		for i := 1; i <= 8; i++ {
			for j := 0; j < len(s); j++ {
				if IsPrintable(s[j]) {
					if IsPresentInArray(j, tab) {
						SetColorHex(col)
					}
					Affichage(StartLine(rune(s[j])) + i)
					SetColorANSI("white")
				}
			}
			fmt.Println()
		}
	} else if len(col) > 11 && col[0:4] == "hsl(" && col[len(col)-1] == ')' {
		// HSL
		col = strings.ReplaceAll(col, "hsl(", "")
		col = strings.ReplaceAll(col, ")", "")
		//fmt.Println(col)
		for i := 1; i <= 8; i++ {
			for j := 0; j < len(s); j++ {
				if IsPrintable(s[j]) {
					if IsPresentInArray(j, tab) {
						SetColorHSL(col)
					}
					Affichage(StartLine(rune(s[j])) + i)
					SetColorANSI("white")
				}
			}
			fmt.Println()
		}

	} else if len(col) > 11 && col[:4] == "rgb(" && col[len(col)-1] == ')' {
		col = strings.ReplaceAll(col, "rgb(", "")
		col = strings.ReplaceAll(col, ")", "")
		for i := 1; i <= 8; i++ {
			for j := 0; j < len(s); j++ {
				if IsPrintable(s[j]) {
					if IsPresentInArray(j, tab) {
						SetColorRGB(col)
					}
					Affichage(StartLine(rune(s[j])) + i)
					SetColorANSI("white")
				}
			}
			fmt.Println()
		}
	} else {
		//tes
		for i := 1; i <= 8; i++ {
			for j := 0; j < len(s); j++ {
				if IsPrintable(s[j]) {
					if IsPresentInArray(j, tab) {
						SetColorANSI(col)

					}
					Affichage(StartLine(rune(s[j])) + i)
					SetColorANSI("white")
				}
			}
			fmt.Println()
		}
	}

}

func AffMultLine(s string) {
	for i := 1; i <= 8; i++ {
		for j := 0; j < len(s); j++ {
			if IsPrintable(s[j]) {
				Affichage(StartLine(rune(s[j])) + i)
			}
		}
		fmt.Println()
	}
}

func ChaineValide(s string) bool {

	tab := []rune(s)
	return len(tab) == len(s)
}

func StartLine(r rune) int {
	return (int(r-32) * 9) + 1
}

func IsPrintable(r byte) bool {
	if r < 32 || r > 126 {
		return false
	}
	return true
}

func Affichage(m int) {
	file, err := os.Open("template/standard.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for i := 1; i <= m-1; i++ {
		scanner.Scan()
	}
	if scanner.Scan() {
		fmt.Print(scanner.Text())
	}
}

func Antislash(s string) string {
	s = strings.ReplaceAll(s, "\n", "\\n")
	return s
}
