package main

import "fmt"

type SammysangTV struct {
	currentChan   int
	currentVolume int
	tvOn          bool
}

func (tv *SammysangTV) getVolume() int {
	fmt.Println("SammysangTV volume is", tv.currentVolume)
	return tv.currentVolume
}

func (tv *SammysangTV) setVolume(volume int) {
	fmt.Println("Setting SammysangTV volume to", volume)
	tv.currentVolume = volume
}

func (tv *SammysangTV) getChannel() int {
	fmt.Println("SammysangTV channel is", tv.currentChan)
	return tv.currentChan
}

func (tv *SammysangTV) setChannel(channel int) {
	fmt.Println("Setting SammysangTV channel to", channel)
	tv.currentChan = channel
}

func (tv *SammysangTV) setOnState(tvOn bool) {
	if tv.tvOn == true {
		fmt.Println("SammysangTV is on")
	} else {
		fmt.Println("SammysangTV is off")
	}
	tv.tvOn = tvOn
}
