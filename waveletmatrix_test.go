package waveletmatrix

import (
	"testing"
)

func TestBuildAndAccess(t *testing.T) {
	src := []uint64{5, 1, 0, 4, 2, 2, 0, 3}
	wm, _ := NewWM(src)
	if wm.Size() != uint64(len(src)) {
		t.Error("Exprected", len(src), "Got", wm.Size())
	}
	for i := 0; i < len(src); i++ {
		v, found := wm.Lookup(uint64(i))
		if !found {
			t.Error("Not Found:", i)
		}
		if v != src[i] {
			t.Error("Expected", src[i], "Got", v)
		}
	}
	if v, found := wm.Lookup(uint64(len(src))); found {
		t.Error("Unexpected", v)
	}

	// Rank
	if r, _ := wm.Rank(0, 0); r != uint64(0) {
		t.Error("Expected", 0, "Got", r)
	}
	if r, _ := wm.Rank(3, 6); r != uint64(0) {
		t.Error("Expected", 0, "Got", r)
	}
	if r, _ := wm.Rank(0, 6); r != uint64(1) {
		t.Error("Expected", 1, "Got", r)
	}
	if r, _ := wm.Rank(0, 7); r != uint64(2) {
		t.Error("Expected", 2, "Got", r)
	}
	if r, _ := wm.Rank(2, 6); r != uint64(2) {
		t.Error("Expected", 2, "Got", r)
	}
	if r, _ := wm.Rank(5, 6); r != uint64(1) {
		t.Error("Expected", 1, "Got", r)
	}
	if _, found := wm.Rank(1, 10); found {
		t.Error("Expected", false, "Got", found)
	}

	// Select
	if pos, _ := wm.Select(2, 1); pos != uint64(4) {
		t.Error("Expected", 4, "Got", pos)
	}
	if pos, _ := wm.Select(2, 2); pos != uint64(5) {
		t.Error("Expected", 5, "Got", pos)
	}
	// c >= alphabetNum
	if _, found := wm.Select(10, 1); found {
		t.Error("Unexpected")
	}
	// Invalid rank
	if _, found := wm.Select(1, 2); found {
		t.Error("Unexpected")
	}
	if pos, _ := wm.SelectFromPos(0, 3, 1); pos != uint64(6) {
		t.Error("Expected", 6, "Got", pos)
	}
	if _, found := wm.SelectFromPos(0, 3, 2); found {
		t.Error("Unexpected")
	}


	if r := wm.RankLessThan(4, 5); r != uint64(3) {
		t.Error("Expected", 3, "Got", r)
	}
	if r := wm.RankMoreThan(3, 5); r != uint64(2) {
		t.Error("Expected", 2, "Got", r)
	}
	r, rl, rm := wm.RankAll(10, 0, 6)
	if r != NotFound || rl != NotFound || rm != NotFound {
		t.Error("Unexpected")
	}

	if f := wm.Freq(2); f != uint64(2) {
		t.Error("Expected", 2, "Got", f)
	}

	if f := wm.FreqRange(2, 5, 2, 6); f != uint64(3) {
		t.Error("Expected", 3, "Got", f)
	}
	if f := wm.FreqRange(10, 11, 1, 6); f != uint64(0) {
		t.Error("Expected", 0, "Got", f)
	}
	if f := wm.FreqRange(3, 2, 1, 6); f != uint64(0) {
		t.Error("Expected", 0, "Got", f)
	}
	if f := wm.FreqRange(1, 2, 7, 6); f != uint64(0) {
		t.Error("Expected", 0, "Got", f)
	}

	if f := wm.FreqSum(0, 3); f != uint64(5) {
		t.Error("Expected", 5, "Got", f)
	}

	pos, val := wm.MaxRange(1, 6)
	if pos != uint64(3) {
		t.Error("Expected", 3, "Got", pos)
	}
	if val != uint64(4) {
		t.Error("Expected", 4, "Got", val)
	}

	pos, val = wm.MinRange(1, 6)
	if pos != uint64(2) {
		t.Error("Expected", 2, "Got", pos)
	}
	if val != uint64(0) {
		t.Error("Expected", 0, "Got", val)
	}

	pos, val = wm.QuantileRange(1, 6, 3)
	if pos != uint64(5) {
		t.Error("Expected", 5, "Got", pos)
	}
	if val != uint64(2) {
		t.Error("Expected", 2, "Got", val)
	}

	result := wm.ListModeRange(1, 3, 0, 8, 3)
	if size := len(result); size != 2 {
		t.Error("Expected", 2, "Got", size)
	}
	if result[0].C != uint64(2) {
		t.Error("Expected", 2, "Got", result[0].C)
	}
	if result[0].Freq != uint64(2) {
		t.Error("Expected", 2, "Got", result[0].Freq)
	}
	if result[1].C != uint64(1) {
		t.Error("Expected", 1, "Got", result[1].C)
	}
	if result[1].Freq != uint64(1) {
		t.Error("Expected", 1, "Got", result[1].Freq)
	}

	result = wm.ListMaxRange(1, 5, 0, 8, 3)
	if size := len(result); size != 3 {
		t.Error("Expected", 3, "Got", size)
	}
	if result[0].C != uint64(4) {
		t.Error("Expected", 4, "Got", result[0].C)
	}
	if result[0].Freq != uint64(1) {
		t.Error("Expected", 1, "Got", result[0].Freq)
	}
	if result[1].C != uint64(3) {
		t.Error("Expected", 3, "Got", result[1].C)
	}
	if result[1].Freq != uint64(1) {
		t.Error("Expected", 1, "Got", result[1].Freq)
	}
	if result[2].C != uint64(2) {
		t.Error("Expected", 2, "Got", result[2].C)
	}
	if result[2].Freq != uint64(2) {
		t.Error("Expected", 2, "Got", result[2].Freq)
	}
	result = wm.ListMinRange(0, 5, 0, 8, 3)
	if size := len(result); size != 3 {
		t.Error("Expected", 3, "Got", size)
	}
	if result[0].C != uint64(0) {
		t.Error("Expected", 0, "Got", result[0].C)
	}
	if result[0].Freq != uint64(2) {
		t.Error("Expected", 2, "Got", result[0].Freq)
	}
	if result[1].C != uint64(1) {
		t.Error("Expected", 1, "Got", result[1].C)
	}
	if result[1].Freq != uint64(1) {
		t.Error("Expected", 1, "Got", result[1].Freq)
	}
	if result[2].C != uint64(2) {
		t.Error("Expected", 0, "Got", result[2].C)
	}
	if result[2].Freq != uint64(2) {
		t.Error("Expected", 2, "Got", result[2].Freq)
	}

}

func TestMarshalize(t *testing.T) {
	src := []uint64{5, 1, 0, 4, 2, 2, 0, 3}
	wm, _ := NewWM(src)
	
	buf, _ := wm.MarshalBinary()
	wm2, err := NewWMFromBinary(buf)
	if err != nil {
		t.Error("Unexpected error in UnmarshalBinary()")
	}

	if wm2.Size() != uint64(len(src)) {
		t.Error("Exprected", len(src), "Got", wm2.Size())
	}
	for i := 0; i < len(src); i++ {
		v, found := wm2.Lookup(uint64(i))
		if !found {
			t.Error("Not Found:", i)
		}
		if v != src[i] {
			t.Error("Expected", src[i], "Got", v)
		}
	}
}
