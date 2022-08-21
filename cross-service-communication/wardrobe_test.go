package main

import (
	"reflect"
	"testing"
)

func TestInitWardrobe(t *testing.T) {
	res := InitWardobe()
	if reflect.TypeOf(res) != reflect.TypeOf(&Wardrobe{}) {
		t.Errorf("Expected InitWardrobe to return Wardrobe, returned %t", reflect.TypeOf(res))
	}
}

func TestAddGarment(t *testing.T) {
	wardrobe := InitWardobe()
	wardrobe.AddGarment()
	wardrobe.AddGarment()
	if wardrobe.maxId != 1 {
		t.Errorf("Expected maxId == 1, Found %+v", wardrobe.maxId)
	}
	if wardrobe.garment[0].status != "dirty" {
		t.Errorf("Expected garment to be dirty, Found %+v", wardrobe.garment[0])
	}
}
