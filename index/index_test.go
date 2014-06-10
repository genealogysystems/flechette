package index

import (
	_ "fmt"
	"testing"
	"github.com/paulsmith/gogeos/geos"
)


func TestNewIndex(t *testing.T) {
	var (
		idx  = NewIndex()
		size = idx.rtree.Size()
	)

	if size != 0 {
		t.Error("Expected 0, got ", size)
	}
}


func TestNewElement(t *testing.T) {

	geom, _ := geos.FromWKT("LINESTRING (0 0, 10 10, 20 20)")

	elem, err := NewElement("123",geom)
	if err != nil {
		t.Error(err)
	}

	if elem == nil {
		t.Error("Expected elem to not be nil")
	}
}

func TestElementInsertion(t *testing.T) {
	
	idx  := NewIndex()

	geom, _ := geos.FromWKT("LINESTRING (0 0, 10 10, 20 20)")

	elem, err := NewElement("123",geom)
	if err != nil {
		t.Error(err)
	}

	idx.Insert(elem)
	size := idx.rtree.Size()

	if size != 1 {
		t.Error("Expected 1, got ", size)
	}
}


func BenchmarkInsetion(b *testing.B) {
	idx  := NewIndex()
    for i := 0; i < b.N; i++ {
    	geom, err := geos.FromWKT("LINESTRING (0 0, 10 10, 20 20)")
    	if err != nil {
        	b.Error(err)
        }
        elem, err := NewElement("123",geom)
        if err != nil {
        	b.Error(err)
        }
        idx.Insert(elem)
    }
}
