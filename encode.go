package main

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path"
)

func main() {

	out, err := os.Create("./output.gif")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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

	err = gif.EncodeAll(out, g)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Generated image to output.gif \n")
}
