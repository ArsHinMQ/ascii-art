package ascii_art

import (
	"fmt"
	"os"
)

type AsciiArt [][]string

func (asciiArt *AsciiArt) Print2Terminal() {
	for _, row := range *asciiArt {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func (asciiArt *AsciiArt) Convert2Text() string {
	res := ""
	for _, row := range *asciiArt {
		for _, char := range row {
			res += char
		}
		res += "\n"
	}
	return res
}

func (asciiArt *AsciiArt) Write2File(fname string) error {
	var f *os.File
	var err error

	if _, err = os.Stat(fname); err == nil {
		f, err = os.Open(fname)
	} else {
		f, err = os.Create(fname)
	}
	if err != nil {
		return err
	}
	defer f.Close()

	art := asciiArt.Convert2Text()
	_, err = f.WriteString(art)

	return err
}
