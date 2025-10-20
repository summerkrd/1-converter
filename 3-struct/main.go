package main

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func newBin() *Bin {
	bin := &Bin{}
	return bin
}

type BinList struct {
	Bins []Bin
}

func newBinList() *BinList {
	binList := &BinList{}
	return binList
}

func main() {

}
