package main

import (
	"bytes"
	"log"

	"github.com/google/gousb"
)

const VID = 0x28e9
const PID = 0x028a

func getBuffer(t1 string, t2 string, t3 string) []byte {
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
	buffer.Write([]byte{58, 1, 3, 0, byte(len(t1) + 3), 1, 1, 1})
	buffer.Write([]byte(t1))
	buffer.Write([]byte{14, 12, 13})

	// Second Text
	buffer.Write([]byte{58, 1, 4, 0, byte(len(t2) + 3), 1, 2, 1})
	buffer.Write([]byte(t2))
	buffer.Write([]byte{14, 12, 13})

	// Third Text
	buffer.Write([]byte{58, 1, 5, 0, byte(len(t3) + 3), 1, 3, 1})
	buffer.Write([]byte(t3))
	buffer.Write([]byte{14, 12, 13})

	// Version..Status..
	buffer.Write([]byte{58, 1, 6, 0, 4, 0, 1, 1})
	buffer.Write([]byte{0})
	buffer.Write([]byte{14, 12, 13})

	// Font
	buffer.Write([]byte{58, 1, 9, 0, 0})
	buffer.Write([]byte{14, 12, 13})

	return buffer.Bytes()
}

func main() {
	// List devices
	log.Print("Searching Handspinner\n")

	ctx := gousb.NewContext()
	defer ctx.Close()

	dev, err := ctx.OpenDeviceWithVIDPID(VID, PID)
	if err != nil {
		log.Printf("Failed to find device: %s", err.Error())
		return
	}
	if dev == nil {
		log.Print("No device found")
		return
	}
	defer dev.Close()

	log.Printf("Handspinner found! (%s)", dev.String())

	// Configuring device
	log.Print("Configuring device")
	dev.SetAutoDetach(true)
	conf, err := dev.Config(1)
	if err != nil {
		log.Printf("Failed to configure device :%s", err.Error())
		return
	}
	defer conf.Close()
	log.Printf("Handspinner configured! (%s)", conf.String())

	log.Print("Configuring interface")
	intf, err := conf.Interface(0, 0)
	if err != nil {
		log.Printf("Failed to configure interface :%s", err.Error())
		return
	}
	defer intf.Close()

	log.Printf("Handspinner interface configured! (%s)", intf.String())

	writer, err := intf.OutEndpoint(1)
	if err != nil {
		log.Printf("Failed to open write endpoint: %s", err.Error())
		return
	}

	log.Printf("Writer acquired: %+v", writer)

	// Construct Buffer
	buf := getBuffer("Hello", "World", "Toms")

	// Write Data
	log.Printf("Writing data...")
	for i := 0; i < len(buf); i += 32 {
		toSend := buf[i:]
		if len(toSend) > 32 {
			toSend = toSend[:32]
		}
		n, err := writer.Write(toSend)
		if err != nil {
			log.Printf("Failed to write data: %s", err.Error())
		}
		log.Printf("%d bytes OK", n)
	}
	log.Print("Write done, closing device")
}
