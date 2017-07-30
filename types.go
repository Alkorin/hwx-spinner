package main

import (
	"bytes"
)

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

type Configuration struct {
	Text1   Text
	Text2   Text
	Text3   Text
	Message Message
}

func (c Configuration) Bytes() []byte {
	var buffer bytes.Buffer

	// First Image
	buffer.Write([]byte{58, 1, 0, 243, 243, 0, 1, 1})
	buffer.Write(make([]byte, 240))
	buffer.Write([]byte{14, 12, 13})

	// Second Image
	buffer.Write([]byte{58, 1, 1, 243, 243, 0, 1, 1})
	buffer.Write(make([]byte, 240))
	buffer.Write([]byte{14, 12, 13})

	// Third Image
	buffer.Write([]byte{58, 1, 2, 243, 243, 0, 1, 1})
	buffer.Write(make([]byte, 240))
	buffer.Write([]byte{14, 12, 13})

	// First Text
	buffer.Write([]byte{58, 1, 3})
	buffer.Write(c.Text1.Bytes())
	buffer.Write([]byte{14, 12, 13})

	// Second Text
	buffer.Write([]byte{58, 1, 4})
	buffer.Write(c.Text2.Bytes())
	buffer.Write([]byte{14, 12, 13})

	// Third Text
	buffer.Write([]byte{58, 1, 5})
	buffer.Write(c.Text3.Bytes())
	buffer.Write([]byte{14, 12, 13})

	// Version..Status..
	buffer.Write([]byte{58, 1, 6})
	buffer.Write(c.Message.Bytes())
	buffer.Write([]byte{14, 12, 13})

	// Font
	buffer.Write([]byte{58, 1, 9, 0, 0})
	buffer.Write([]byte{14, 12, 13})

	return buffer.Bytes()
}
