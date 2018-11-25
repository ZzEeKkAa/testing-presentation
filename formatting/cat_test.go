package formatting

import (
	"image/color"
	"testing"

	"github.com/ZzEeKkAa/testing-presentation/messagediff"
)

var redColor = color.RGBA{255, 0, 0, 255}

func TestNewCat(t *testing.T) {
	var (
		catColor = redColor
		catName  = "Lucky"
	)
	got := NewCat(catColor, catName)
	want := &Cat{
		color: catColor,
		name:  "Tom",
	}
	diff := messagediff.GitDiff(want, got)
	if diff != "" {
		t.Fatalf("NewCat(%#v,%v) want(-), got(+):\n%v",
			catColor, catName, diff)
	}
}

//--- FAIL: TestNewCat (0.00s)
//    cat_test.go:24: NewCat(color.RGBA{R:0xff, G:0x0, B:0x0, A:0xff},Lucky) want(-), got(+):
//        - name: Tom
//        + name: Lucky
//FAIL
