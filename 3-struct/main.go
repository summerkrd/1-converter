package main

import "3-struct/storage"

type BinStorage = storage.BinStorage
type Storage = storage.Storage

func main() {
	var storage BinStorage = &Storage{}
	
}
