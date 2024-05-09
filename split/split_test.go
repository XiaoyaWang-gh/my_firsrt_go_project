package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}

func TestGroupSplit(t *testing.T) {
	type prepare struct {
		s    string
		sep  string
		want []string
	}

	data := map[string]prepare{
		"college":   {s: "csu-18-33", sep: "-", want: []string{"csu", "18", "33"}},
		"threebody": {s: "叶文洁·罗辑·程心", sep: "·", want: []string{"叶文", "罗辑", "程心"}},
		"seasons":   {s: "spring*summer*fall*winter", sep: "*", want: []string{"spring", "summer", "fall", "winter"}},
		"marks":     {s: "↑↓↑↓↑↓↑", sep: "↓", want: []string{"↑", "↑", "↑", "↑"}},
	}

	for name, d := range data {
		got := Split(d.s, d.sep)
		if !reflect.DeepEqual(got, d.want) {
			t.Errorf("name:%s, expected:%#v, got:%#v", name, d.want, got)
		}
	}
}
