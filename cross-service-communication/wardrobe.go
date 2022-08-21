package main

// Garment is a single piece of clothing
type Garment struct {
	id     uint64
	status string
}

// Wardrobe contains all the known garments
type Wardrobe struct {
	maxId   uint64
	garment map[uint64]Garment
}

func InitWardobe() *Wardrobe {
	return &Wardrobe{garment: make(map[uint64]Garment)}
}

func (wardrobe *Wardrobe) AddGarment() {
	// if no garment in wardrobe - add garment
	if _, ok := wardrobe.garment[0]; !ok {
		wardrobe.garment[0] = Garment{id: 0, status: "dirty"}
		wardrobe.maxId = 0
		return
	}
	// otherwise: add a new garment and increment maxId
	newId := wardrobe.maxId + 1
	wardrobe.garment[newId] = Garment{id: newId, status: "dirty"}
	wardrobe.maxId = newId
}
