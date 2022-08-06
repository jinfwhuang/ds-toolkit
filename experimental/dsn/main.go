package main

import (
	"github.com/jinfwhuang/ds-toolkit/experimental/dsn/arweave"
	estuary "github.com/jinfwhuang/ds-toolkit/experimental/dsn/filecoin-estuary"
	"github.com/jinfwhuang/ds-toolkit/experimental/dsn/storj"
)

var (
	data = []byte("123")
)

func main() {

}

func arweaveExample() {
	id, err := arweave.Write(data)
	println(id)
	if err != nil {
		println(err.Error())
	}

	// This is an example id that should be retrieved from Write function.
	// Arweave takes couple of minutes until the id is accepted from the chain,
	// that is why we use another random id that is already in the chain here.
	ret, err := arweave.Read("IUYBL-mW7OpG7Em_kwIpucrg43Br64nGbeMM01yja4w")
	println(string(ret))
	if err != nil {
		println(err.Error())
	}
}

func estuaryExample() {
	resp, err := estuary.Write(data)
	println(resp)
	if err != nil {
		println(err.Error())
	}
}

func storjExample() {
	key := "key_1"

	err := storj.Write(key, data)
	if err != nil {
		println(err.Error())
	}

	retr, _ := storj.Read(key)
	println(string(retr))
	if err != nil {
		println(err.Error())
	}
}
