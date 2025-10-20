package bins

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func NewBin() *Bin {
	bin := &Bin{}
	return bin
}

type BinList struct {
	Bins []Bin
}

func NewBinList() *BinList {
	binList := &BinList{}
	return binList
}
