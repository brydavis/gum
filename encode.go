package main

import (
	// "flag"
	"bytes"
	"fmt"
	"image"
	// "image/color"
	// "image/draw"
	"image/gif"
	"image/jpeg"

	// "io"
	"io/ioutil"
	// "math/rand"
	"os"
	"path"
	// "time"
)

func main() {
	// flag.Parse()
	// rand.Seed(time.Now().UTC().UnixNano())
	out, err := os.Create("./output.gif")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} // generate some QR code look a like image

	// imgRect := image.Rect(0, 0, 100, 100)
	// img := image.NewGray(imgRect)
	// draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
	// for y := 0; y < 100; y += 10 {
	// 	for x := 0; x < 100; x += 10 {
	// 		fill := &image.Uniform{color.Black}
	// 		if rand.Intn(10)%2 == 0 {
	// 			fill = &image.Uniform{color.White}
	// 		}
	// 		draw.Draw(img, image.Rect(x, y, x+10, y+10), fill, image.ZP, draw.Src)
	// 	}
	// }

	dir, _ := ioutil.ReadDir(".")

	var imgs []*image.Paletted

	for _, file := range dir {
		switch ext := path.Ext(file.Name()); ext {
		case ".jpg":

			f, _ := os.Open(file.Name())
			img, _ := jpeg.Decode(f)
			var opt gif.Options
			opt.NumColors = 256

			buf := new(bytes.Buffer)
			gif.Encode(buf, img, &opt)

			g, _ := gif.DecodeAll(buf)

			imgs = append(imgs, g.Image[0])

		default:
			fmt.Println(ext)
		}
	}

	g := new(gif.GIF)

	g.Image = imgs

	var opt gif.Options
	opt.NumColors = 256
	g.Delay = []int{100, 100, 100}
	g.LoopCount = 1000

	// you can add more parameters if you want // ok, write out the data into the new GIF fil
	err = gif.EncodeAll(out, g) // put num of colors to 256 <------- here
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Generated image to output.gif \n")
}
