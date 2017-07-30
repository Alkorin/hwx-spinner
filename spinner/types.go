package Spinner

import (
	"bytes"
	"fmt"
)

//go:generate enumer -type=Color -yaml
//go:generate enumer -type=Mode -yaml
//go:generate enumer -type=Type -yaml

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
