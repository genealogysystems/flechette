package main

import (
	"fmt"
	"github.com/genealogysystems/flechette/index"
	"github.com/genealogysystems/flechette/server"
	//github.com/paulsmith/gogeos/geos"
)

func main() {
	fmt.Println(server.Version())
	var idx = index.NewIndex()
	var elems = []*index.Element{
		index.NewElement("123"),
		index.NewElement("456"),
	}

	for _, elem := range elems {
		idx.Insert(elem)
	}
	/*
	   line, _ := geos.FromWKT("LINESTRING (0 0, 10 10, 20 20)")
	   buf, _ := line.Buffer(2.5)
	   fmt.Println(buf)
	   // Output: POLYGON ((18.2322330470336311 21.7677669529663689, 18.61…
	*/
}
