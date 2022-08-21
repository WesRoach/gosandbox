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
	intoWash := make(chan uint64)
	intoDry := make(chan uint64)
	intoFold := make(chan uint64)
	go ServiceWash(wardrobe, intoWash, intoDry)
	go ServiceDry(wardrobe, intoDry, intoFold)
	go ServiceFold(wardrobe, intoFold)

	// Print all garments in wardrobe
	// for garmentId := range wardrobe.garment {
	// 	fmt.Printf("garmentId: %v\n", garmentId)
	// }

	// Send address of Garments from wardrobe into ServiceWash
	fmt.Printf("Starting to send Garments into Wash.\n")
	for garmentId := range wardrobe.garment {
		time.Sleep(time.Second * 1)
		fmt.Printf("%v: ServiceWorker sending %+v into Wash.\n", time.Now(), garmentId)
		intoWash <- garmentId
	}
}
