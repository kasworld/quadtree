// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package quadtree

import (
	"testing"

	// "github.com/kasworld/ivector2d"
	"github.com/kasworld/idgen"
	"github.com/kasworld/rect"
)

type qtobj struct {
	id idgen.IDInt
	rt rect.Rect
}

func (po *qtobj) GetID() idgen.IDInt {
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
