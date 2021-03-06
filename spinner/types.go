package Spinner

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"os"
)

//go:generate enumer -type=Color -yaml
//go:generate enumer -type=Mode -yaml
//go:generate enumer -type=Type -yaml
//go:generate enumer -type=Coordinates -yaml

type Color byte

const (
	RED Color = iota
	BLUE
	MIX
)

type Mode byte

const (
	LEFT_TO_RIGHT Mode = iota
	UP_TO_DOWN
	FIX
)

type Text struct {
	Enabled bool
	Value   string
	Color   Color
	Mode    Mode
}

func (t Text) Bytes() []byte {
	// { 16 bits len, Enabled, Color, Mode, Value }
	if t.Enabled {
		b := bytes.NewBuffer([]byte{0, byte(len(t.Value) + 3), 1, byte(t.Color + 1), byte(t.Mode + 1)})
		b.Write([]byte(t.Value))
		return b.Bytes()
	} else {
		return []byte{0, 3, 0, 1, 1}
	}
}

func (t Text) String() string {
	return fmt.Sprintf("{Value: %s, Color: %s, Mode: %s}", t.Value, t.Color, t.Mode)
}

type Coordinates int

const (
	POLAR Coordinates = iota
	CARTESIAN
)

type Image struct {
	Enabled     bool
	File        string
	Coordinates Coordinates
	Color       Color
	Mode        Mode
}

func (i Image) Bytes() []byte {
	// { 16 bits len, Enabled, Color, Mode, Value }
	if i.Enabled {
		b := bytes.NewBuffer([]byte{243, 243, 1, byte(i.Color + 1), byte(i.Mode + 1)})
		file, err := os.Open(i.File)
		if err != nil {
			log.Panicf("Failed to open file: %s", err.Error())
		}
		img, _, err := image.Decode(file)
		if err != nil {
			log.Panicf("Failed to decode file %s: %s", i.File, err.Error())
		}

		if i.Coordinates == CARTESIAN {
			// Resample img
			imgDx := float64(img.Bounds().Dx())
			imgDy := float64(img.Bounds().Dy())
			imgBiggestSide := math.Max(imgDx, imgDy)

			imgNew := image.NewRGBA(image.Rect(0, 0, 120, 16))
			for x := 0; x < 120; x++ {
				for y := 0; y < 16; y++ {
					theta := 2 * math.Pi * float64(x) / 120
					r := float64(26-y) / float64(26)

					imgX := int(r*math.Cos(theta)*imgBiggestSide/2 + imgDx/2)
					imgY := int(r*math.Sin(theta)*imgBiggestSide/2 + imgDy/2)
					imgNew.Set(x, y, img.At(imgX, imgY))
				}
			}

			img = imgNew
		}

		if img.Bounds().Dx() != 120 {
			log.Panicf("Image should be 120px wide")
		}
		if img.Bounds().Dy() != 16 {
			log.Panicf("Image should be 16px tall")
		}
		data := make([]byte, 0, 240)
		for x := 0; x < 120; x++ {
			for y1 := 0; y1 < 2; y1++ {
				var value byte
				for y2 := 0; y2 < 8; y2++ {
					value <<= 1
					r, g, b, _ := img.At(img.Bounds().Min.X+x, img.Bounds().Min.Y+8*y1+y2).RGBA()
					y, _, _ := color.RGBToYCbCr(uint8(r), uint8(g), uint8(b))
					if y > 128 {
						value |= 1
					}
				}
				data = append(data, value)
			}
		}
		b.Write(data)
		return b.Bytes()
	} else {
		b := bytes.NewBuffer([]byte{0, 243, 0, 1, 1})
		b.Write(make([]byte, 240))
		return b.Bytes()
	}
}

func (i Image) String() string {
	return fmt.Sprintf("{File: %s, Color: %s, Mode: %s, Coordinates: %s}", i.File, i.Color, i.Mode, i.Coordinates)
}

type Type byte

const (
	SPEED Type = iota
	VOLT
	VERSION
)

type Message struct {
	Enabled bool
	Type    Type
	Color   Color
	Mode    Mode
}

func (m Message) Bytes() []byte {
	// { 16 bits len, Enabled, Color, Mode, Value }
	if m.Enabled {
		return []byte{0, 4, 1, byte(m.Color + 1), byte(m.Mode + 1), byte(m.Type + 1)}
	} else {
		return []byte{0, 4, 0, 1, 1, 0}
	}
}

func (m Message) String() string {
	return fmt.Sprintf("{Type: %s, Color: %s, Mode: %s}", m.Type, m.Color, m.Mode)
}
