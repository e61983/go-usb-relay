package relay

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestList_OneDevice_GetOne(t *testing.T) {
	l := List()
	assert.Equal(t, 1, len(l))
}

func TestTurnAllOn(t *testing.T) {
	l := List()
	relay := l[0]
	err := relay.Open()
	assert.Nil(t, err)
	defer relay.Close()
	err = relay.TurnAllOn()
	time.Sleep(time.Millisecond * 500)
	assert.Nil(t, err)
}

func TestTurnAllOff(t *testing.T) {
	l := List()
	relay := l[0]
	err := relay.Open()
	assert.Nil(t, err)
	defer relay.Close()
	err = relay.TurnAllOff()
	time.Sleep(time.Millisecond * 500)
	assert.Nil(t, err)
}

func TestTurnOn(t *testing.T) {
	l := List()
	relay := l[0]
	err := relay.Open()
	assert.Nil(t, err)
	defer relay.Close()
	for i := 1; i <= 8; i++ {
		time.Sleep(time.Millisecond * 500)
		err = relay.TurnOn(ChannelNumber(i))
		assert.Nil(t, err)
	}
}

func TestGetStatus(t *testing.T) {
	l := List()
	relay := l[0]
	err := relay.Open()
	assert.Nil(t, err)
	defer relay.Close()
	status, err := relay.GetStatus()
	assert.Nil(t, err)
	expected := &ChannelStatus{Channel_1: 1, Channel_2: 1, Channel_3: 1, Channel_4: 1, Channel_5: 1, Channel_6: 1, Channel_7: 1, Channel_8: 1}
	assert.Equal(t, expected, status)
}

func TestTurnOff(t *testing.T) {
	l := List()
	relay := l[0]
	err := relay.Open()
	assert.Nil(t, err)
	defer relay.Close()
	for i := 1; i <= 8; i++ {
		time.Sleep(time.Millisecond * 500)
		err = relay.TurnOff(ChannelNumber(i))
		assert.Nil(t, err)
	}
}
func TestSetSN_AVI_01_GetErrorMessage(t *testing.T) {
	l := List()
	relay := l[0]
	err := relay.Open()
	assert.Nil(t, err)
	defer relay.Close()
	err = relay.SetSN("AVI_01")
	expected := errors.New("The length of `AVI_01` is large than 5 bytes.")
	assert.Equal(t, expected, err)
}

func TestSetSN_AVI01_OK(t *testing.T) {
	l := List()
	relay := l[0]
	err := relay.Open()
	assert.Nil(t, err)
	defer relay.Close()
	err = relay.SetSN("AVI01")
	assert.Nil(t, err)
	sn, err := relay.GetSN()
	assert.Nil(t, err)
	assert.Equal(t, "AVI01", sn)
}
