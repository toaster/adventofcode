package sif

import (
	"bytes"
	"math"
	"strconv"
)

type Image struct {
	width  int
	height int
	layers [][]int
}

func New(width, height int, raw string) *Image {
	imgSize := width * height
	if len(raw)%imgSize != 0 {
		panic("invalid input")
	}
	img := &Image{width: width, height: height}
	for li := 0; li < len(raw)/imgSize; li++ {
		l := []int{}
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				pi := y*width + x
				ii := pi + imgSize*li
				p, _ := strconv.Atoi(raw[ii : ii+1])
				l = append(l, p)
			}
		}
		img.layers = append(img.layers, l)
	}
	return img
}

func CountOfOnesByTwosOnLayerWithFewestZeros(width, height int, raw string) int {
	img := New(width, height, raw)
	var layer int
	minZeros := math.MaxInt64
	for li := 0; li < len(img.layers); li++ {
		if zeros := img.countDigitsOnLayer(0, li); zeros < minZeros {
			layer = li
			minZeros = zeros
		}
	}
	return img.countDigitsOnLayer(1, layer) * img.countDigitsOnLayer(2, layer)
}

func RenderImage(width, height int, raw string) string {
	img := New(width, height, raw)
	out := bytes.NewBuffer([]byte{})
	for y := 0; y < height; y++ {
	Loop:
		for x := 0; x < width; x++ {
			pi := y*img.width + x
			for _, l := range img.layers {
				px := l[pi]
				if px == 2 {
					continue
				}
				if px == 1 {
					out.WriteString("X")
				} else {
					out.WriteString(".")
				}
				continue Loop
			}
			out.WriteString(" ")
		}
		out.WriteString("\n")
	}
	return out.String()
}

func (i *Image) countDigitsOnLayer(digit, layer int) int {
	count := 0

	for _, pixel := range i.layers[layer] {
		if pixel == digit {
			count++
		}
	}
	return count
}
