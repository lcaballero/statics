package handling

import (
	. "github.com/lcaballero/exam/assert"
	"testing"
)

func Test_Parts_008(t *testing.T) {
	t.Log("nigh all with 0 element should produce empty Parts")
	p := Parts{}

	g := p.NighAll()
	IsTrue(t, g.IsEmpty())
	IsEqInt(t, g.Len(), 0)
	IsEqStrings(t, g.First(), "")
	IsEqStrings(t, g.Last(), "")
}

func Test_Parts_007(t *testing.T) {
	t.Log("nigh all with 1 element should produce empty Parts")
	p := Parts{"a"}

	g := p.NighAll()
	IsTrue(t, g.IsEmpty())
	IsEqInt(t, g.Len(), 0)
	IsEqStrings(t, g.First(), "")
	IsEqStrings(t, g.Last(), "")
}

func Test_Parts_006(t *testing.T) {
	t.Log("nigh all should produce all but the last element")
	a, b, c, d, e := "a", "b", "c", "d", "e"
	p := Parts{a, b, c, d, e}

	g := p.NighAll()
	IsFalse(t, g.IsEmpty())
	IsEqInt(t, g.Len(), 4)
	IsEqStrings(t, g.First(), a)
	IsEqStrings(t, g.Last(), d)
}

func Test_Parts_005(t *testing.T) {
	t.Log("5 len Parts should NOT be empty")
	a, b, c, d, e := "a", "b", "c", "d", "e"
	p := Parts{a, b, c, d, e}
	IsFalse(t, p.IsEmpty())
	IsEqInt(t, p.Len(), 5)
	IsEqStrings(t, p.First(), a)
	IsEqStrings(t, p.Last(), e)
}

func Test_Parts_004(t *testing.T) {
	t.Log("2 len Parts should NOT be empty")
	a := "a"
	b := "b"
	p := Parts{a, b}
	IsFalse(t, p.IsEmpty())
	IsEqInt(t, p.Len(), 2)
	IsEqStrings(t, p.First(), a)
	IsEqStrings(t, p.Last(), b)
}

func Test_Parts_003(t *testing.T) {
	t.Log("1 len Parts should NOT be empty")
	a := "a"
	p := Parts{a}
	IsFalse(t, p.IsEmpty())
	IsEqInt(t, p.Len(), 1)
	IsEqStrings(t, p.First(), a)
	IsEqStrings(t, p.Last(), a)
}

func Test_Parts_002(t *testing.T) {
	t.Log("0 len Parts should be empty")
	p := Parts{}
	IsTrue(t, p.IsEmpty())
	IsZero(t, p.Len())
	IsEmptyString(t, p.First())
	IsEmptyString(t, p.Last())
}

func Test_Parts_001(t *testing.T) {
	t.Log("nil Parts should be empty")
	var p Parts
	IsTrue(t, p.IsEmpty())
	IsZero(t, p.Len())
	IsEmptyString(t, p.First())
	IsEmptyString(t, p.Last())
}
