package ascii_art

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

type Accuracy int8
type Tile [][]image.Gray

type Scale struct {
	x int
	y int
}

func Image2Art(img image.Image, asciiChar []rune, accuracy Accuracy) *AsciiArt {
	if accuracy == 0 {
		accuracy = 4
	}

	if len(asciiChar) == 0 {
		asciiChar = []rune{
			'$',
			'@',
			'B',
			'%',
			'8',
			'&',
			'W',
			'M',
			'#',
			'*',
			'o',
			'a',
			'h',
			'k',
			'b',
			'd',
			'p',
			'q',
			'w',
			'm',
			'Z',
			'O',
			'0',
			'Q',
			'L',
			'C',
			'J',
			'U',
			'Y',
			'X',
			'z',
			'c',
			'v',
			'u',
			'n',
			'x',
			'r',
			'j',
			'f',
			't',
			'/',
			'\\',
			'|',
			'(',
			')',
			'1',
			'{',
			'}',
			'[',
			']',
			'?',
			'-',
			'_',
			'+',
			'~',
			'<',
			'>',
			'i',
			'!',
			'l',
			'I',
			';',
			':',
			',',
			'"',
			'^',
			'`',
			'\'',
			'.',
			' ',
		}
	}

	imgSet := imgRGBA2GrayScale(&img)

	scale, err := defineResultAccuracy(&img, accuracy)
	if err != nil {
		log.Fatal(err)
	}

	tiles := splitImg2Tiles(imgSet, scale)

	return convertTiles2Art(tiles, asciiChar, scale)
}

func OpenImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func imgRGBA2GrayScale(img *image.Image) *image.RGBA {
	b := (*img).Bounds()
	imgPtr := image.NewRGBA(b)

	for y := 0; y <= b.Max.Y; y++ {
		for x := 0; x <= b.Max.X; x++ {
			oldPixel := (*img).At(x, y)
			r, g, b, _ := oldPixel.RGBA()
			yx := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			pixel := color.Gray{uint8(yx / 256)}
			imgPtr.Set(x, y, pixel)
		}
	}

	return imgPtr
}

func defineResultAccuracy(img *image.Image, accuracy Accuracy) (Scale, error) {

	if accuracy <= 0 {
		return Scale{0, 0}, fmt.Errorf("accurancy must be higher than 0")
	}

	b := (*img).Bounds()
	x, y := b.Max.X/int(accuracy), b.Max.Y/int(accuracy)

	return Scale{x, y}, nil
}

func splitImg2Tiles(img *image.RGBA, scale Scale) *Tile {
	width := (*img).Bounds().Dx() / scale.x
	height := (*img).Bounds().Dy() / scale.y

	tiles := make(Tile, scale.y+1)
	for i := 0; i <= scale.y; i++ {
		row := make([]image.Gray, scale.x+1)
		for j := 0; j <= scale.x; j++ {
			tile := image.NewGray(image.Rect(0, 0, width, height))
			draw.Draw(tile, tile.Bounds(), img, image.Point{j * width, i * height}, draw.Src)

			row[j] = *tile
		}
		tiles[i] = row
	}

	return &tiles
}

func convertTiles2Art(tiles *Tile, asciiChar []rune, scale Scale) *AsciiArt {
	characters := make(AsciiArt, scale.y+1)

	for i, row := range *tiles {
		cRow := make([]string, scale.x+1)
		for j, t := range row {
			p := t.At(t.Bounds().Min.X, t.Bounds().Min.X)
			r, b, g, _ := p.RGBA()
			brightness := int((r+b+g)/3) / 256
			if brightness > 255 {
				brightness = 255
			}
			code := brightness * len(asciiChar) / 256
			if code == len(asciiChar) {
				code = code - 1
			}
			cRow[j] = string(asciiChar[code])
		}
		characters[i] = cRow
	}

	return &characters
}
