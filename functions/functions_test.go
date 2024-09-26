package ascii

import (
	"reflect"
	"testing"
)

func TestSplitNewLine(t *testing.T) {
	got1 := SplitNewLine("\\n\\n")
	want1 := []string{"", ""}
	got2 := SplitNewLine(`\na\n`)
	want2 := []string{"", "a", ""}
	array_got := [][]string{got1, got2}
	array_want := [][]string{want1, want2}
	for i, arr := range array_got {
		if !reflect.DeepEqual(arr, array_want[i]) {
			t.Errorf("got %q, wanted %q", arr, array_want[i])
		}
	}
}

func TestCreateMap(t *testing.T) {
  
}
