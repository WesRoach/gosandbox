package main

// Garment is a single piece of clothing
type Garment struct {
	id     uint64
	status string
}

// Wardrobe contains all the known garments
type Wardrobe struct {
	maxId   uint64
	garment []Garment
}

func InitWardobe() *Wardrobe {
	return &Wardrobe{garment: make([]Garment, 0)}
}

func (wardrobe *Wardrobe) AddGarment() {
	// if no garment in wardrobe - add garment
	if len(wardrobe.garment) == 0 {
		wardrobe.garment = append(wardrobe.garment, Garment{id: 0, status: "dirty"})
		wardrobe.maxId = 0
		return
	}
	// otherwise: add a new garment and increment maxId
	newId := wardrobe.maxId + 1
	wardrobe.garment = append(wardrobe.garment, Garment{id: newId, status: "dirty"})
	wardrobe.maxId = newId
}
