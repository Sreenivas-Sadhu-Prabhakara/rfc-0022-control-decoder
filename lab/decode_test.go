package lab

import (
	"reflect"
	"strings"
	"testing"
)

func TestEveryOpcodeAndBoundary(t *testing.T) {
	const input = "00 80 01 00 FF 02 11 22 03 33 44 04 55 05 66 77 06 88 07 99 08 aa"
	want := []Command{
		{0, 0, "Open primary", []int{128}},
		{2, 1, "Open auxiliary", []int{0, 255}},
		{5, 2, "Accept primary", []int{17, 34}},
		{8, 3, "Accept auxiliary", []int{51, 68}},
		{11, 4, "Reject primary", []int{85}},
		{13, 5, "Reject auxiliary", []int{102, 119}},
		{16, 6, "CEASE", []int{136}},
		{18, 7, "CLOSING", []int{153}},
		{20, 8, "UNCEASE", []int{170}},
	}
	got, err := DecodeHex(input)
	if err != nil {
		t.Fatal(err)
	}
	if got.Bytes != 22 || got.Hex != strings.ToUpper(input) || !reflect.DeepEqual(got.Commands, want) {
		t.Fatalf("unexpected decode: %+v", got)
	}
}

func TestMalformedPayloads(t *testing.T) {
	for _, tc := range []struct{ input, want string }{
		{" \n", "empty payload"},
		{"00 12 z3", "invalid hex at byte offset 2"},
		{"00 123", "expected two hex digits at byte offset 1"},
		{"00 12 FE", "unknown opcode 0xFE at byte offset 2"},
		{"00 12 01 34", "truncated opcode 0x01 at byte offset 2: expected 2 parameter bytes, got 1"},
		{"06", "truncated opcode 0x06 at byte offset 0: expected 1 parameter bytes, got 0"},
	} {
		got, err := DecodeHex(tc.input)
		if err == nil || !strings.Contains(err.Error(), tc.want) || len(got.Commands) != 0 {
			t.Errorf("DecodeHex(%q) = %+v, %v, want %q and no commands", tc.input, got, err, tc.want)
		}
	}
}
