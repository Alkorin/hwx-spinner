package main

import (
	"log"

	"github.com/google/gousb"
)

const VID = 0x28e9
const PID = 0x028a

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
	var c Configuration

	c.Text1.Enabled = true
	c.Text1.Value = "HelloWorld"
	c.Text1.Color = BLUE
	c.Text1.Mode = FIX

	c.Message.Enabled = true
	c.Message.Type = SPEED

	buf := c.Bytes()

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
