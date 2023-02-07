package main

import (
	"fmt"

	"github.com/imersao-full-cycle/simulator/application/route"
)

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringJson, _ := route.ExportJSONPositions()
	
	for _, v := range stringJson {
		fmt.Println(v)
	}
}
