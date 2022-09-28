package tools

import (
	"fmt"
	"testing"
)

func TestFloat64ToBytes(t *testing.T) {
	tCases := []struct {
		name string
		in   float64
		out  []byte
	}{
		{
			name: "should convert 1",
			in:   1,
			out:  []byte(fmt.Sprint(1)),
		},
		{
			name: "should convert 1.1",
			in:   1.1,
			out:  []byte(fmt.Sprint(1.1)),
		},
		{
			name: "should convert 0.1",
			in:   0.1,
			out:  []byte(fmt.Sprint(0.1)),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			if out := Float64ToBytes(tc.in); string(out) != string(tc.out) {
				tt.Errorf("unexpected return (actual=%s, expected=%s)", string(out), string(tc.out))
			}
		})
	}
}

func TestBytesToFloat64(t *testing.T) {
	tCases := []struct {
		name string
		in   []byte
		out  float64
	}{
		{
			name: "should convert 1",
			in:   Float64ToBytes(1),
			out:  1,
		},
		{
			name: "should convert 1.1",
			in:   Float64ToBytes(1.1),
			out:  1.1,
		},
		{
			name: "should convert 0.1",
			in:   Float64ToBytes(0.1),
			out:  0.1,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			out, err := BytesToFloat64(tc.in)
			if err != nil {
				tt.Errorf("unexpected error (err=%+v)", err)
			}
			if out != tc.out {
				tt.Errorf("unexpected return (actual=%f, expected=%f)", out, tc.out)
			}
		})
	}
}
