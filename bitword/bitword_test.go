package bitword

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToAndFromStr(t *testing.T) {
	cases := []struct {
		src  string
		n    int
		want []byte
	}{
		{"", 1,
			[]byte{}},
		{"", 2,
			[]byte{}},
		{"", 4,
			[]byte{}},
		{"", 8,
			[]byte{}},
		{"a", 1,
			[]byte{0, 1, 1, 0, 0, 0, 0, 1}},
		{"a", 2,
			[]byte{0x1, 0x2, 0x0, 0x1}},
		{"a", 4,
			[]byte{0x6, 0x1}},
		{"a", 8,
			[]byte{0x61}},
		{"\x00", 4,
			[]byte{0, 0}},
		{"\x01\x02\xff", 1,
			[]byte{
				0, 0, 0, 0, 0, 0, 0, 1,
				0, 0, 0, 0, 0, 0, 1, 0,
				1, 1, 1, 1, 1, 1, 1, 1}},
		{"\x01\x02\xff", 2,
			[]byte{0, 0, 0, 1, 0, 0, 0, 2, 3, 3, 3, 3}},
		{"\x01\x02\xff", 4,
			[]byte{0, 1, 0, 2, 0xf, 0xf}},
		{"\x01\x02\xff", 8,
			[]byte{1, 2, 0xff}},
		{"我", 1,
			[]byte{
				1, 1, 1, 0, 0, 1, 1, 0,
				1, 0, 0, 0, 1, 0, 0, 0,
				1, 0, 0, 1, 0, 0, 0, 1,
			},
		},
		{"我", 2,
			[]byte{
				3, 2, 1, 2,
				2, 0, 2, 0,
				2, 1, 0, 1,
			},
		},
		{"我", 4,
			[]byte{0xe, 0x6, 0x8, 0x8, 0x9, 0x1}},
		{"我", 8,
			[]byte{0xe6, 0x88, 0x91},
		},
	}

	for i, c := range cases {
		res := FromStr(c.src, c.n)

		if !reflect.DeepEqual(res, c.want) {
			t.Errorf("test %d: got %#v, want %#v",
				i+1, res, c.want)
		}

		str := ToStr(res, c.n)
		if str != c.src {
			t.Fatalf(" expect: %v; but: %v", c.src, str)
		}
	}
}

func TestToStrIncomplete(t *testing.T) {
	ta := require.New(t)
	got := ToStr([]byte{1, 2, 3}, 4)
	want := "\x12\x30"
	ta.Equal(want, got)
}

func TestToAndFromStrPanic(t *testing.T) {

	ta := require.New(t)

	ta.Panics(func() { FromStr("1", -1) }, "ToBitWords, n= -1")
	ta.Panics(func() { FromStr("1", 0) }, "ToBitWords, n= 0")
	ta.Panics(func() { FromStr("1", 3) }, "ToBitWords, n= 3")
	ta.Panics(func() { FromStr("1", 5) }, "ToBitWords, n= 5")
	ta.Panics(func() { FromStr("1", 6) }, "ToBitWords, n= 6")
	ta.Panics(func() { FromStr("1", 7) }, "ToBitWords, n= 7")
	ta.Panics(func() { FromStr("1", 9) }, "ToBitWords, n= 9")

	ta.Panics(func() { ToStr([]byte{1}, -1) }, "ToBitWords, n= -1")
	ta.Panics(func() { ToStr([]byte{1}, 0) }, "ToBitWords, n= 0")
	ta.Panics(func() { ToStr([]byte{1}, 3) }, "ToBitWords, n= 3")
	ta.Panics(func() { ToStr([]byte{1}, 5) }, "ToBitWords, n= 5")
	ta.Panics(func() { ToStr([]byte{1}, 6) }, "ToBitWords, n= 6")
	ta.Panics(func() { ToStr([]byte{1}, 7) }, "ToBitWords, n= 7")
	ta.Panics(func() { ToStr([]byte{1}, 9) }, "ToBitWords, n= 9")
}

func TestToAndFromStrs(t *testing.T) {

	cases := []struct {
		input []string
		n     int
		want  [][]byte
	}{
		{[]string{"a", "bc", "d"}, 4,
			[][]byte{
				{6, 1},
				{6, 2, 6, 3},
				{6, 4},
			},
		},
		{[]string{"a", "bc", "d"}, 2,
			[][]byte{
				{1, 2, 0, 1},
				{1, 2, 0, 2, 1, 2, 0, 3},
				{1, 2, 1, 0},
			},
		},
	}

	for i, c := range cases {
		rst := FromStrs(c.input, c.n)
		if !reflect.DeepEqual(c.want, rst) {
			t.Fatalf("%d-th: input: %v; want: %v; actual: %v",
				i+1, c.input, c.want, rst)
		}

		strs := ToStrs(rst, c.n)
		if !reflect.DeepEqual(c.input, strs) {
			t.Fatalf("%d-th expect: %v; but: %v", i+1, c.input, strs)
		}
	}
}
