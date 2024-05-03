package main

import "fmt"

func main() {
	tv1 := &SammysangTV{
		currentChan:   1,
		currentVolume: 10,
		tvOn:          false,
	}
	tv2 := &SohneeTV{
		vol:     10,
		channel: 1,
		isOn:    true,
	}
	tv2.channelUp()
	tv2.volumeDown()
	tv2.goToChannel(11)
	tv2.channelDown()
	tv2.volumeUp()
	tv2.turnOn()

	fmt.Println("---------------------------")
	ssAdapter := &sammysangAdapter{
		ssTv: tv1,
	}

	ssAdapter.channelUp()
	ssAdapter.volumeDown()
	ssAdapter.goToChannel(11)
	ssAdapter.channelDown()
	ssAdapter.volumeUp()
	ssAdapter.turnOn()
}
