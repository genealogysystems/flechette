package main

import (
	"fmt"
	"github.com/genealogysystems/flechette/server"
	//github.com/paulsmith/gogeos/geos"
)

func main() {
	fmt.Println(server.Version())
	/*
	   line, _ := geos.FromWKT("LINESTRING (0 0, 10 10, 20 20)")
	   buf, _ := line.Buffer(2.5)
	   fmt.Println(buf)
	   // Output: POLYGON ((18.2322330470336311 21.7677669529663689, 18.61…
	*/
}
