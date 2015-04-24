package quadtree

import (
	"testing"

	// "github.com/kasworld/ivector2d"
	"github.com/kasworld/rect"
)

type qtobj struct {
	id int64
	rt rect.Rect
}

func (po *qtobj) GetID() int64 {
	return po.id
}

func (po *qtobj) GetRect() rect.Rect {
	return po.rt
}

func TestQueryByPos(t *testing.T) {
	qt := New(rect.Rect{0, 0, 32, 32})
	o := &qtobj{
		12,
		rect.Rect{8, 8, 8, 8},
	}
	qt.Insert(o)
	if qt.QueryByPos(nil, [2]int{4, 4}) != false {
		t.Fail()
	}
	if qt.QueryByPos(nil, [2]int{8, 8}) != true {
		t.Fail()
	}
	if qt.QueryByPos(nil, [2]int{16, 16}) != false {
		t.Fail()
	}
}
