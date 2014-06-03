package index

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	var idx = NewIndex()
	fmt.Println(idx.rtree.Size())
}
