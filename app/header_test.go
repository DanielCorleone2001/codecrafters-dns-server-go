package main

import "testing"

func Test_Header(t *testing.T) {
	h := NewDefaultHeader()
	h.ToBytes()

}
