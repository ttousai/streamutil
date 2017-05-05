package streamutil

import (
	"testing"
)

func TestGetAddressType(t *testing.T) {
	cases := []struct {
		ins  SedInstruction
		want addressType
	}{
		{ins: SedInstruction{Address: "/foo/"}, want: regexpAddressType},
		{ins: SedInstruction{Address: "12"}, want: lineNumberAddressType},
		{ins: SedInstruction{Address: "2@"}, want: invalidAddressType},
		{ins: SedInstruction{Address: ""}, want: noAddressType},
		{ins: SedInstruction{}, want: noAddressType},
	}

	for _, c := range cases {
		got, err := getAddressType(c.ins)
		if got != c.want {
			t.Errorf("address: %v, wanted: %v, got: %v\n", c.ins.Address, c.want, got)
		}
		if err != nil {
			t.Error(err)
		}
	}
}

func TestSed(t *testing.T) {
	f := "sample.txt"

	cases := []struct {
		ins  SedInstruction
		want string
	}{
		// sed -n '/seven/p' file
		{ins: SedInstruction{Address: "/seven/", SedOptions: SedOptions{Silent: true}}, want: "line seven"},
		// sed -n '10p' file
		{ins: SedInstruction{Address: "10", SedOptions: SedOptions{Silent: true}}, want: "line seven"},
		{ins: SedInstruction{Address: "10", SedOptions: SedOptions{Silent: true}}, want: "line seven"},
	}

	for _, c := range cases {
		out := Sed(c.ins, f)
		l := <-out
		if l != c.want {
			t.Errorf("wanted: %v, got: %v\n", c.want, l)
		}
	}
}

func TestSedPipe(t *testing.T) {
	t.Error("not implemented")
}

func TestSedRaw(t *testing.T) {
	t.Error("not implemented")
}

func TestSedRawPipe(t *testing.T) {
	t.Error("not implemented")
}
