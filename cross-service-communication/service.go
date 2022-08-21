package main

import (
	"fmt"
	"time"
)

func ServiceWash(wardrobe *Wardrobe, receiveChannel chan uint64, sendChannel chan uint64) {
	for {
		garment := <-receiveChannel
		fmt.Printf("%v: ServiceWash received %+v.\n", time.Now(), garment)
		time.Sleep(time.Second * 1)
		sendChannel <- garment
		fmt.Printf("%v: ServiceWash passed %+v to Dryer.\n", time.Now(), garment)
	}
}

func ServiceDry(wardrobe *Wardrobe, receiveChannel chan uint64, sendChannel chan uint64) {
	for {
		garment := <-receiveChannel
		fmt.Printf("%v: ServiceDry received %+v.\n", time.Now(), garment)
		time.Sleep(time.Second * 2)
		sendChannel <- garment
		fmt.Printf("%v: ServiceDry passed %+v to Fold.\n", time.Now(), garment)
	}
}

func ServiceFold(wardrobe *Wardrobe, receiveChannel chan uint64) {
	for {
		garment := <-receiveChannel
		fmt.Printf("%v: ServiceFold received %+v.\n", time.Now(), garment)
		time.Sleep(time.Second * 1)
	}
}
