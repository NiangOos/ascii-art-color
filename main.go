package main

import (
	src "color/src"
	"os"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			src.ErrorMessageColor()
		}
	}()

	if len(os.Args) == 2 {
		if !src.ChaineValide(os.Args[1]) {
			src.ErrorMessageColor()
		} else {
			src.BasicAffich()
		}
	} else if len(os.Args) == 3 {
		opt := os.Args[1]
		if len(opt) > 8 && opt[:8] == "--color=" {
			src.ColArgs3(opt)
		} else {
			src.ErrorMessageColor()
		}
	} else if len(os.Args) == 4 {
		opt := os.Args[1]
		if len(opt) > 8 && opt[:8] == "--color=" {
			src.ColArgs4(opt)
		} else {
			src.ErrorMessageColor()
		}
	} else {
		src.ErrorMessageColor()
	}
}

//setColor(inttouint8(HSLtoRGB(setHSL("120, 50, 50"))))
//setColorANSI("blue")
//setColorHex("0E91FF")
//fmt.Println(StartLineStr("hel"))
//setColor(51, 255, 212)
