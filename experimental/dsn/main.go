package main

import (
	"github.com/jinfwhuang/ds-toolkit/experimental/dsn/arweave"
)

func main() {

}

func arweaveExample() {
	id, _ := arweave.Write([]byte("123"))
	println(id)

	ret, err := arweave.Read("IUYBL-mW7OpG7Em_kwIpucrg43Br64nGbeMM01yja4w")
	println(string(ret))
	if err != nil {
		println(err.Error())
	}
}
