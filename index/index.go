package index

import (
	"fmt"
	"github.com/dhconnelly/rtreego"
)

func NewElement(id string) *Element {

	p1 := rtreego.Point{0.4, 0.5}
	r1, _ := rtreego.NewRect(p1, []float64{1, 2})

	return &Element{id, r1}

}

type Element struct {
	id    string
	where *rtreego.Rect
}

func (t *Element) Bounds() *rtreego.Rect {
	return t.where
}

func NewIndex() *Index {
	rt := rtreego.NewTree(2, 25, 50)
	/*
		p1 := rtreego.Point{0.4, 0.5}
		p2 := rtreego.Point{6.2, -3.4}
		r1, _ := rtreego.NewRect(p1, []float64{1, 2})
		r2, _ := rtreego.NewRect(p2, []float64{1.7, 2.7})

		rt.Insert(&Element{"123", r1})
		rt.Insert(&Element{"abc", r2})

		fmt.Println(rt.Size())
	*/
	return &Index{rt}
}

type Index struct {
	rtree *rtreego.Rtree
}

func (t *Index) Insert(elem *Element) {
	t.rtree.Insert(elem)
	fmt.Println(elem.id)
	fmt.Println(t.rtree.Size())
	return
}
