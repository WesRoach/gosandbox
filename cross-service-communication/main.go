package main

import (
	"fmt"
	"time"
)

func main() {

	// Create a bunch of garments
	wardrobe := InitWardobe()
	for i := 0; i < 20; i++ {
		wardrobe.AddGarment()
	}

	// Start Services
	intoWash := make(chan *Garment)
	intoDry := make(chan *Garment)
	intoFold := make(chan *Garment)
	go ServiceWash(intoWash, intoDry)
	go ServiceDry(intoDry, intoFold)
	go ServiceFold(intoFold)

	// Print all garments in wardrobe
	for garmentId := range wardrobe.garment {
		fmt.Printf("garmentId: %v\n", garmentId)
	}
	fmt.Printf("%+v, %+v\n", cap(wardrobe.garment), len(wardrobe.garment))

	// Send address of Garments from wardrobe into ServiceWash
	fmt.Printf("Starting to send Garments into Wash.\n")
	for garmentId := range wardrobe.garment {
		// time.Sleep(time.Second * 1)
		fmt.Printf("%v: ServiceWorker sending %+v into Wash.\n", time.Now(), garmentId)
		intoWash <- &wardrobe.garment[garmentId]
	}

}
