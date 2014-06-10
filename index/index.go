package index

import (
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
func NewElement(id string, geom *geos.Geometry) (elem *Element, geosError error) {

	var (
		minX float64
		minY float64
		maxX float64
		maxY float64
		err error
	)

	// Get the envelope
	bounds, err := geom.Envelope()
	if err != nil {
    	return nil, err
    }
	shell, err := bounds.Shell()
	if err != nil {
    	return nil, err
    }
	points, err := shell.NPoint()
	if err != nil {
    	return nil, err
    }

    // Set the initial point
	initialPoint, err := shell.Point(0)
	if err != nil {
    	return nil, err
    }
	minX, err = initialPoint.X()
	if err != nil {
    	return nil, err
    }
	minY, err = initialPoint.Y()
	if err != nil {
    	return nil, err
    }
	maxX, err = initialPoint.X()
	if err != nil {
    	return nil, err
    }
	maxY, err = initialPoint.Y()

	// Get min and max x/y
	for i := 1; i < points; i++ {
		point, err := shell.Point(i)
		if err != nil {
	    	return nil, err
	    }
		x, err := point.X()
		if err != nil {
	    	return nil, err
	    }
		y, err := point.Y()
		if err != nil {
	    	return nil, err
	    }
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
	r1, err := rtreego.NewRect(rtreego.Point{minX,minY}, []float64{maxX-minX,maxY-minY})
	if err != nil {
    	return nil, err
    }

	pgeom := geom.Prepare()

	return &Element{id, pgeom, r1}, nil

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