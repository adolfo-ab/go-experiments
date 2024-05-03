package main

import "fmt"

type sammysangAdapter struct {
	ssTv *SammysangTV
}

func (ss *sammysangAdapter) turnOn() {
	fmt.Println("SammySang is now on")
	ss.ssTv.tvOn = true
}

func (ss *sammysangAdapter) turnOff() {
	fmt.Println("SammySang is now off")
	ss.ssTv.tvOn = false
}

func (ss *sammysangAdapter) volumeUp() int {
	ss.ssTv.setVolume(ss.ssTv.getVolume() + 1)
	fmt.Println("Increasing SammySang volume to", ss.ssTv.getVolume())
	return ss.ssTv.getVolume()
}

func (ss *sammysangAdapter) volumeDown() int {
	ss.ssTv.setVolume(ss.ssTv.getVolume() - 1)
	fmt.Println("Increasing SammySang volume to", ss.ssTv.getVolume())
	return ss.ssTv.getVolume()
}

func (ss *sammysangAdapter) channelUp() int {
	ss.ssTv.setChannel(ss.ssTv.getChannel() + 1)
	fmt.Println("Increasing SammySang channel to", ss.ssTv.getChannel())
	return ss.ssTv.getVolume()
}

func (ss *sammysangAdapter) channelDown() int {
	ss.ssTv.setChannel(ss.ssTv.getChannel() - 1)
	fmt.Println("Increasing SammySang channel to", ss.ssTv.getChannel())
	return ss.ssTv.getChannel()
}

func (ss *sammysangAdapter) goToChannel(ch int) {
	ss.ssTv.setChannel(ch)
	fmt.Println("Setting SammySang channel to", ss.ssTv.getChannel())
}
