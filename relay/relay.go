package relay

import (
	"fmt"
	"github.com/karalabe/hid"
	"runtime"
)

const (
	OFF = iota
	ON
)

type IoStatus int

const (
	C1 = iota + 1
	C2
	C3
	C4
	C5
	C6
	C7
	C8
	ALL
)

type ChannelNumber int

type ChannelStatus struct {
	Channel_1 IoStatus
	Channel_2 IoStatus
	Channel_3 IoStatus
	Channel_4 IoStatus
	Channel_5 IoStatus
	Channel_6 IoStatus
	Channel_7 IoStatus
	Channel_8 IoStatus
}

type Relay struct {
	info *hid.DeviceInfo
	dev  *hid.Device
}

func List() []*Relay {
	list := make([]*Relay, 0, 5)
	relayInfos := hid.Enumerate(0x16C0, 0x05DF)

	if len(relayInfos) <= 0 {
		return list
	}
	for i, info := range relayInfos {
		relay, err := info.Open()
		if err == nil {
			list = append(list, &Relay{info: &relayInfos[i]})
		}
		relay.Close()
	}
	return list
}

func (this *Relay) Open() error {
	dev, err := this.info.Open()
	if err == nil {
		this.dev = dev
	}
	return err
}

func (this *Relay) Close() error {
	return this.dev.Close()
}

func (this *Relay) setIO(s IoStatus, no ChannelNumber) error {
	cmd := make([]byte, 9)
	cmd[0] = 0x0
	if no < C1 && no > ALL {
		return fmt.Errorf("channel number (%d) is illegal", no)
	}

	if no == ALL {
		if s == ON {
			cmd[1] = 0xFE
		} else {
			cmd[1] = 0xFC
		}
	} else {
		if s == ON {
			cmd[1] = 0xFF
		} else {
			cmd[1] = 0xFD
		}
		cmd[2] = byte(no)
	}

	_, err := this.dev.SendFeatureReport(cmd)
	return err
}

func (this *Relay) GetStatus() (*ChannelStatus, error) {
	cmd := make([]byte, 9)
	_, err := this.dev.GetFeatureReport(cmd)
	if err != nil {
		return nil, err
	}

	// Remove HID report ID on Windows, others OSes don't need it.
	if runtime.GOOS == "windows" {
		cmd = cmd[1:]
	}

	status := &ChannelStatus{}
	status.Channel_1 = IoStatus(cmd[7] >> 0 & 0x01)
	status.Channel_2 = IoStatus(cmd[7] >> 1 & 0x01)
	status.Channel_3 = IoStatus(cmd[7] >> 2 & 0x01)
	status.Channel_4 = IoStatus(cmd[7] >> 3 & 0x01)
	status.Channel_5 = IoStatus(cmd[7] >> 4 & 0x01)
	status.Channel_6 = IoStatus(cmd[7] >> 5 & 0x01)
	status.Channel_7 = IoStatus(cmd[7] >> 6 & 0x01)
	status.Channel_8 = IoStatus(cmd[7] >> 7 & 0x01)
	return status, err
}

func (this *Relay) TurnOn(num ChannelNumber) error {
	return this.setIO(ON, num)
}

func (this *Relay) TurnOff(num ChannelNumber) error {
	return this.setIO(OFF, num)
}

func (this *Relay) TurnAllOn() error {
	return this.setIO(ON, ALL)
}

func (this *Relay) TurnAllOff() error {
	return this.setIO(OFF, ALL)
}

func (this *Relay) SetSN(sn string) error {
	if len(sn) > 5 {
		return fmt.Errorf("The length of `%s` is large than 5 bytes.", sn)
	}
	cmd := make([]byte, 9)
	cmd[0] = 0x00
	cmd[1] = 0xFA
	copy(cmd[2:], sn)
	_, err := this.dev.SendFeatureReport(cmd)
	if err != nil {
		return err
	}
	return err
}

func (this *Relay) GetSN() (string, error) {
	cmd := make([]byte, 9)
	_, err := this.dev.GetFeatureReport(cmd)
	var sn string
	if err != nil {
		sn = ""
	} else {
		// Remove HID report ID on Windows, others OSes don't need it.
		if runtime.GOOS == "windows" {
			cmd = cmd[1:]
		}
		sn = string(cmd[:5])
	}
	return sn, err
}
