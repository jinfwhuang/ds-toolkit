<<<<<<< HEAD
package main

import (
	"path/filepath"
	"runtime"

	"github.com/jinfwhuang/ds-toolkit/experimental/dsn/arweave"
	"github.com/jinfwhuang/ds-toolkit/experimental/dsn/ceramic"
	estuary "github.com/jinfwhuang/ds-toolkit/experimental/dsn/filecoin-estuary"
	"github.com/jinfwhuang/ds-toolkit/experimental/dsn/sia"
	"github.com/jinfwhuang/ds-toolkit/experimental/dsn/storj"
)

var (
	data = []byte("123")
)

func main() {
}

func arweaveExample() {
	id, err := arweave.Write(data, fileInRuntimeDir("/arweave/wallet.json"))
	if err != nil {
		println(err.Error())
	}
	println(id)

	// This is an example id that should be retrieved from Write function.
	// Arweave takes couple of minutes until the id is accepted from the chain,
	// that is why we use another random id that is already in the chain here.
	ret, err := arweave.Read("IUYBL-mW7OpG7Em_kwIpucrg43Br64nGbeMM01yja4w")
	if err != nil {
		println(err.Error())
	}
	println(string(ret))
}

func estuaryExample() {
	dataPath := fileInRuntimeDir("/test_blob.json")
	resp, err := estuary.Write(dataPath)
	if err != nil {
		println(err.Error())
	}
	println(resp)
}

func storjExample() {
	key := "key_1"

	err := storj.Write(key, data)
	if err != nil {
		println(err.Error())
	}

	retr, err := storj.Read(key)
	if err != nil {
		println(err.Error())
	}
	println(string(retr))
}

func ceramicExample() {
	streamId, err := ceramic.Write(string(data))
	if err != nil {
		println(err.Error())
	}
	println(streamId)

	// This is an example id of a Stream that should be obtained from the above call.
	// However, the writing via HTTP does not seem to be maintained and currently does not work.
	// The example ID is an ID of a stream generated via the CLI.
	retr, err := ceramic.Read("kjzl6cwe1jw147z6im9aagzi92icf72n7evjt5kdhevd97ygkaxjqbicyxz618i")
	println(retr)
	if err != nil {
		println(err.Error())
	}
}

func siaExample() {
	siaDir := "example/testfileGo"
	err := sia.Write(fileInRuntimeDir("/test_file"), siaDir)
	if err != nil {
		println(err.Error())
	}

	// Most likely Read will fail, as Sia block mine time is ~10 minutes.
	retr, err := sia.Read(siaDir)
	if err != nil {
		println(err.Error())
	}
	println(retr)
}

func fileInRuntimeDir(file string) string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename) + file
}
||||||| parent of 032b42b (arweave test)
=======
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
>>>>>>> 032b42b (arweave test)
