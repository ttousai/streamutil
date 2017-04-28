package streamutil

import (
	"fmt"
	"testing"
)

func TestSed(t *testing.T) {
	f := "sample.txt"

	cases := []struct {
		in  string
		ins SedInstruction
	}{
		{f, SedInstruction{"address1", "pattern1", "action1"}},
		{f, SedInstruction{"address2", "pattern2", "action2"}},
		{f, SedInstruction{"address3", "pattern3", "action3"}},
	}

	for _, v := range cases {
		c, err := Sed(v.ins, SedOptions{}, f)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(<-c)
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
