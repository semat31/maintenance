package main

import (
	"fmt"
	"github.com/semat31/maintenance/store"

	//"github.com/semat31/maintenance/code"
)

func main() {

	device:=store.Device{
		Code: "29VA1121",
		Name: "Valve",
		Categorie: "Грузоподъемная техника",
		Serial: "W42355 Y000624",
		Commissioning: "2015",
	}

	fmt.Println(device)

}
