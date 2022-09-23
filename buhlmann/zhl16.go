package buhlmann

import (
	"encoding/json"
	"io/ioutil"
)

type ZHL16 struct {
	Compartments [16]Compartment
}

func NewZHL16(filename string) (*ZHL16, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	zhl16 := ZHL16{}

	err = json.Unmarshal(data, &zhl16.Compartments)
	if err != nil {
		return nil, err
	}

	return &zhl16, nil
}

func (z *ZHL16) AddStop(atm float64, time float64) {
	for i := range z.Compartments {
		z.Compartments[i].AddStop(atm, time)
	}
}
