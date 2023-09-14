
# ASCII ART CREATOR

Create ASCII art using GoLang

## Currently Supports
- Image to ASCII art

## How to Use
- Image to ASCII art
```
// using default values
art := asciiArt.Image2Art(img, []rune{}, 0)
```
```
// using custom characters
art := asciiArt.Image2Art(img, []rune{' ', '.'}, 0)
```
```
// using custom accuracy - the lower the number, the more accurate(default=4)
art := asciiArt.Image2Art(img, []rune{' ', '.'}, 1)
```
Full example
```
import (
  asciiArt "ascii-art"
  "fmt"
)

func main() {
  // You can use an existing image of type image.Image or open your image using this method
  img, err := asciiArt.OpenImage("path/to/img.jpg")
  if err != nil {
    fmt.Println(err)
    return
  }
  // convert image to ASCIIArt - using default values
  art := asciiArt.Image2Art(img, []rune{}, 0)
}
```
- output
    - Print to terminal
    ```
    art.Print2Terminal()
    ```
    - Convert to text
    ```
    artText := art.Convert2Text()
	fmt.Println(artText)
    ```
    - Convert to file
    ```
    err = art.Write2File("out.txt")
    if err != nil {
      fmt.Println(err)
    }
    ```

