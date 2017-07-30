package Spinner

import (
	"bytes"
	"io/ioutil"
	"path"
	"strings"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Image1  Image
	Image2  Image
	Image3  Image
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

	// Clean image paths if necessary
	if c.Image1.Enabled && !path.IsAbs(c.Image1.File) {
		c.Image1.File = path.Join(path.Dir(conf), c.Image1.File)
	}
	if c.Image2.Enabled && !path.IsAbs(c.Image2.File) {
		c.Image2.File = path.Join(path.Dir(conf), c.Image2.File)
	}
	if c.Image3.Enabled && !path.IsAbs(c.Image3.File) {
		c.Image3.File = path.Join(path.Dir(conf), c.Image3.File)
	}

	return &c, nil
}

func (c Configuration) Bytes() []byte {
	var buffer bytes.Buffer

	// First Image
	buffer.Write([]byte{58, 1, 0})
	buffer.Write(c.Image1.Bytes())
	buffer.Write([]byte{14, 12, 13})

	// Second Image
	buffer.Write([]byte{58, 1, 1})
	buffer.Write(c.Image2.Bytes())
	buffer.Write([]byte{14, 12, 13})

	// Third Image
	buffer.Write([]byte{58, 1, 2})
	buffer.Write(c.Image3.Bytes())
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

func (c Configuration) String() string {
	s := make([]string, 0)
	if c.Image1.Enabled {
		s = append(s, "Image1: "+c.Image1.String())
	}
	if c.Image2.Enabled {
		s = append(s, "Image2: "+c.Image2.String())
	}
	if c.Image3.Enabled {
		s = append(s, "Image3: "+c.Image3.String())
	}
	if c.Text1.Enabled {
		s = append(s, "Text1: "+c.Text1.String())
	}
	if c.Text2.Enabled {
		s = append(s, "Text2: "+c.Text2.String())
	}
	if c.Text3.Enabled {
		s = append(s, "Text3: "+c.Text3.String())
	}
	if c.Message.Enabled {
		s = append(s, "Message: "+c.Message.String())
	}
	return strings.Join(s, "\n")
}
