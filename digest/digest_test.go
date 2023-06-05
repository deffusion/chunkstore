package digest

import (
	"fmt"
	"testing"
)

func TestDigest(t *testing.T) {
	d1, _ := New("68cf29c8adf41458783f4907695067dcc21d751c")
	fmt.Printf("bytes %x\n", d1.Bytes())
	d2, err := New(d1.Bytes())
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d1 == d2)
}
