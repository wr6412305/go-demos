package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X, Y int
}

func (pt Point) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"X":%d,"Y":%d}`, pt.X, pt.Y)), nil
}

func marshal3() {
	if data, err := json.Marshal(Point{50, 50}); err == nil {
		fmt.Printf("%s\n", data)
	}
}
