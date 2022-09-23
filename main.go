package main

import (
	"github.com/edwardsdl/turbo-sniffle/buhlmann"
	"log"
)

func main() {
	zhl16, err := buhlmann.NewZHL16("zhl16c.json")
	if err != nil {
		log.Fatalln(err)
	}

	zhl16.AddStop(4, 10)

	for i, c := range zhl16.Compartments {
		log.Printf("Compartment %v: %v", i, c.PPN2)
	}
}
