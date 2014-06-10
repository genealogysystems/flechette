package index

import (
	_ "fmt"
	"github.com/dhconnelly/rtreego"
	"github.com/paulsmith/gogeos/geos"
)

type Element struct {
	id    string
	geom *geos.PGeometry
	envelope *rtreego.Rect
}

func (t *Element) Bounds() *rtreego.Rect {
	return t.envelope
}

// Create a new Element
func NewElement(id string, geom *geos.Geometry) *Element {

	// Get the envelope
	bounds, _ := geom.Envelope()
	shell, _ := bounds.Shell()
	points, _ := shell.NPoint()

	initialPoint, _ := shell.Point(0)
	minX, _ := initialPoint.X()
	minY, _ := initialPoint.Y()
	maxX, _ := initialPoint.X()
	maxY, _ := initialPoint.Y()

	// Get min and max x/y
	for i := 1; i < points; i++ {
		point, _ := shell.Point(i)
		x, _ := point.X()
		y, _ := point.Y()

		if x <= minX && y <= minY {
			minX = x
			minY = y
		}
		if x >= maxX && y >= maxY {
			maxX = x
			maxY = y
		}
	}

	// Create rectangle from min/max points
	r1, _ := rtreego.NewRect(rtreego.Point{minX,minY}, []float64{maxX-minX,maxY-minY})

	pgeom := geom.Prepare()

	return &Element{id, pgeom, r1}

}

type Index struct {
	rtree *rtreego.Rtree
}

func (t *Index) Insert(elem *Element) {
	t.rtree.Insert(elem)
}

func (t *Index) Remove(elem *Element) {
	t.rtree.Delete(elem)
}

func NewIndex() *Index {
	rt := rtreego.NewTree(2, 25, 50)
	return &Index{rt}
}