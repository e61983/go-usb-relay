package main

import (
	"flag"
	"log"
	"os"

	"github.com/e61983/go-usb-relay/relay"
)

var (
	sn    = flag.String("sn", "", "Set module SN.")
	num   = flag.Int("n", 1, "Selected relay channel")
	onOff = flag.Bool("o", false, "on or off")
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	// List all relay module
	l := relay.List()

	if len(l) <= 0 {
		log.Fatal("No device")
	}

	r1 := l[0]

	err := r1.Open()

	if err != nil {
		log.Fatal(err)
	}
	defer r1.Close()

	if *sn != "" {
		err = r1.SetSN(*sn)
		if err != nil {
			log.Fatal(err)
		} else {
			os.Exit(0)
		}
	}

	if *onOff {
		log.Println("Channel", *num, "ON")
		r1.TurnOn(relay.ChannelNumber(*num))
	} else {
		log.Println("Channel", *num, "OFF")
		r1.TurnOff(relay.ChannelNumber(*num))
	}
}
