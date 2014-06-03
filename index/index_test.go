package index

import (
	"testing"
)

func TestNewIndex(t *testing.T) {
	var (
		idx = NewIndex()
		size = idx.rtree.Size()
	)

	if size != 0 {
		t.Error("Expected 0, got ", size)
	}
}
