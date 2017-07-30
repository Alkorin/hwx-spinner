package Spinner

import (
	"bytes"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Text1   Text
	Text2   Text
	Text3   Text
	Message Message
}

func NewConfiguration(conf string) (*Configuration, error) {
	file, err := ioutil.ReadFile(conf)
	if err != nil {
		return nil, err
	}

	var c Configuration
	err = yaml.Unmarshal(file, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
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
