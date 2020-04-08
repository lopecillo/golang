package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var map_make map[string]Vertex

var map_literal = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var map_literal_short = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {

	map_make = make(map[string]Vertex)
	map_make["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(map_make["Bell Labs"])

	fmt.Println(map_literal["Bell Labs"])

	fmt.Println(map_literal_short["Bell Labs"])

}