package handling

import (
	. "github.com/lcaballero/exam/assert"
	"testing"
)

const prefix = "/assets/"

func Test_MatchAssets_001(t *testing.T) {
	t.Log("")
}

func Test_PathtMatch_003(t *testing.T) {
	t.Log("should produce 'path' and 'file' when there is a path")

	var path Path = "/src/pages/file.html"
	m := AssetVars{}
	parts := path.Parts()

	m.AcquireVars(parts)

	IsEqStrings(t, m.Path(), "/src/pages")
	IsEqStrings(t, m.File(), "file.html")
}

func Test_PathMatch_002(t *testing.T) {
	t.Log("should produce 'path' and 'file'")

	var path Path = "/file.html"
	m := AssetVars{}
	parts := path.Parts()

	m.AcquireVars(parts)

	IsEqStrings(t, m.Path(), "/")
	IsEqStrings(t, m.File(), "file.html")
}

func Test_AssetMatch_002(t *testing.T) {
	t.Log("should match when prefix is present")
	vars := AssetVars{}
	isMatch := vars.AssetMatch("/assets/src/pages/file.html", prefix)
	IsNotNil(t, vars)
	IsTrue(t, isMatch)
}

func Test_AssetMatch_001(t *testing.T) {
	t.Log("should not produce a NotImplemented error")
	vars := AssetVars{}
	isMatch := vars.AssetMatch("/src/pages/file.html", prefix)
	IsNotNil(t, vars)
	IsFalse(t, isMatch)
}
