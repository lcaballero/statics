package handling

import (
	. "github.com/lcaballero/exam/assert"
	"testing"
)

func Test_Path_001(t *testing.T) {
	t.Log("Path split should produce dir, file, and parts")
	var p Path = "/assets/src/file.html"
	parts := p.Parts()

	IsEqInt(t, parts.Len(), 3)
	IsEqStrings(t, parts.First(), "assets")
	IsEqStrings(t, parts.Last(), "file.html")
}
