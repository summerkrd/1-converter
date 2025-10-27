package bins

import "time"

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
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
